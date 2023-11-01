// Package instance contains structs for the Instance payload.
package instance

import (
	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
	"github.com/softiron/hypercloud-api/internal/time"
)

//go:generate go run "github.com/dmarkham/enumer" -type State -linecomment -text

// State represents the state of an instance.
type State int

// State values.
const (
	InitState      State = iota // init
	PendingState                // pending
	HoldState                   // hold
	ActiveState                 // active
	StoppedState                // stopped
	SuspendedState              // suspended
	DoneState                   // done
	_
	OffState                      // off
	UndeployedState               // undeployed
	CloningState                  // cloning
	CloningFailedState            // cloning_failed
	AnyState           State = -2 // all
	NotDoneState       State = -1 // not_done
)

// V1 returns the v1 hc.InstanceState that corresponds to the State.
func (s State) V1() hc.InstanceState {
	return hc.InstanceState(s)
}

//go:generate go run "github.com/dmarkham/enumer" -type Action -linecomment -text

// Action is the action to perform on an intstance.
type Action int

// Action values.
const (
	NoneAction               Action = iota // none
	MigrateAction                          // migrate
	LiveMigrateAction                      // live-migrate
	ShutdownAction                         // shutdown
	ShutdownHardAction                     // shutdown-hard
	UndeployAction                         // undeploy
	UndeployHardAction                     // undeploy-hard
	HoldAction                             // hold
	ReleaseAction                          // release
	StopAction                             // stop
	SuspendAction                          // suspend
	ResumeAction                           // resume
	BootAction                             // boot
	DeleteAction                           // delete
	DeleteRecreateAction                   // delete-recreate
	RebootAction                           // reboot
	RebootHardAction                       // reboot-hard
	RescheduleAction                       // resched
	UnrescheduleAction                     // unresched
	PowerOffAction                         // poweroff
	PowerOffHardAction                     // poweroff-hard
	DiskAttachAction                       // disk-attach
	DiskDetachAction                       // disk-detach
	NicAttachAction                        // nic-attach
	NicDetachAction                        // nic-detach
	DiskSnapshotCreateAction               // disk-snapshot-create
	DiskSnapshotDeleteAction               // disk-snapshot-delete
	TerminateAction                        // terminate
	TerminateHardAction                    // terminate-hard
	DiskResizeAction                       // disk-resize
	DeployAction                           // deploy
	ChownAction                            // chown
	ChmodAction                            // chmod
	UpdateconfAction                       // updateconf
	RenameAction                           // rename
	ResizeAction                           // resize
	UpdateAction                           // update
	SnapshotCreateAction                   // snapshot-create
	SnapshotDeleteAction                   // snapshot-delete
	SnapshotRevertAction                   // snapshot-revert
	DiskSaveasAction                       // disk-saveas
	DiskSnapshotRevertAction               // disk-snapshot-revert
	RecoverAction                          // recover
	RetryAction                            // retry
	MonitorAction                          // monitor
	DiskSnapshotRenameAction               // disk-snapshot-rename
	AliasAttachAction                      // alias-attach
	AliasDetachAction                      // alias-detach
	PoffMigrateAction                      // poff-migrate
	PoffHardMigrateAction                  // poff-hard-migrate
)

