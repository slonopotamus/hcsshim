package annotations

const (
	// ContainerMemorySizeInMB overrides the container memory size set
	// via the OCI spec.
	//
	// Note: This annotation is in MB. OCI is in Bytes. When using this override
	// the caller MUST use MB or sizing will be wrong.
	//
	// Note: This is only present because CRI does not (currently) have a
	// `WindowsPodSandboxConfig` for setting this correctly. It should not be
	// used via OCI runtimes and rather use
	// `spec.Windows.Resources.Memory.Limit`.
	ContainerMemorySizeInMB = "io.microsoft.container.memory.sizeinmb"

	// ContainerProcessorCount overrides the container processor count
	// set via the OCI spec.
	//
	// Note: For Windows Process Containers CPU Count/Limit/Weight are mutually
	// exclusive and the caller MUST only set one of the values.
	//
	// Note: This is only present because CRI does not (currently) have a
	// `WindowsPodSandboxConfig` for setting this correctly. It should not be
	// used via OCI runtimes and rather use `spec.Windows.Resources.CPU.Count`.
	ContainerProcessorCount = "io.microsoft.container.processor.count"

	// ContainerProcessorLimit overrides the container processor limit
	// set via the OCI spec.
	//
	// Limit allows values 1 - 10,000 where 10,000 means 100% CPU. (And is the
	// default if omitted)
	//
	// Note: For Windows Process Containers CPU Count/Limit/Weight are mutually
	// exclusive and the caller MUST only set one of the values.
	//
	// Note: This is only present because CRI does not (currently) have a
	// `WindowsPodSandboxConfig` for setting this correctly. It should not be
	// used via OCI runtimes and rather use
	// `spec.Windows.Resources.CPU.Maximum`.
	ContainerProcessorLimit = "io.microsoft.container.processor.limit"

	// ContainerProcessorWeight overrides the container processor
	// weight set via the OCI spec.
	//
	// Weight allows values 0 - 10,000. (100 is the default)
	//
	// Note: For Windows Process Containers CPU Count/Limit/Weight are mutually
	// exclusive and the caller MUST only set one of the values.
	//
	// Note: This is only present because CRI does not (currently) have a
	// `WindowsPodSandboxConfig` for setting this correctly. It should not be
	// used via OCI runtimes and rather use `spec.Windows.Resources.CPU.Shares`.
	ContainerProcessorWeight = "io.microsoft.container.processor.weight"

	// ContainerStorageQoSBandwidthMaximum overrides the container
	// storage bandwidth per second set via the OCI spec.
	//
	// Note: This is only present because CRI does not (currently) have a
	// `WindowsPodSandboxConfig` for setting this correctly. It should not be
	// used via OCI runtimes and rather use
	// `spec.Windows.Resources.Storage.Bps`.
	ContainerStorageQoSBandwidthMaximum = "io.microsoft.container.storage.qos.bandwidthmaximum"

	// ContainerStorageQoSIopsMaximum overrides the container storage
	// maximum iops set via the OCI spec.
	//
	// Note: This is only present because CRI does not (currently) have a
	// `WindowsPodSandboxConfig` for setting this correctly. It should not be
	// used via OCI runtimes and rather use
	// `spec.Windows.Resources.Storage.Iops`.
	ContainerStorageQoSIopsMaximum = "io.microsoft.container.storage.qos.iopsmaximum"

	// GPUVHDPath overrides the default path to search for the gpu vhd
	GPUVHDPath = "io.microsoft.lcow.gpuvhdpath"

	// ContainerGPUCapabilities is used to find the gpu capabilities on the container spec
	ContainerGPUCapabilities = "io.microsoft.container.gpu.capabilities"

	// VirtualMachineKernelDrivers indicates what drivers to install in the pod.
	// This value should contain a list of comma separated directories containing all
	// files and information needed to install given driver(s). For windows, this may
	// include .sys, .inf, .cer, and/or other files used during standard installation with pnputil.
	// For LCOW, this may include a vhd file that contains kernel modules as *.ko files.
	VirtualMachineKernelDrivers = "io.microsoft.virtualmachine.kerneldrivers"

	// DeviceExtensions contains a comma separated list of full paths to device extension files.
	// The content of these are added to a container's hcs create document.
	DeviceExtensions = "io.microsoft.container.wcow.deviceextensions"

	// HostProcessInheritUser indicates whether to ignore the username passed in to run a host process
	// container as and instead inherit the user token from the executable that is launching the container process.
	HostProcessInheritUser = "microsoft.com/hostprocess-inherit-user"

	// HostProcessContainer indicates to launch a host process container (job container in this repository).
	HostProcessContainer = "microsoft.com/hostprocess-container"

	// AllowOvercommit indicates if we should allow over commit memory for UVM.
	// Defaults to true. For physical backed memory, set to false.
	AllowOvercommit = "io.microsoft.virtualmachine.computetopology.memory.allowovercommit"

	// EnableDeferredCommit indicates if we should allow deferred memory commit for UVM.
	// Defaults to false. For virtual memory with deferred commit, set to true.
	EnableDeferredCommit = "io.microsoft.virtualmachine.computetopology.memory.enabledeferredcommit"

	// EnableColdDiscardHint indicates whether to enable cold discard hint, which allows the UVM
	// to trim non-zeroed pages from the working set (if supported by the guest operating system).
	EnableColdDiscardHint = "io.microsoft.virtualmachine.computetopology.memory.enablecolddiscardhint"

	// MemorySizeInMB overrides the container memory size set via the
	// OCI spec.
	//
	// Note: This annotation is in MB. OCI is in Bytes. When using this override
	// the caller MUST use MB or sizing will be wrong.
	MemorySizeInMB = "io.microsoft.virtualmachine.computetopology.memory.sizeinmb"

	// MemoryLowMMIOGapInMB indicates the low MMIO gap in MB
	MemoryLowMMIOGapInMB = "io.microsoft.virtualmachine.computetopology.memory.lowmmiogapinmb"

	// MemoryHighMMIOBaseInMB indicates the high MMIO base in MB
	MemoryHighMMIOBaseInMB = "io.microsoft.virtualmachine.computetopology.memory.highmmiobaseinmb"

	// MemoryHighMMIOBaseInMB indicates the high MMIO gap in MB
	MemoryHighMMIOGapInMB = "io.microsoft.virtualmachine.computetopology.memory.highmmiogapinmb"

	// ProcessorCount overrides the hypervisor isolated vCPU count set
	// via the OCI spec.
	//
	// Note: Unlike Windows process isolated container QoS Count/Limt/Weight on
	// the UVM are not mutually exclusive and can be set together.
	ProcessorCount = "io.microsoft.virtualmachine.computetopology.processor.count"

	// ProcessorLimit overrides the hypervisor isolated vCPU limit set
	// via the OCI spec.
	//
	// Limit allows values 1 - 100,000 where 100,000 means 100% CPU. (And is the
	// default if omitted)
	//
	// Note: Unlike Windows process isolated container QoS Count/Limt/Weight on
	// the UVM are not mutually exclusive and can be set together.
	ProcessorLimit = "io.microsoft.virtualmachine.computetopology.processor.limit"

	// ProcessorWeight overrides the hypervisor isolated vCPU weight set
	// via the OCI spec.
	//
	// Weight allows values 0 - 10,000. (100 is the default if omitted)
	//
	// Note: Unlike Windows process isolated container QoS Count/Limt/Weight on
	// the UVM are not mutually exclusive and can be set together.
	ProcessorWeight = "io.microsoft.virtualmachine.computetopology.processor.weight"

	// VPMemCount indicates the max number of vpmem devices that can be used on the UVM
	VPMemCount = "io.microsoft.virtualmachine.devices.virtualpmem.maximumcount"

	// VPMemSize indicates the size of the VPMem devices.
	VPMemSize = "io.microsoft.virtualmachine.devices.virtualpmem.maximumsizebytes"

	// PreferredRootFSType indicates what the preferred rootfs type should be for an LCOW UVM.
	// valid values are "initrd" or "vhd"
	PreferredRootFSType = "io.microsoft.virtualmachine.lcow.preferredrootfstype"

	// BootFilesRootPath indicates the path to find the LCOW boot files to use when creating the UVM
	BootFilesRootPath = "io.microsoft.virtualmachine.lcow.bootfilesrootpath"

	// KernelDirectBoot indicates that we should skip UEFI and boot directly to `kernel`
	KernelDirectBoot = "io.microsoft.virtualmachine.lcow.kerneldirectboot"

	// VPCIEnabled indicates that pci support should be enabled for the LCOW UVM
	VPCIEnabled = "io.microsoft.virtualmachine.lcow.vpcienabled"

	// VPMemNoMultiMapping indicates that we should disable LCOW vpmem layer multi mapping
	VPMemNoMultiMapping = "io.microsoft.virtualmachine.lcow.vpmem.nomultimapping"

	// KernelBootOptions is used to specify kernel options used while booting a linux kernel
	KernelBootOptions = "io.microsoft.virtualmachine.lcow.kernelbootoptions"

	// StorageQoSBandwidthMaximum indicates the maximum number of bytes per second. If `0`
	// will default to the platform default.
	StorageQoSBandwidthMaximum = "io.microsoft.virtualmachine.storageqos.bandwidthmaximum"

	// StorageQoSIopsMaximum indicates the maximum number of Iops. If `0` will
	// default to the platform default.
	StorageQoSIopsMaximum = "io.microsoft.virtualmachine.storageqos.iopsmaximum"

	// FullyPhysicallyBacked indicates that the UVM should use physically backed memory only,
	// including for additional devices added later.
	FullyPhysicallyBacked = "io.microsoft.virtualmachine.fullyphysicallybacked"

	// DisableCompartmentNamespace sets whether to disable namespacing the network compartment in the UVM
	// for WCOW.
	DisableCompartmentNamespace = "io.microsoft.virtualmachine.disablecompartmentnamespace"

	// VSMBNoDirectMap specifies that no direct mapping should be used for any VSMBs added to the UVM
	VSMBNoDirectMap = "io.microsoft.virtualmachine.wcow.virtualSMB.nodirectmap"

	// CPUGroupID specifies the cpugroup ID that a UVM should be assigned to if any
	CPUGroupID = "io.microsoft.virtualmachine.cpugroup.id"

	// SaveAsTemplate annotation must be used with a pod & container creation request.
	// If this annotation is present in the request then it will save the UVM (pod)
	// and the container(s) inside it as a template. However, this also means that this
	// pod and the containers inside this pod will permananetly stay in the
	// paused/templated state and can not be resumed again.
	SaveAsTemplate = "io.microsoft.virtualmachine.saveastemplate"

	// TemplateID should be used when creating a pod or a container from a template.
	// When creating a pod from a template use the ID of the templated pod as the
	// TemplateID and when creating a container use the ID of the templated container as
	// the TemplateID. It is the client's responsibility to make sure that the sandbox
	// within which a cloned container needs to be created must also be created from the
	// same template.
	TemplateID = "io.microsoft.virtualmachine.templateid"

	// NetworkConfigProxy holds the address of the network config proxy service.
	// If set, network setup will be attempted via ncproxy.
	NetworkConfigProxy = "io.microsoft.network.ncproxy"

	// NcproxyContainerID indicates whether or not to use the hcsshim container ID
	// when setting up ncproxy and computeagent
	NcproxyContainerID = "io.microsoft.network.ncproxy.containerid"

	// EncryptedScratchDisk indicates whether or not the container scratch disks
	// should be encrypted or not
	EncryptedScratchDisk = "io.microsoft.virtualmachine.storage.scratch.encrypted"

	// SecurityPolicy is used to specify a security policy for opengcs to enforce
	SecurityPolicy = "io.microsoft.virtualmachine.lcow.securitypolicy"
	// ContainerProcessDumpLocation specifies a path inside of containers to save process dumps to. As
	// the scratch space for a container is generally cleaned up after exit, this is best set to a volume mount of
	// some kind (vhd, bind mount, fileshare mount etc.)
	ContainerProcessDumpLocation = "io.microsoft.container.processdumplocation"

	// WCOWProcessDumpType specifies the type of dump to create when generating a local user mode
	// process dump for Windows containers. The supported options are "mini", and "full".
	// See DumpType: https://docs.microsoft.com/en-us/windows/win32/wer/collecting-user-mode-dumps
	WCOWProcessDumpType = "io.microsoft.wcow.processdumptype"

	// RLimitCore specifies the core rlimit value for a container. This will need to be set
	// in order to have core dumps generated for a given container.
	RLimitCore = "io.microsoft.lcow.rlimitcore"

	// LCOWDevShmSizeInKb specifies the size of LCOW /dev/shm.
	LCOWDevShmSizeInKb = "io.microsoft.lcow.shm.size-kb"

	// LCOWPrivileged is used to specify that the container should be run in privileged mode
	LCOWPrivileged = "io.microsoft.virtualmachine.lcow.privileged"

	// KubernetesContainerType is the annotation used by CRI to define the `ContainerType`.
	KubernetesContainerType = "io.kubernetes.cri.container-type"

	// KubernetesSandboxID is the annotation used by CRI to define the
	// KubernetesContainerType == "sandbox"` ID.
	KubernetesSandboxID = "io.kubernetes.cri.sandbox-id"
)
