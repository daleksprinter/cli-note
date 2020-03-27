package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
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

func New(context *cli.Context) error {
	fmt.Println(context.Args())
	return nil
}

var commands = []cli.Command{
	{
		Name:    "new",
		Aliases: []string{"n"},
		Usage:   "Add New Note",
		Action:  New,
	},
}

func main() {
	fmt.Println(load())

	app := cli.NewApp()
	app.Name = "cli-note"
	app.Usage = "Usage: note {command} {args...}"
	app.Version = "0.1"
	app.Commands = commands

	app.Run(os.Args)

}
