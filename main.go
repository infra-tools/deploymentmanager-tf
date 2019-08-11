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
    canIpForward: true
`

type Top struct {
	ResouceList []Resource `yaml:"resources"`
}

type Resource struct {
	Name       string      `yaml:"name"`
	Type       string      `yaml:"type"`
	Properties interface{} `yaml:"properties"`
}

type Output struct {
	Resource interface{} `hcl:"resource,block"`
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

	yaml.Unmarshal(d, structType)

	f := hclwrite.NewEmptyFile()
	test2 := gohcl.EncodeAsBlock(structType, "resource")

	f.Body().AppendBlock(test2)

	fmt.Print(string(f.Bytes()))
}