// Template is the API payload based on the legacy xmlrpc backend.
type Template struct {
	Values                    map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	AutomaticDSRequirements   string            `json:"automatic_dsrequirements" yaml:"automatic_dsrequirements"`
	AutomaticNICRequirements  string            `json:"automatic_nicrequirements" yaml:"automatic_nicrequirements"`
	AutomaticRequirements     string            `json:"automatic_requirements" yaml:"automatic_requirements"`
	CPU                       float32           `json:"cpu" yaml:"cpu"`
	CPUCost                   string            `json:"cpu_cost" yaml:"cpu_cost"`
	CPUModel                  CPUModel          `json:"cpu_model" yaml:"cpu_model" xml:"CPU_MODEL"`
	CloningTemplateID         string            `json:"cloning_template_id" yaml:"cloning_template_id"`
	Context                   Context           `json:"context" yaml:"context"`
	Disk                      []Disk            `json:"disk,omitempty" yaml:"disk,omitempty"`
	DiskCost                  string            `json:"disk_cost" yaml:"disk_cost"`
	Emulator                  string            `json:"emulator" yaml:"emulator"`
	Features                  string            `json:"features" yaml:"features"`
	Graphics                  Graphics          `json:"graphics" yaml:"graphics"`
	HypervOptions             string            `json:"hyperv_options" yaml:"hyperv_options"`
	Imported                  string            `json:"imported" yaml:"imported"`
	Input                     string            `json:"input" yaml:"input"`
	InstanceGroup             []string          `json:"instance_group,omitempty" yaml:"instance_group,omitempty"`
	InstanceID                int               `json:"instance_id" yaml:"instance_id"`
	Memory                    int               `json:"memory" yaml:"memory"`
	MemoryCost                string            `json:"memory_cost" yaml:"memory_cost"`
	MemoryMax                 string            `json:"memory_max" yaml:"memory_max"`
	MemoryResizeMode          string            `json:"memory_resize_mode" yaml:"memory_resize_mode"`
	MemorySlots               string            `json:"memory_slots" yaml:"memory_slots"`
	NIC                       []NIC             `json:"nic,omitempty" yaml:"nic,omitempty"`
	NICAlias                  []NICAlias        `json:"nic_alias,omitempty" yaml:"nic_alias,omitempty"`
	NICDefault                string            `json:"nic_default" yaml:"nic_default"`
	NumaNode                  string            `json:"numa_node" yaml:"numa_node"`
	OS                        OS                `json:"os" yaml:"os"`
	PCI                       string            `json:"pci" yaml:"pci"`
	Raw                       string            `json:"raw" yaml:"raw"`
	SchedAction               []SchedAction     `json:"sched_action,omitempty" yaml:"sched_action,omitempty"`
	SecurityGroupRule         []string          `json:"security_group_rule,omitempty" yaml:"security_group_rule,omitempty"`
	Snapshot                  []Snapshot        `json:"snapshot,omitempty" yaml:"snapshot,omitempty"`
	SpiceOptions              string            `json:"spice_options" yaml:"spice_options"`
	SubmitOnHold              string            `json:"submit_on_hold" yaml:"submit_on_hold"`
	TemplateID                int               `json:"template_id" yaml:"template_id"`
	TmMADSystem               string            `json:"tm_madsystem" yaml:"tm_madsystem"`
	Topology                  string            `json:"topology" yaml:"topology"`
	VCPU                      int               `json:"vcpu" yaml:"vcpu"`
	VCPUMax                   string            `json:"vcpu_max" yaml:"vcpu_max"`
	VRouterID                 string            `json:"vrouter_id" yaml:"vrouter_id"`
	VRouterKeepalivedID       string            `json:"vrouter_keepalived_id" yaml:"vrouter_keepalived_id"`
	VRouterKeepalivedPassword string            `json:"vrouter_keepalived_password" yaml:"vrouter_keepalived_password"`
}

// CPUModel is the API payload based on the legacy xmlrpc backend.
type CPUModel struct {
	Values map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	Model  string            `json:"model" yaml:"model" xml:"MODEL"`
}

// Context is the API payload based on the legacy xmlrpc backend.
type Context struct {
	Values       map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	DiskID       int               `json:"disk_id" yaml:"disk_id"`
	Network      bool              `json:"network" yaml:"network"`
	SSHPublicKey string            `json:"ssh_public_key,omitempty" yaml:"ssh_public_key,omitempty"`
	Target       string            `json:"target" yaml:"target"`
}

// OS is the API payload based on the legacy xmlrpc backend.
type OS struct {
	Values map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	UUID   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Arch   string            `json:"arch,omitempty" yaml:"arch,omitempty"`
	Boot   string            `json:"boot,omitempty" yaml:"boot,omitempty"`
}

