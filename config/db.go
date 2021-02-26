package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Crud struct {
	Ddl struct {
		Create     string `yaml:"create"`
		Insert     string `yaml:"insert"`
		Select     string `yaml:"select"`
		SelectByPk string `yaml:"selectByPk"`
	} `yaml:"ddl"`
}

func LoadDdl(configFile string) Crud {
	log.Println("Parsing YAML file", configFile)

	if configFile == "" {
		log.Fatalln("Please provide yaml file ")
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Error reading YAML file: ", err)
	}

	var crud Crud
	err = yaml.Unmarshal(yamlFile, &crud)
	if err != nil {
		log.Fatalln("Error parsing YAML file: ", err)
	}

	log.Println("Result: ", crud.Ddl)
	return crud
}
