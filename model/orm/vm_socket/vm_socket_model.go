package orm

type VmInfo struct {
	Config struct {
		Cpus struct {
			BootVcpus   int  `json:"boot_vcpus"`
			MaxVcpus    int  `json:"max_vcpus"`
			Topology    any  `json:"topology"`
			KvmHyperv   bool `json:"kvm_hyperv"`
			MaxPhysBits int  `json:"max_phys_bits"`
			Affinity    any  `json:"affinity"`
			Features    struct {
				Amx bool `json:"amx"`
			} `json:"features"`
		} `json:"cpus"`
		Memory struct {
			Size           int    `json:"size"`
			Mergeable      bool   `json:"mergeable"`
			HotplugMethod  string `json:"hotplug_method"`
			HotplugSize    any    `json:"hotplug_size"`
			HotpluggedSize any    `json:"hotplugged_size"`
			Shared         bool   `json:"shared"`
			Hugepages      bool   `json:"hugepages"`
			HugepageSize   any    `json:"hugepage_size"`
			Prefault       bool   `json:"prefault"`
			Zones          any    `json:"zones"`
			Thp            bool   `json:"thp"`
		} `json:"memory"`
		Payload struct {
			Firmware  any    `json:"firmware"`
			Kernel    string `json:"kernel"`
			Cmdline   string `json:"cmdline"`
			Initramfs any    `json:"initramfs"`
		} `json:"payload"`
		RateLimitGroups any `json:"rate_limit_groups"`
		Disks           []struct {
			Path              string `json:"path"`
			Readonly          bool   `json:"readonly"`
			Direct            bool   `json:"direct"`
			Iommu             bool   `json:"iommu"`
			NumQueues         int    `json:"num_queues"`
			QueueSize         int    `json:"queue_size"`
			VhostUser         bool   `json:"vhost_user"`
			VhostSocket       any    `json:"vhost_socket"`
			RateLimitGroup    any    `json:"rate_limit_group"`
			RateLimiterConfig any    `json:"rate_limiter_config"`
			ID                any    `json:"id"`
			DisableIoUring    bool   `json:"disable_io_uring"`
			DisableAio        bool   `json:"disable_aio"`
			PciSegment        int    `json:"pci_segment"`
			Serial            any    `json:"serial"`
			QueueAffinity     any    `json:"queue_affinity"`
		} `json:"disks"`
		Net any `json:"net"`
		Rng struct {
			Src   string `json:"src"`
			Iommu bool   `json:"iommu"`
		} `json:"rng"`
		Balloon any `json:"balloon"`
		Fs      any `json:"fs"`
		Pmem    any `json:"pmem"`
		Serial  struct {
			File   string `json:"file"`
			Mode   string `json:"mode"`
			Iommu  bool   `json:"iommu"`
			Socket any    `json:"socket"`
		} `json:"serial"`
		Console struct {
			File   any    `json:"file"`
			Mode   string `json:"mode"`
			Iommu  bool   `json:"iommu"`
			Socket any    `json:"socket"`
		} `json:"console"`
		DebugConsole struct {
			File   any    `json:"file"`
			Mode   string `json:"mode"`
			Iobase int    `json:"iobase"`
		} `json:"debug_console"`
		Devices        any  `json:"devices"`
		UserDevices    any  `json:"user_devices"`
		Vdpa           any  `json:"vdpa"`
		Vsock          any  `json:"vsock"`
		Pvpanic        bool `json:"pvpanic"`
		Iommu          bool `json:"iommu"`
		SgxEpc         any  `json:"sgx_epc"`
		Numa           any  `json:"numa"`
		Watchdog       bool `json:"watchdog"`
		PciSegments    any  `json:"pci_segments"`
		Platform       any  `json:"platform"`
		Tpm            any  `json:"tpm"`
		LandlockEnable bool `json:"landlock_enable"`
		LandlockRules  any  `json:"landlock_rules"`
	} `json:"config"`
	State            string `json:"state"`
	MemoryActualSize int    `json:"memory_actual_size"`
	DeviceTree       any    `json:"device_tree"`
}
