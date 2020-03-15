package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func trimLast(in string) (out string) {
	if len(in) == 0 {
		return
	}

	// Ignore the last byte and return the resulting string
	return in[:len(in)-1]
}

func getDirs() (dirs []string, err error) {
	var fis []os.FileInfo
	// Get the children of the current directory
	if fis, err = ioutil.ReadDir("./"); err != nil {
		return
	}

	// Iterate through children
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
	// Get the current working directory
	if cwd, err = os.Getwd(); err != nil {
		out.Errorf("error getting current working directory: %v", err)
		return
	}
	// Defer changing back to the current working directory after our execution is complete
	defer os.Chdir(cwd)

	// Change directory to the provided target directory
	if err = os.Chdir(dir); err != nil {
		out.Errorf("error switching to directory \"%s\": %v", dir, err)
		return
	}

	out.Successf("Switched to \"%s\"", dir)

	// Run provided func
	return fn()
}

func getUserInput() (input string, err error) {
	r := bufio.NewReader(os.Stdin)
	if input, err = r.ReadString('\n'); err != nil {
		return
	}

	input = trimLast(input)
	return
}

func requestPermission(destination string) (ok bool) {
	var (
		input string
		err   error
	)

	out.Warningf("Attempting to update branch \"%s\", are you sure?", destination)

	if input, err = getUserInput(); err != nil {
		out.Errorf("error getting user input: %v", err)
		return
	}

	// Convert to lowercase
	input = strings.ToLower(input)

	if input != "y" && input != "yes" {
		return
	}

	return true
}
