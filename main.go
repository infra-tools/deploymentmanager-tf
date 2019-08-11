package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclwrite"
	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  label: test2
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string `hcl:"a"`
	B struct {
		Label    string `yaml:"label" hcl:"test,label"`
		RenamedC int    `yaml:"c" hcl:"c"`
		D        []int  `yaml:",flow" hcl:"d"`
	} `hcl:"b,block"`
}

func main() {
	t := T{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- d dump:\n%s\n\n", string(d))

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&t, f.Body())
	fmt.Printf("--- hcl dump:\n%s\n\n", string(f.Bytes()))
}
