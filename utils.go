package main

import (
	"io/ioutil"
	"os"
)

func trimLast(in string) (out string) {
	return in[:len(in)-1]
}

func getDirs() (dirs []string, err error) {
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

func executeWithinDir(dir string, fn func() error) (err error) {
	var cwd string
	if cwd, err = os.Getwd(); err != nil {
		out.Error("error getting current working directory: %v", err)
		return
	}
	defer os.Chdir(cwd)

	if err = os.Chdir(dir); err != nil {
		out.Error("error switching to directory \"%s\": %v", dir, err)
		return
	}

	out.Success("Switched to \"%s\"", dir)

	return fn()
}
