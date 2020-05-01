package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	ymlfile = "%s/dennis.yaml"
)

type Root struct {
	Tests Tests `yaml:"tests"`
}

type Tests []Test

type Test struct {
	Cmd   string     `yaml:"cmd"`
	Cases []Testcase `yaml:"cases"`
}

type Testcase struct {
	Name string `yaml:"name"`
	Args string `yaml:"args"`
	Out  int    `yaml:"out"`
}

func main() {
	runTests()
}

func runTests() {
	p, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	f := fmt.Sprintf(ymlfile, p)
	b, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	r := Root{}
	err = yaml.Unmarshal(b, &r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for _, t := range r.Tests {
		fmt.Printf("Running tests for: %s\n", t.Cmd)
		fmt.Println("==============================================================================================")
		for i, c := range t.Cases {
			fmt.Printf("%d - Test case: [%s]\n", i, c.Name)
			fmt.Println("----------------------------------------------------------------------------------------------")
			aa := strings.Split(c.Args, ",")

			cmd := exec.Command(t.Cmd, aa...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Start(); err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			out := 0
			if err := cmd.Wait(); err != nil {
				out = 1
			}

			if out != c.Out {
				err := fmt.Errorf("-> Failed, got %v, want %v", out, c.Out)
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			} else {
				fmt.Println("-> Passed")
			}
			fmt.Println("")
		}
		fmt.Println("=============================================END==============================================")
	}
}
