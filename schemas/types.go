package schemas

import (
	"github.com/juliosueiras/deploymentmanager-tf/schemas/base_types/compute"
)

type GenericResource struct {
	Type  interface{}
	Label string
}

// BaseCompute Dictionary for Deployment Base Compute Types
var BaseCompute = map[string]GenericResource{
	"compute.v1.instance": GenericResource{
		Type:  &compute.Instance{},
		Label: "google_compute_instance",
	},
}

// BaseTypes Dictionary for Deployment Base Types
var BaseTypes = map[string]map[string]GenericResource{
	"compute": BaseCompute,
}
