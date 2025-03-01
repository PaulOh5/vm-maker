package dto

type VmExecution struct {
	CloudImageFilePath string
	VmSocketSocketPath string
	Username           string
	Hostname           string
	PasswordHash       string
	KernelPath         string
	OsDisk             string
	Disks              []string
}
