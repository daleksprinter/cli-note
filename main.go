package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Dir    string `yaml:"dir"`
	Editor string `yaml:"editor"`
}

func load() (config, error) {
	buf, err := ioutil.ReadFile("./config.yaml")
	var c config
	if err != nil {
		return config{}, err
	}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return config{}, err
	}

	return c, err

}
func main() {
	fmt.Println(load())
}
