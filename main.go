package main

import (
	"fmt"
	"io/ioutil"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclwrite"
	"github.com/juliosueiras/deploymentmanager-tf/schemas"
	"github.com/juliosueiras/deploymentmanager-tf/yamloader"
	"github.com/juliosueiras/deploymentmanager-tf/cobracmder"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"reflect"
	"strings"
)

type Top struct {
	ResouceList []Resource `yaml:"resources"`
}

type Resource struct {
	Name       string      `yaml:"name"`
	Type       string      `yaml:"type"`
	Properties interface{} `yaml:"properties"`
}

func failAndExit(e error) {
	fmt.Println(e)
	os.Exit(1)
}

func main() {
	rootCmd := cobracmder.Setup(loadFile)
	if err := rootCmd.Execute(); err != nil {
		failAndExit(err)
	}
}

func loadFile(file string, output string) {
	var loader yamloader.Yamloader
	loader.LoadFile(file)
	data := loader.GetFileContent()

	t := Top{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	f := hclwrite.NewEmptyFile()

	for _, v := range t.ResouceList {
		mainType := strings.Split(v.Type, ".")[0]
		tempType := schemas.BaseTypes[mainType][v.Type].Label

		structType := reflect.New(reflect.ValueOf(schemas.BaseTypes[mainType][v.Type].Type).Elem().Type()).Interface()

		properties := v.Properties
		properties.(map[interface{}]interface{})["type"] = tempType
		properties.(map[interface{}]interface{})["typeName"] = v.Name

		d, err := yaml.Marshal(&v.Properties)

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = yaml.Unmarshal(d, structType)

		if err != nil {
			log.Fatalf("error: %v", err)
		}
		test2 := gohcl.EncodeAsBlock(structType, "resource")

		f.Body().AppendBlock(test2)
	}

	err = ioutil.WriteFile(output, f.Bytes(), 0644)
	if err != nil {
		failAndExit(err)
	}

	fmt.Print(string(f.Bytes()))
}
