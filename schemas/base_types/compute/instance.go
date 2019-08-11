package compute

type Instance struct {
	Name         string `yaml:"name" hcl:",label"`
	Type         string `yaml:"type" hcl:",label"`
	CanIpForward bool   `yaml:"canIpForward" hcl:"can_ip_forward"`
}