// UserTemplate is the API payload based on the legacy xmlrpc backend.
type UserTemplate struct {
	Values                map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	EFI                   string            `json:"efi,omitempty" yaml:"efi,omitempty"`
	VCenterCCRRef         string            `json:"vcenter_ccrref,omitempty" yaml:"vcenter_ccrref,omitempty"`
	VCenterDSRef          string            `json:"vcenter_dsref,omitempty" yaml:"vcenter_dsref,omitempty"`
	VCenterInstanceID     string            `json:"vcenter_instance_id,omitempty" yaml:"vcenter_instance_id,omitempty"`
	Error                 string            `json:"error,omitempty" yaml:"error,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty" yaml:"hypervisor,omitempty"`
	Info                  string            `json:"info,omitempty" yaml:"info,omitempty"`
	Logo                  string            `json:"logo,omitempty" yaml:"logo,omitempty"`
	LXDSecurityPrivileged bool              `json:"lxdsecurity_privileged,omitempty" yaml:"lxdsecurity_privileged,omitempty"`
	SchedRequirements     string            `json:"sched_requirements,omitempty" yaml:"sched_requirements,omitempty"`
	SnapshotSchedule      string            `json:"snapshot_schedule,omitempty" yaml:"snapshot_schedule,omitempty"`
}

// History is the API payload based on the legacy xmlrpc backend.
type History struct {
	ID                 int       `json:"id" yaml:"id"`
	SequenceNumber     int       `json:"sequence_number" yaml:"sequence_number"`
	Hostname           string    `json:"hostname" yaml:"hostname"`
	HostID             int       `json:"host_id" yaml:"host_id"`
	ClusterID          int       `json:"cluster_id" yaml:"cluster_id"`
	StartTime          time.Time `json:"start_time,omitempty" yaml:"start_time,omitempty"`
	EndTime            time.Time `json:"end_time,omitempty" yaml:"end_time,omitempty"`
	InstanceMAD        string    `json:"instance_mad" yaml:"instance_mad"`
	TransferManagerMAD string    `json:"transfer_manager_mad" yaml:"transfer_manager_mad"`
	DatastoreID        int       `json:"datastore_id" yaml:"datastore_id"`
	PrologStartTime    time.Time `json:"prolog_start_time,omitempty" yaml:"prolog_start_time,omitempty"`
	PrologEndtime      time.Time `json:"prolog_end_time,omitempty" yaml:"prolog_end_time,omitempty"`
	RunningStartTime   time.Time `json:"running_start_time,omitempty" yaml:"running_start_time,omitempty"`
	RunningEndTime     time.Time `json:"running_end_time,omitempty" yaml:"running_end_time,omitempty"`
	EpilogStartTime    time.Time `json:"epilog_start_time,omitempty" yaml:"epilog_start_time,omitempty"`
	EpilogEndTime      time.Time `json:"epilog_end_time,omitempty" yaml:"epilog_end_time,omitempty"`
	Action             Action    `json:"action,omitempty" yaml:"action,omitempty"`
	UserID             int       `json:"user_id" yaml:"user_id"`
	GroupID            int       `json:"group_id" yaml:"group_id"`
	RequestID          int       `json:"request_id" yaml:"request_id"`
}

// Monitoring is the API payload based on the legacy xmlrpc backend.
type Monitoring struct {
	CPU                             float64    `json:"cpu" yaml:"cpu"`
	DiskReadBytes                   int        `json:"disk_read_bytes" yaml:"disk_read_bytes"`
	DiskReadIOps                    int        `json:"disk_read_iops" yaml:"disk_read_iops"`
	DiskWriteBytes                  int        `json:"disk_write_bytes" yaml:"disk_write_bytes"`
	DiskWriteIOps                   int        `json:"disk_write_iops" yaml:"disk_write_iops"`
	DiskSize                        []DiskSize `json:"disk_size,omitempty" yaml:"disk_size,omitempty"`
	ID                              int        `json:"id" yaml:"id"`
	Memory                          int        `json:"memory" yaml:"memory"`
	NetRX                           int        `json:"netrx" yaml:"netrx"`
	NetTX                           int        `json:"nettx" yaml:"nettx"`
	Timestamp                       time.Time  `json:"timestamp" yaml:"timestamp"`
	VCenterESXHost                  string     `json:"vcenter_esxhost,omitempty" yaml:"vcenter_esxhost,omitempty"`
	VCenterGuestState               string     `json:"vcenter_guest_state,omitempty" yaml:"vcenter_guest_state,omitempty"`
	VCenterRPName                   string     `json:"vcenter_rpname,omitempty" yaml:"vcenter_rpname,omitempty"`
	VCenterVMWareToolsRunningStatus string     `json:"vcenter_vmware_tools_running_status,omitempty" yaml:"vcenter_vmware_tools_running_status,omitempty"`
	VCenterVMWareToolsVersion       string     `json:"vcenter_vmware_tools_version,omitempty" yaml:"vcenter_vmware_tools_version,omitempty"`
	VCenterVMWareToolsVersionStatus string     `json:"vcenter_vmware_tools_version_status,omitempty" yaml:"vcenter_vmware_tools_version_status,omitempty"`
	VCenterVMName                   string     `json:"vcenter_vmname,omitempty" yaml:"vcenter_vmname,omitempty"`
}

