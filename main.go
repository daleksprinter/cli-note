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

var (
	c config
)

func load() error {
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	return nil
}

func New(context *cli.Context) error {
	fmt.Println(context.Args())
	return nil
}

func List(context *cli.Context) error {
	return nil
}

func View(context *cli.Context) error {
	return nil
}

func Edit(context *cli.Context) error {
	return nil
}

func Delete(context *cli.Context) error {
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
	load()
	fmt.Println(c)

	app := cli.NewApp()
	app.Name = "cli-note"
	app.Usage = "Usage: note {command} {args...}"
	app.Version = "0.1"
	app.Commands = commands

	app.Run(os.Args)

}
