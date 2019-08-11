package main

import (
	"fmt"
	"log"
	"strings"
	//
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclwrite"
	"github.com/juliosueiras/deploymentmanager-tf/schemas"
	"gopkg.in/yaml.v2"
	"reflect"
)

var data = `
resources:
- name: test-vm
  type: compute.v1.instance
  properties:
    networkInterfaces:
    - name: Interface 0
      network: https://www.googleapis.com/compute/v1/projects/test/global/networks/test-vpc
      subnetwork: https://www.googleapis.com/compute/v1/projects/test/regions/northamerica-northeast1/subnetworks/shared-services-subnet
    canIpForward: true
    metadata:
      items:
      - key: bitnami-base-password
        value: redacted
      - key: google-monitoring-enable
        value: '0'
      - key: google-logging-enable
        value: '0'
`

type Top struct {
	ResouceList []Resource `yaml:"resources"`
}

type Resource struct {
	Name       string      `yaml:"name"`
	Type       string      `yaml:"type"`
	Properties interface{} `yaml:"properties"`
}

func main() {
	t := Top{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	mainType := strings.Split(t.ResouceList[0].Type, ".")[0]
	tempType := schemas.BaseTypes[mainType][t.ResouceList[0].Type].Label

	structType := reflect.New(reflect.ValueOf(schemas.BaseTypes[mainType][t.ResouceList[0].Type].Type).Elem().Type()).Interface()

	properties := t.ResouceList[0].Properties
	properties.(map[interface{}]interface{})["type"] = tempType
	properties.(map[interface{}]interface{})["name"] = t.ResouceList[0].Name

	d, err := yaml.Marshal(&t.ResouceList[0].Properties)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err2 := yaml.Unmarshal(d, structType)

	if err2 != nil {
		log.Fatalf("error: %v", err2)
	}
	f := hclwrite.NewEmptyFile()
	test2 := gohcl.EncodeAsBlock(structType, "resource")

	f.Body().AppendBlock(test2)

	fmt.Print(string(f.Bytes()))
}
