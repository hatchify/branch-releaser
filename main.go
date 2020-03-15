package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/hatchify/scribe"
	"github.com/hatchify/stringset"
)

var out = scribe.New("Branch releaser")

func main() {
	var (
		// Source branch where data will be merged from
		source string
		// Destination branch where data will be merged into
		destination string
		// Whether or not the current request is for the current directory or its children
		recursive bool
		// How the error should be handled
		onError string

		err error
	)

	flag.StringVar(&source, "source", "", "Source branch to merge from")
	flag.StringVar(&destination, "destination", "", "Destination branch to merge into")
	flag.BoolVar(&recursive, "r", false, "Whether or not the current request is for the current directory or its children")
	flag.StringVar(&onError, "onError", "exit", "How to handle errors (e.g. exit, warn, ignore)")
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
		out.Errorf("error getting current working directory: %v", err)
		return
	}
	// Ensure we switch to the current working directory after our releases have completed
	defer os.Chdir(cwd)

	successful := stringset.New()
	errored := stringset.New()

	out.Notificationf("Beginning recursive release for the children of \"%s\"", cwd)

	var dirs []string
	// Get the child directories for the current directory
	if dirs, err = getDirs(); err != nil {
		out.Errorf("error getting target directories: %v", err)
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
		}); err == nil {
			successful.Set(dir)
			continue
		}

		switch onError {
		case "exit":
			out.Error(err.Error())
			return

		case "warn":
			out.Warning(err.Error())
		case "ignore":
		}

		err = nil
		errored.Set(dir)
	}

	successResults := successful.Slice()
	sort.Slice(successResults, func(a, b int) (less bool) {
		return successResults[a] < successResults[b]
	})

	outBuf := bytes.NewBuffer(nil)
	outBuf.WriteString("Successful:\n")

	for _, result := range successResults {
		line := fmt.Sprintf("\t- %s\n", result)
		outBuf.WriteString(line)
	}

	fmt.Print("\n")
	out.Success(outBuf.String())

	errorResults := errored.Slice()
	sort.Slice(errorResults, func(a, b int) (less bool) {
		return errorResults[a] < errorResults[b]
	})

	if len(errorResults) > 0 {
		errBuf := bytes.NewBuffer(nil)
		errBuf.WriteString("Errored:\n")

		for _, result := range errorResults {
			line := fmt.Sprintf("\t- %s\n", result)
			errBuf.WriteString(line)
		}

		fmt.Print("\n")
		out.Error(errBuf.String())
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
