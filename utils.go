package main

import (
	"io/ioutil"
	"os"
)

func trimLast(in string) (out string) {
	return in[:len(in)-1]
}

func getDirs(recursive bool) (dirs []string, err error) {
	if !recursive {
		// Set target dir as the current one
		dirs = []string{"./"}
		return
	}

	var fis []os.FileInfo
	if fis, err = ioutil.ReadDir("./"); err != nil {
		return
	}

	for _, fi := range fis {
		// Check to see if entry is a directory
		if !fi.IsDir() {
			// This is not a directory, continue
			continue
		}

		// Check to see if the entry is hidden
		if fi.Name()[0] == '.' {
			// Ignore hidden directories, return
			continue
		}

		// Append entry name to directories list
		dirs = append(dirs, fi.Name())
	}

	return
}