// DiskSize is the API payload based on the legacy xmlrpc backend.
type DiskSize struct {
	ID   int `json:"id" yaml:"id"`
	Size int `json:"size" yaml:"size"`
}

// Snapshots is the API payload based on the legacy xmlrpc backend.
type Snapshots struct {
	AllowOrphans string          `json:"allow_orphans" yaml:"allow_orphans"`
	CurrentBase  int             `json:"current_base" yaml:"current_base"`
	NextSnapshot int             `json:"next_snapshot" yaml:"next_snapshot"`
	Snapshot     []ImageSnapshot `json:"snapshot,omitempty" yaml:"snapshot,omitempty"`
}

// DiskSnapshots is a Snapshots for a specific disk.
type DiskSnapshots struct {
	Snapshots
	DiskID int `json:"disk_id" yaml:"disk_id"`
}

// Disk is the API payload based on the legacy xmlrpc backend.
type Disk struct {
	Values                map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	Format                string            `json:"format" yaml:"format"`
	AllowOrphans          string            `json:"allow_orphans" yaml:"allow_orphans"`
	Clone                 bool              `json:"clone" yaml:"clone"`
	CloneTarget           string            `json:"clone_target" yaml:"clone_target"`
	ClusterID             int               `json:"cluster_id" yaml:"cluster_id"`
	Datastore             string            `json:"datastore" yaml:"datastore"`
	DatastoreID           int               `json:"datastore_id" yaml:"datastore_id"`
	DevPrefix             string            `json:"dev_prefix" yaml:"dev_prefix"`
	DiskID                int               `json:"disk_id" yaml:"disk_id"`
	DiskSnapshotTotalSize int               `json:"disk_snapshot_total_size" yaml:"disk_snapshot_total_size"`
	DiskType              string            `json:"disk_type" yaml:"disk_type"`
	Driver                string            `json:"driver" yaml:"driver"`
	Image                 string            `json:"image" yaml:"image"`
	ImageID               int               `json:"image_id" yaml:"image_id"`
	ImageState            int               `json:"image_state" yaml:"image_state"`
	LnTarget              string            `json:"ln_target" yaml:"ln_target"`
	OriginalSize          int               `json:"original_size" yaml:"original_size"`
	Persistent            bool              `json:"persistent" yaml:"persistent"`
	PoolName              string            `json:"pool_name" yaml:"pool_name"`
	Readonly              bool              `json:"readonly" yaml:"readonly"`
	Save                  bool              `json:"save" yaml:"save"`
	Size                  int               `json:"size" yaml:"size"`
	Source                string            `json:"source" yaml:"source"`
	Target                string            `json:"target" yaml:"target"`
	TmMAD                 string            `json:"tm_mad" yaml:"tm_mad"`
	TmMADSystem           string            `json:"tm_madsystem" yaml:"tm_madsystem"`
	Type                  string            `json:"type" yaml:"type"`
	VCenterDSRef          string            `json:"vcenter_dsref,omitempty" yaml:"vcenter_dsref,omitempty"`
	VCenterInstanceID     string            `json:"vcenter_instance_id,omitempty" yaml:"vcenter_instance_id,omitempty"`
}

