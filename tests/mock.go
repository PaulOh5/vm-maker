package tests

import (
	"vm-maker/model/dto"

	"github.com/google/uuid"
)

func CreateDummyVmExecution() dto.VmExecution {
	vmUUID := uuid.New().String()
	vm := dto.VmExecution{
		CloudImageFilePath: "/tmp/" + vmUUID + ".img",
		VmSocketSocketPath: "/tmp/" + vmUUID + ".sock",
		Username:           "cloud",
		Hostname:           "host1",
		PasswordHash:       "$6$7125787751a8d18a$sHwGySomUA1PawiNFWVCKYQN.Ec.Wzz0JtPPL1MvzFrkwmop2dq7.4CYf03A5oemPQ4pOFCCrtCelvFBEle/K.",
		KernelPath:         "/opt/vm-maker/hypervisor-fw",
		OsDisk:             "/opt/vm-maker/ubuntu.raw",
		Disks:              []string{},
	}

	return vm
}
