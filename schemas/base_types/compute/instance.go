package compute

//import "fmt"

type Instance struct {
	Name              string             `yaml:"name" hcl:",label"`
	Type              string             `yaml:"type" hcl:",label"`
	CanIpForward      bool               `yaml:"canIpForward" hcl:"can_ip_forward"`
	MachineType       bool               `yaml:"machineType" hcl:"machine_type"`
	MetadataMaps      Metadata           `yaml:"metadata" hcl:"metadata"`
	NetworkInterfaces []NetworkInterface `yaml:"networkInterfaces" hcl:"network_interface,block"`
}

type NetworkInterface struct {
	Network    string `yaml:"network" hcl:"network"`
	Subnetwork string `yaml:"subnetwork" hcl:"subnetwork"`
}

type Metadata map[string]string

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
