package compute

type Instance struct {
	TypeName           string            `yaml:"typeName" hcl:",label"`
	Type               string            `yaml:"type" hcl:",label"`
	Name               string            `yaml:"name,omitempty" hcl:"name,omitempty"`
	MinCpuPlatform     string            `yaml:"minCpuPlatform,omitempty" hcl:"min_cpu_platform,omitempty"`
	CanIpForward       *bool             `yaml:"canIpForward,omitempty" hcl:"can_ip_forward,omitempty"`
	DeletionProtection *bool             `yaml:"deletionProtection" hcl:"deletion_protection,omitempty"`
	Hostname           *bool             `yaml:"hostname" hcl:"hostname,omitempty"`
	Labels             map[string]string `yaml:"labels" hcl:"labels,omitempty"`
	MachineType        *bool             `yaml:"machineType" hcl:"machine_type,omitempty"`
	Zone               map[string]string `yaml:"zone" hcl:"zone,omitempty"`

	ShieldedInstanceConfig ShieldedInstanceConfig `yaml:"shieldedInstanceConfig" hcl:"shielded_instance_config,block,omitempty"`
	GuestAccelerators      []AcceleratorConfig    `yaml:"guestAccelerators" hcl:"guest_accelerator,block,omitempty"`
	ServiceAccounts        []ServiceAccount       `yaml:"serviceAccounts" hcl:"service_account,block,omitempty"`
	Scheduling             []Scheduling           `yaml:"scheduling" hcl:"scheduling,block,omitempty"`
	MetadataMaps           Metadata               `yaml:"metadata" hcl:"metadata,omitempty"`
	NetworkInterfaces      []NetworkInterface     `yaml:"networkInterfaces" hcl:"network_interface,block,omitempty"`
	Tags                   Tags                   `yaml:"tags" hcl:"tags,omitempty"`
}

type NetworkInterface struct {
	Network    string `yaml:"network" hcl:"network"`
	Subnetwork string `yaml:"subnetwork" hcl:"subnetwork"`
}

type AcceleratorConfig struct {
	Count string `yaml:"count" hcl:"count,omitempty"`
	Type  string `yaml:"type" hcl:"type,omitempty"`
}

type ServiceAccount struct {
	Email  string   `yaml:"email" hcl:"email,omitempty"`
	Scopes []string `yaml:"scopes" hcl:"scopes,omitempty"`
}

type Scheduling struct {
	NodeAffinities    []SchedulingNodeAffinity `yaml:"nodeAffinities" hcl:"node_affinities,block,omitempty"`
	AutomaticRestart  *bool                    `yaml:"automaticRestart" hcl:"automatic_restart"`
	OnHostMaintenance string                   `yaml:"onHostMaintenance" hcl:"on_host_maintenance,omitempty"`
	Preemptible       *string                  `yaml:"preemptible" hcl:"preemptible,omitempty"`
}

type ShieldedInstanceConfig struct {
	EnableIntegrityMonitoring string `yaml:"enableIntegrityMonitoring" hcl:"enable_integrity_monitoring,omitempty"`
	EnableSecureBoot          string `yaml:"enableSecureBoot" hcl:"enable_secure_boot,omitempty"`
	EnableVtpm                string `yaml:"enableVtpm" hcl:"enable_vtpm,omitempty"`
}

type SchedulingNodeAffinity struct {
	Key      string   `yaml:"key" hcl:"key,omitempty"`
	Operator string   `yaml:"operator" hcl:"operator,omitempty"`
	Values   []string `yaml:"values" hcl:"values,omitempty"`
}

type Metadata map[string]string
type Tags []string

func (i *Metadata) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s map[interface{}]interface{}
	metadata := Metadata{}

	if err := unmarshal(&s); err != nil {
		panic(err)
	}

	if s["items"] != nil {
		for _, v := range s["items"].([]interface{}) {
			v := v.(map[interface{}]interface{})
			key := v["key"].(string)
			value := v["value"].(string)
			metadata[key] = value
		}
	}

	*i = metadata
	return nil
}

func (i *Tags) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s map[interface{}]interface{}
	tags := Tags{}

	if err := unmarshal(&s); err != nil {
		panic(err)
	}

	if s["items"] != nil {
		for _, v := range s["items"].([]interface{}) {
			v := v.(string)
			tags = append(tags, v)
		}
	}

	*i = tags
	return nil
}
