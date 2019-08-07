package main

import (
	"flag"
	"os"

	"github.com/hatchify/output"
)

var out = output.NewWrapper("Branch releaser")

func main() {
	var (
		// Source branch where data will be merged from
		source string
		// Destination branch where data will be merged into
		destination string
		// Whether or not the current request is for the current directory or its children
		recursive bool

		err error
	)

	flag.StringVar(&source, "source", "", "Source branch to merge from")
	flag.StringVar(&destination, "destination", "", "Destination branch to merge into")
	flag.BoolVar(&recursive, "r", false, "Whether or not the current request is for the current directory or its children")
	flag.Parse()
	defer close(err)

	if len(source) == 0 {
		out.Error("invalid source branch, cannot be empty")
		return
	}

	if len(destination) == 0 {
		out.Error("invalid source branch, cannot be empty")
		return
	}

	if err = release(source, destination); err != nil {
		return
	}
}

func close(err error) {
	if err == nil {
		os.Exit(0)
	}

	os.Exit(1)
}