// Graphics is the API payload based on the legacy xmlrpc backend.
type Graphics struct {
	Values       map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	Listen       string            `json:"listen" yaml:"listen"`
	Port         int               `json:"port" yaml:"port"`
	GraphicsType string            `json:"type" yaml:"type"`
	Password     string            `json:"password" yaml:"password"`
	Keymap       string            `json:"keymap" yaml:"keymap" xml:"KEYMAP"`
}

// NIC is the API payload based on the legacy xmlrpc backend.
type NIC struct {
	Values               map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	ID                   int               `json:"id" yaml:"id"`
	NetworkID            int               `json:"network_id" yaml:"network_id"`
	IP                   string            `json:"ip" yaml:"ip"`
	MAC                  string            `json:"mac" yaml:"mac"`
	Model                string            `json:"model" yaml:"model"`
	VirtioQueues         string            `json:"virtio_queues" yaml:"virtio_queues"`
	Phydev               string            `json:"phy_dev" yaml:"phy_dev"`
	SecurityGroups       string            `json:"security_groups" yaml:"security_groups"`
	VCenterInstanceID    string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	VCenterNetRef        string            `json:"vcenter_net_ref" yaml:"vcenter_net_ref"`
	VCenterPortgroupType string            `json:"vcenter_portgroup_type" yaml:"vcenter_portgroup_type"`
}

// NICAlias is the API payload based on the legacy xmlrpc backend.
type NICAlias struct {
	Values               map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	AliasID              string            `json:"alias_id" yaml:"alias_id"`
	Parent               string            `json:"parent" yaml:"parent"`
	ParentID             string            `json:"parent_id" yaml:"parent_id"`
	VCenterInstanceID    string            `json:"vcenter_instance_id,omitempty" yaml:"vcenter_instance_id,omitempty"`
	VCenterNetRef        string            `json:"vcenter_net_ref,omitempty" yaml:"vcenter_net_ref,omitempty"`
	VCenterPortgroupType string            `json:"vcenter_portgroup_type,omitempty" yaml:"vcenter_portgroup_type,omitempty"`
}

// SchedAction is the API payload based on the legacy xmlrpc backend.
type SchedAction struct {
	Values   map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	Action   string            `json:"action" yaml:"action"`
	Args     string            `json:"args" yaml:"args"`
	Days     string            `json:"days" yaml:"days"`
	EndType  string            `json:"end_type" yaml:"end_type"`
	EndValue string            `json:"end_value" yaml:"end_value"`
	ID       string            `json:"id" yaml:"id"`
	Repeat   string            `json:"repeat" yaml:"repeat"`
	Time     string            `json:"time" yaml:"time"`
	Warning  string            `json:"warning" yaml:"warning"`
}

// Snapshot is the API payload based on the legacy xmlrpc backend.
type Snapshot struct {
	Values         map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	Action         string            `json:"action,omitempty" yaml:"action,omitempty"`
	Active         bool              `json:"active" yaml:"active"`
	HypervisorID   string            `json:"hypervisor_id" yaml:"hypervisor_id"`
	Name           string            `json:"name" yaml:"name"`
	SnapshotID     int               `json:"snapshot_id" yaml:"snapshot_id"`
	SystemDiskSize int               `json:"system_disk_size" yaml:"system_disk_size"`
	Time           time.Time         `json:"time,omitempty" yaml:"time,omitempty"`
}

// ImageSnapshot is the API payload based on the legacy xmlrpc backend.
type ImageSnapshot struct {
	Active   string `json:"active" yaml:"active"`
	Children string `json:"children" yaml:"children"`
	Date     int    `json:"date" yaml:"date"`
	ID       int    `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Parent   int    `json:"parent" yaml:"parent"`
	Size     int    `json:"size" yaml:"size"`
}

// GroupRole is the API payload based on the legacy xmlrpc backend.
type GroupRole struct {
	HostAffined     string `json:"host_affined" yaml:"host_affined"`
	HostAntiAffined string `json:"host_anti_affined" yaml:"host_anti_affined"`
	ID              int    `json:"id" yaml:"id"`
	Name            string `json:"name" yaml:"name"`
	Policy          string `json:"policy" yaml:"policy"`
}
