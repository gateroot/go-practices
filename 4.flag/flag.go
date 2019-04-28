package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Name    string `short:"n" long:"name" description:"A name" required:"true"`
	Args    struct {
		ID string
	} `positional-args:"yes" required:"yes"`
}

var opts Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[OPTIONS]"
	_, err := parser.Parse()
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	fmt.Printf("Name: %s\n", opts.Name)
	fmt.Printf("ID: %s\n", opts.Args.ID)
}
