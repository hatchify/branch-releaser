package main

import (
	"bufio"
	"os"

	"github.com/hatchify/stringset"
)

func getMeta(filename string) (meta *stringset.StringSet) {
	var (
		f   *os.File
		err error
	)

	// Initialize new stringset
	meta = stringset.New()

	// Attempt to open meta file
	if f, err = os.Open(filename); err != nil {
		// Meta file does not exist. This file is not required, so we do not need to throw an error.
		return
	}
	// Defer the closing of the file
	defer f.Close()

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

		// Set current line contents as a meta key
		meta.Set(string(bs))
	}

	return
}
