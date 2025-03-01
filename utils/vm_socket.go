package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"time"

	"vm-maker/config"
	"vm-maker/model/dto"
	orm "vm-maker/model/orm/vm_socket"
)

type VmSocket struct {
	Cmd        *exec.Cmd
	socketPath string
	client     *http.Client
	execution  *dto.VmExecution
}

func (vmSocket *VmSocket) get(path string) (string, error) {
	url := url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   path,
	}

	resp, err := vmSocket.client.Get(url.String())
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"unexpected status code: %d, status: %s, body: %s",
			resp.StatusCode,
			resp.Status,
			respBody,
		)
	}

	return string(respBody), nil
}

func (vmSocket *VmSocket) put(path string, body any) error {
	url := url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   path,
	}

	var req *http.Request
	var err error

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal json: %v", err)
		}

		req, err = http.NewRequest(http.MethodPut, url.String(), bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
	} else {
		req, err = http.NewRequest(http.MethodPut, url.String(), nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := vmSocket.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf(
			"error: %d, status: %s, body: %s",
			resp.StatusCode,
			resp.Status,
			respBody,
		)
	}

	return nil
}

func (vmSocket *VmSocket) Ping() error {
	_, err := vmSocket.get("/api/v1/vmm.ping")
	return err
}

func (vmSocket *VmSocket) CreateVm() error {
	disks := make([]map[string]string, len(vmSocket.execution.Disks)+2)
	disks[0] = map[string]string{"path": vmSocket.execution.OsDisk}
	disks[1] = map[string]string{"path": vmSocket.execution.CloudImageFilePath}

	for i, disk := range vmSocket.execution.Disks {
		disks[i+2] = map[string]string{"path": disk}
	}

	vmConfig := map[string]any{
		"payload": map[string]any{
			"kernel":  vmSocket.execution.KernelPath,
			"cmdline": "root=/dev/vda1",
		},
		"disks": disks,
		"console": map[string]any{
			"mode": "Null",
		},
		"serial": map[string]any{
			"mode": "Pty",
		},
	}

	err := vmSocket.put("/api/v1/vm.create", vmConfig)
	return err
}

func (vmSocket *VmSocket) GetVmInfo() (*orm.VmInfo, error) {
	res, err := vmSocket.get("/api/v1/vm.info")
	if err != nil {
		return nil, err
	}

	var vmInfo orm.VmInfo
	if err := json.Unmarshal([]byte(res), &vmInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vm info: %v", err)
	}

	return &vmInfo, nil
}

func (vmSocket *VmSocket) Boot() error {
	err := vmSocket.put("/api/v1/vm.boot", nil)
	if err != nil {
		return err
	}

	vmInfo, err := vmSocket.GetVmInfo()
	if err != nil {
		return err
	}

	log.Printf("serial file path: %s", vmInfo.Config.Serial.File)
	return nil
}

func CreateVmSocket(
	vm dto.VmExecution,
	settings config.Settings,
	timeoutMiliSeconds float32,
) (*VmSocket, error) {
	if timeoutMiliSeconds == 0 {
		timeoutMiliSeconds = 10_000 // 10 seconds
	}

	cmd, err := RunShellProcess(
		settings.BinPathBook.CloudHyperVisor,
		"--api-socket",
		vm.VmSocketSocketPath,
	)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			dialer := net.Dialer{}
			return dialer.DialContext(ctx, "unix", vm.VmSocketSocketPath)
		},
	}

	client := &http.Client{Transport: transport}

	vmSocket := &VmSocket{
		Cmd:        cmd,
		socketPath: vm.VmSocketSocketPath,
		client:     client,
		execution:  &vm,
	}

	if err := waitForSocketReady(vmSocket, timeoutMiliSeconds); err != nil {
		err := cmd.Wait()

		if stderrBuf, ok := cmd.Stderr.(*bytes.Buffer); ok {
			if err != nil {
				return nil, fmt.Errorf(
					"%s%v", stderrBuf.String(), err,
				)
			}

			return nil, err
		}

		return nil, err
	}

	return vmSocket, nil
}

func waitForSocketReady(vmSocket *VmSocket, timeoutSec float32) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeoutSec)*time.Millisecond,
	)
	defer cancel()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("failed to connect to vm socket: %v", ctx.Err())
		case <-ticker.C:
			err := vmSocket.Ping()
			if err == nil {
				return nil
			}
		}
	}
}
