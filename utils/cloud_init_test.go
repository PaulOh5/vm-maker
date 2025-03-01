package utils

import (
	"os"
	"testing"
	"vm-maker/config"
	"vm-maker/tests"
)

func TestCreateCloudInit(t *testing.T) {
	settings := config.SetupSettings()
	vm := tests.CreateDummyVmExecution()

	defer os.Remove(vm.CloudImageFilePath)

	if err := CreateCloudInit(vm, settings); err != nil {
		t.Fatalf("CreateCloudInit() failed, error = %v", err)
	} else {
		if _, err := os.Stat(vm.CloudImageFilePath); os.IsNotExist(err) {
			t.Fatalf("CreateCloudInit() failed, cloud-init file not created")
		}
	}
}
