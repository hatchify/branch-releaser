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
		out.Error("invalid destination branch, cannot be empty")
		return
	}

	warn := getMeta(".branch-releaser-warn")

	if warn.Has(destination) && !requestPermission(destination) {
		out.Warning("Aborting..")
		return
	}

	if !recursive {
		// Recursive flag isn't set, run release within the current directory then
		err = release(source, destination)
		return
	}

	var cwd string
	// Get the current working directory
	if cwd, err = os.Getwd(); err != nil {
		out.Error("error getting current working directory: %v", err)
		return
	}
	// Ensure we switch to the current working directory after our releases have completed
	defer os.Chdir(cwd)

	out.Notification("Beginning recursive release for the children of \"%s\"", cwd)

	var dirs []string
	// Get the child directories for the current directory
	if dirs, err = getDirs(); err != nil {
		out.Error("error getting target directories: %v", err)
		return
	}

	// Get ignored list
	ignored := getMeta(".branch-releaser-ignore")

	// Iterate through all the child directories
	for _, dir := range dirs {
		if ignored.Has(dir) {
			continue
		}

		// Execute release for current child
		if err = executeWithinDir(dir, func() (err error) {
			return release(source, destination)
		}); err != nil {
			return
		}
	}
}

func close(err error) {
	if err == nil {
		// No error exists, exit with a OK status code
		os.Exit(0)
	}

	// Error exists, exit with an error status code
	os.Exit(1)
}
