# Simple VM Maker

This is a simple Virtual Machine maker application built with Go. It is a API that allows users to create and manager virtual machines.
- The application is built based on the [cloud-hypervisor](https://github.com/cloud-hypervisor/cloud-hypervisor)

## Prerequisites

- Go 1.23.2
- linux kernel 3.0 or later for using kvm
    - check if your system supports kvm by running the following command:
    ```bash
    lsmod | grep kvm
   ```
- cloud-hypervisor installed on your system
  - Go to the [cloud-hypervisor](https://github.com/cloud-hypervisor/cloud-hypervisor/releases/) and donwload the `cloud-hypervisor`
  - Move the `cloud-hypervisor` to the `/usr/local/bin` directory (or any other directory and modify the `CLOUD_HYPERVISOR_BIN_PATH` in the `.env` file)
