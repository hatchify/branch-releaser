package main

import (
	"bufio"
	"os"

	"github.com/hatchify/stringset"
)

func getIgnored() (ignored *stringset.StringSet) {
	var (
		f   *os.File
		err error
	)

	// Attempt to open ignore file
	if f, err = os.Open(".branch-releaser-ignore"); err != nil {
		// Ignore file does not exist. This file is not required, so we do not need to throw an error.
		return
	}
	// Defer the closing of the file
	defer f.Close()

	// Initialize new stringset
	ignored = stringset.New()
	// Initialize new scanner
	scn := bufio.NewScanner(f)

	// Iterate through file lines
	for scn.Scan() {
		// Assign bytes to variable
		bs := scn.Bytes()

		// Check byte length
		if len(bs) == 0 {
			// Ignore empty lines, continue
			continue
		}

		// Set current line contents as a stringset key
		ignored.Set(string(bs))
	}

	return
}
