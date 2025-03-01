package utils

import (
	"os"
	"vm-maker/config"
	"vm-maker/model/dto"

	"gopkg.in/yaml.v2"
)

func CreateCloudInit(
	vm dto.VmExecution,
	settings config.Settings,
) error {
	tempDir, err := os.MkdirTemp(settings.TempDirPath, "cloud-init-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	userDataFilePath := tempDir + "/user-data"
	networkDataFilePath := tempDir + "/network-config"
	metaDataFilePath := tempDir + "/meta-data"

	userData := map[string]interface{}{
		"users": []map[string]interface{}{
			{
				"name":        vm.Username,
				"passwd":      vm.PasswordHash,
				"lock_passwd": false,
				"shell":       "/bin/bash",
			},
		},
		"ssh_pwauth": true,
	}

	networkData := map[string]interface{}{
		"version": 2,
		"ethernets": map[string]interface{}{
			"ens4": map[string]interface{}{
				"match": map[string]interface{}{
					"macaddress": "12:34:56:78:90:ab",
				},
				"addresses": []string{
					"192.168.249.2/24",
				},
				"gateway4": "192.168.249.1",
			},
		},
	}

	metaData := map[string]interface{}{
		"instance-id":    "cloud",
		"local-hostname": vm.Hostname,
	}

	if err := writeYamlToFile(userDataFilePath, userData, "#cloud-config\n"); err != nil {
		return err
	}

	if err := writeYamlToFile(networkDataFilePath, networkData, ""); err != nil {
		return err
	}

	if err := writeYamlToFile(metaDataFilePath, metaData, ""); err != nil {
		return err
	}

	if _, err := RunShellCommand(
		settings.BinPathBook.Mkdosfs,
		"-n", "CIDATA",
		"-C",
		vm.CloudImageFilePath,
		"8192",
	); err != nil {
		return err
	}

	_, err = RunShellCommand(
		settings.BinPathBook.Mcopy,
		"-oi",
		vm.CloudImageFilePath,
		"-s",
		userDataFilePath,
		"::",
	)
	if err != nil {
		return err
	}

	_, err = RunShellCommand(
		settings.BinPathBook.Mcopy,
		"-oi",
		vm.CloudImageFilePath,
		"-s",
		networkDataFilePath,
		"::",
	)
	if err != nil {
		return err
	}

	_, err = RunShellCommand(
		settings.BinPathBook.Mcopy,
		"-oi",
		vm.CloudImageFilePath,
		"-s",
		metaDataFilePath,
		"::",
	)
	if err != nil {
		return err
	}

	return nil
}

func writeYamlToFile(filePath string, data map[string]interface{}, prefix string) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(prefix + string(yamlData)))
	return err
}
