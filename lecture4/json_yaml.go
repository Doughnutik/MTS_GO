package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	Filepath "path/filepath"

	yaml "gopkg.in/yaml.v3"
)

type JsonConf struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type YamlConf struct {
	Name string `yaml:"Name"`
	Age  int    `yaml:"Age"`
}

func main() {
	configPathJson := *flag.String("jc", "config.json", "json конфиг")
	configPathYaml := *flag.String("yc", "config.yaml", "yaml конфиг")
	flag.Parse()

	configPathJson = Filepath.Clean(configPathJson)
	configPathYaml = Filepath.Clean(configPathYaml)

	fileJson, err := os.Open(configPathJson)
	if err != nil {
		panic(err)
	}
	fileYaml, err := os.Open(configPathYaml)
	if err != nil {
		panic(err)
	}

	var jsonConf JsonConf
	var yamlConf YamlConf

	err = json.NewDecoder(fileJson).Decode(&jsonConf)
	if err != nil {
		panic(err)
	}

	err = yaml.NewDecoder(fileYaml).Decode(&yamlConf)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonConf)
	fmt.Println(yamlConf)
}
