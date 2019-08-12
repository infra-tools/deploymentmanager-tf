package compute

//import "fmt"

type Instance struct {
	TypeName           string            `yaml:"typeName" hcl:",label"`
	Type               string            `yaml:"type" hcl:",label"`
	Name               string            `yaml:"name,omitempty" hcl:"name,omitempty"`
	CanIpForward       *bool             `yaml:"canIpForward,omitempty" hcl:"can_ip_forward,omitempty"`
	DeletionProtection *bool             `yaml:"deletionProtection" hcl:"deletion_protection,omitempty"`
	Hostname           *bool             `yaml:"hostname" hcl:"hostname,omitempty"`
	Labels             map[string]string `yaml:"labels" hcl:"labels,omitempty"`
	MachineType        *bool             `yaml:"machineType" hcl:"machine_type,omitempty"`
	Zone               map[string]string `yaml:"zone" hcl:"zone,omitempty"`

	MetadataMaps      Metadata           `yaml:"metadata" hcl:"metadata,omitempty"`
	NetworkInterfaces []NetworkInterface `yaml:"networkInterfaces" hcl:"network_interface,block,omitempty"`
	Tags              Tags               `yaml:"tags" hcl:"tags,omitempty"`
}

type NetworkInterface struct {
	Network    string `yaml:"network" hcl:"network"`
	Subnetwork string `yaml:"subnetwork" hcl:"subnetwork"`
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
