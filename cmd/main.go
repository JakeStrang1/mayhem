package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JakeStrang1/mayhem"
	"github.com/JakeStrang1/mayhem/config"
	"gopkg.in/yaml.v2"
)

func main() {
	// Read flags
	configFlag := flag.String("config", "", "a mayhem config file")
	projectFlag := flag.String("project", "", "the path of the mayhem project to be generated")
	flag.Parse()

	// Get configs
	t := Config(*configFlag)

	// Set project path
	if t.ProjectPath == "" && *projectFlag == "" {
		fmt.Println("No project path specified. Generating default project path... myNetwork")
		t.ProjectPath = "myNetwork"
	} else if *projectFlag != "" {
		t.ProjectPath = *projectFlag
	}

	// Create project folder
	err := os.Mkdir(t.ProjectPath, 0755)
	check(err)

	err = mayhem.Generate(*t)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Config(filename string) *config.T {
	if filename == "" {
		filename = "mayhem.yaml"
		fmt.Println("No config file specified. Generating default config... " + filename)
		err := ioutil.WriteFile(filename, []byte(config.Default()), 0644)
		check(err)
	}

	data, err := ioutil.ReadFile(filename)
	check(err)

	t := config.T{}
	err = yaml.Unmarshal(data, &t)
	check(err)
	return &t
}
