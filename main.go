package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v2"
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
	confdir := fmt.Sprintf("%s/.note_conf.yaml", os.ExpandEnv("$HOME"))

	if !Exists(confdir) {
		log.Fatal(fmt.Sprintf("Configure file not found in %s", confdir))
	}

	buf, err := ioutil.ReadFile(confdir)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	return nil
}

func Exists(dirname string) bool {
	_, err := os.Stat(os.ExpandEnv(dirname))
	return err == nil
}

func GetNowTime() string {
	return time.Now().Format("2006-01-02_03:04:05")
}

func New(context *cli.Context) error {
	dir := os.ExpandEnv(c.Dir)
	if !Exists(dir) {
		err := os.Mkdir(dir, 0700)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	filename := fmt.Sprintf("%s/%s_%s", dir, GetNowTime(), context.String("name"))

	_, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cmd := exec.Command(c.Editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func List(context *cli.Context) error {
	ls := fmt.Sprintf("%s `ls -d %s/* | peco`", c.Editor, os.ExpandEnv(c.Dir))
	cmd := exec.Command("bash", "-c", ls)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Delete(context *cli.Context) error {
	rm := fmt.Sprintf("rm `ls -d  %s/* | peco`", os.ExpandEnv(c.Dir))
	cmd := exec.Command("bash", "-c", rm)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Last(context *cli.Context) error {
	return nil
}

func History(context *cli.Context) error {
	return nil
}

func BackUp(context *cli.Context) error {
	return nil
}

var commands = []*cli.Command{
	{
		Name:    "new",
		Aliases: []string{"n"},
		Usage:   "Add New Note",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "Note Title",
			},
		},
		Action: New,
	},
	{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List Notes",
		Action:  List,
	},
	{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "Delete Note",
		Action:  Delete,
	},
}

func main() {
	load()

	app := cli.NewApp()
	app.Name = "Command Line Note Application"
	app.Version = "0.1"
	app.Commands = commands

	app.Run(os.Args)

}
