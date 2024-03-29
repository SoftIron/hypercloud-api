package cloud

import (
	"time"
)

// User is the API payload based on the legacy xmlrpc backend.
type User struct {
	ID                int               `json:"id" yaml:"id"`
	GroupID           int               `json:"group_id" yaml:"group_id"`
	Groups            []int             `json:"groups" yaml:"groups"`
	GroupName         string            `json:"group_name" yaml:"group_name"`
	Name              string            `json:"name" yaml:"name"`
	Password          string            `json:"password" yaml:"password"`
	AuthDriver        string            `json:"auth_driver" yaml:"auth_driver"`
	Enabled           bool              `json:"enabled" yaml:"enabled"`
	LoginToken        []UserLoginToken  `json:"login_token" yaml:"login_token"`
	Template          Template          `json:"template" yaml:"template"`
	DatastoreQuota    []UserDatastore   `json:"datastore_quota" yaml:"datastore_quota"`
	NetworkQuota      []UserNetwork     `json:"network_quota" yaml:"network_quota"`
	InstanceQuota     UserInstance      `json:"instance_quota" yaml:"instance_quota"`
	ImageQuota        []UserImage       `json:"image_quota" yaml:"image_quota"`
	DefaultUserQuotas UserDefaultQuotas `json:"default_user_quotas" yaml:"default_user_quotas"`
}

// UserTemplate is the API payload based on the legacy xmlrpc backend.
type UserTemplate struct {
	Values           map[string]string `json:"values" yaml:"values"`
	RadosGW          bool              `json:"rados_gw" yaml:"rados_gw"`
	RadosGWAccessKey string            `json:"rados_gw_access_key" yaml:"rados_gw_access_key"`
	RadosGWSecretKey string            `json:"rados_gw_secret_key" yaml:"rados_gw_secret_key"`
	SSHPublicKey     string            `json:"ssh_public_key" yaml:"ssh_public_key"`
	TokenPassword    string            `json:"token_password" yaml:"token_password"`
}

// UserLoginToken is the API payload based on the legacy xmlrpc backend.
type UserLoginToken struct {
	Token            string    `json:"token" yaml:"token"`
	ExpirationTime   time.Time `json:"expiration_time" yaml:"expiration_time"`
	EffectiveGroupID int       `json:"effective_group_id" yaml:"effective_group_id"`
}

// UserDatastore is the API payload based on the legacy xmlrpc backend.
type UserDatastore struct {
	ID         string `json:"id" yaml:"id"`
	Images     string `json:"images" yaml:"images"`
	ImagesUsed string `json:"images_used" yaml:"images_used"`
	Size       string `json:"size" yaml:"size"`
	SizeUsed   string `json:"size_used" yaml:"size_used"`
}

// UserNetwork is the API payload based on the legacy xmlrpc backend.
type UserNetwork struct {
	ID         string `json:"id" yaml:"id"`
	Leases     string `json:"leases" yaml:"leases"`
	LeasesUsed string `json:"leases_used" yaml:"leases_used"`
}

// UserInstance is the API payload based on the legacy xmlrpc backend.
type UserInstance struct {
	CPU                  float32 `json:"cpu" yaml:"cpu"`
	CPUUsed              float32 `json:"cpu_used" yaml:"cpu_used"`
	Memory               int     `json:"memory" yaml:"memory"`
	MemoryUsed           int     `json:"memory_used" yaml:"memory_used"`
	RunningCPU           float32 `json:"running_cpu" yaml:"running_cpu"`
	RunningCPUUsed       float32 `json:"running_cpu_used" yaml:"running_cpu_used"`
	RunningMemory        int     `json:"running_memory" yaml:"running_memory"`
	RunningMemoryUsed    int     `json:"running_memory_used" yaml:"running_memory_used"`
	RunningInstances     int     `json:"running_instances" yaml:"running_instances"`
	RunningInstancesUsed int     `json:"running_instances_used" yaml:"running_instances_used"`
	SystemDiskSize       int64   `json:"system_disk_size" yaml:"system_disk_size"`
	SystemDiskSizeUsed   int64   `json:"system_disk_size_used" yaml:"system_disk_size_used"`
	Instances            int     `json:"instances" yaml:"instances"`
	InstancesUsed        int     `json:"instances_used" yaml:"instances_used"`
}

// UserImage is the API payload based on the legacy xmlrpc backend.
type UserImage struct {
	ID                   string `json:"id" yaml:"id"`
	RunningInstances     string `json:"running_instances" yaml:"running_instances"`
	RunningInstancesUsed string `json:"running_instances_used" yaml:"running_instances_used"`
}

// UserDefaultQuotas is the API payload based on the legacy xmlrpc backend.
type UserDefaultQuotas struct {
	DatastoreQuota []UserDatastore `json:"datastore_quota" yaml:"datastore_quota"`
	NetworkQuota   []UserNetwork   `json:"network_quota" yaml:"network_quota"`
	InstanceQuota  UserInstance    `json:"instance_quota" yaml:"instance_quota"`
	ImageQuota     []UserImage     `json:"image_quota" yaml:"image_quota"`
}
