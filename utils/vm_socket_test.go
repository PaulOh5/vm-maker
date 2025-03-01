package utils

import (
	"os"
	"testing"

	"vm-maker/config"
	"vm-maker/tests"
)

func TestCreateVmSocket(t *testing.T) {
	settings := config.SetupSettings()
	vm := tests.CreateDummyVmExecution()

	vmSocket, err := CreateVmSocket(vm, settings, 100)

	if err != nil {
		t.Fatalf("CreateVmSocket() error = %v", err)
	} else {
		if _, err := os.Stat(vm.VmSocketSocketPath); err != nil {
			t.Fatalf("vmSocket socket file not found: %v", err)
		}
		vmSocket.Cmd.Process.Signal(os.Interrupt)
	}
}

func TestVmSocket(t *testing.T) {
	settings := config.SetupSettings()
	dummyVm := tests.CreateDummyVmExecution()

	if err := CreateCloudInit(dummyVm, settings); err != nil {
		t.Fatalf("CreateCloudInit() failed, error = %v", err)
	}
	defer os.Remove(dummyVm.CloudImageFilePath)

	vmSocket, err := CreateVmSocket(dummyVm, settings, 100)
	if err != nil {
		t.Fatalf("CreateVmSocket() failed, error = %v", err)
	} else {
		defer vmSocket.Cmd.Process.Signal(os.Interrupt)
	}

	err = vmSocket.CreateVm()
	if err != nil {
		t.Fatalf("vmSocket.CreateVm() failed, error = %v", err)
	}

	if vmInfo, err := vmSocket.GetVmInfo(); err != nil {
		t.Fatalf("vmSocket.GetVmInfo() failed, error = %v", err)
	} else {
		if vmInfo.State != "Created" {
			t.Fatalf("vmInfo.State = %s, want %s", vmInfo.State, "created")
		}
	}

	if err := vmSocket.Boot(); err != nil {
		t.Fatalf("vmSocket.Boot() failed, error = %v", err)
	}

	if vmInfo, err := vmSocket.GetVmInfo(); err != nil {
		t.Fatalf("vmSocket.GetVmInfo() failed, error = %v", err)
	} else {
		if vmInfo.State != "Running" {
			t.Fatalf("vmInfo.State = %s, want %s", vmInfo.State, "running")
		}
	}
}
