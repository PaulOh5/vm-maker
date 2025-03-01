package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"syscall"
)

func RunShellCommand(bin string, args ...string) (string, error) {
	cmd := exec.Command(bin, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%v%v", stderr.String(), err)
	}

	return out.String(), nil
}

func RunShellProcess(bin string, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(bin, args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	stderrBuf := new(bytes.Buffer)
	cmd.Stderr = stderrBuf

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func IsFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CheckGroup(groupName string) (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, err
	}

	groupIds, err := currentUser.GroupIds()
	if err != nil {
		return false, err
	}

	for _, gid := range groupIds {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			return false, err
		}

		if group.Name == groupName {
			return true, nil
		}
	}

	return false, nil
}
