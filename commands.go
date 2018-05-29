package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// GlobalFlags is cli flags
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "file, f",
		Value: "",
		Usage: "perse from file",
	},
}

// Action is main function
func Action(c *cli.Context) error {
	if c.String("f") != "" {
		_, e := ioutil.ReadFile(c.String("f"))
		if e != nil {
			fmt.Println(e)
			return nil
		}
	}

	result := perseText(c.Args()[0])

	for i := range result {
		fmt.Println(result[i])
	}

	fmt.Println(len(result))

	return nil
}

func perseText(text string) []string {
	a := strings.Split(text, " ")
	s := []string{}

	s = append(s, a[0])
	for i := range a {
		c := 0
		for j := range s {
			if a[i] == s[j] {
				break
			}
			c++

			if c == len(s) {
				s = append(s, a[i])
			}
		}
	}

	return s
}

// CommandNotFound is exception function
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
