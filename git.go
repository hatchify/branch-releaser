package main

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/hatchify/errors"
)

func getBranchName() (name string, err error) {
	return gitCmd("symbolic-ref", "--short", "HEAD")
}

func gitPull() (err error) {
	_, err = gitCmd("pull")
	return
}

func gitCheckout(branch string) (err error) {
	_, err = gitCmd("checkout", branch)
	return

}

func gitMerge(branch string) (updated bool, err error) {
	var out string
	if out, err = gitCmd("merge", branch); err != nil {
		return
	}

	updated = strings.Index(out, "Already") == -1
	return
}

func gitFetch() (err error) {
	if _, err = gitCmd("fetch", "origin"); err != nil {
		return
	}

	return
}

func gitPush() (err error) {
	_, err = gitCmd("push")
	return
}

func gitCmd(command string, args ...string) (out string, err error) {
	// Initialize output buffer
	outBuf := bytes.NewBuffer(nil)
	// Initialize error buffer
	errBuf := bytes.NewBuffer(nil)

	// Prepend command to args list
	args = append([]string{command}, args...)

	// Initialize new git command
	cmd := exec.Command("git", args...)
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	// Run command
	if err = cmd.Run(); err != nil {
		// Set return error as the error response with the last character (newline) omitted
		err = errors.Error(trimLast(errBuf.String()))
		return
	}

	if out = outBuf.String(); len(out) == 0 {
		// Sometimes git CLI returns the response output in the stderr
		// If the output is empty (and we don't have an error response code), we can
		// safely assume the stdout message is actually within the stderr buffer
		out = errBuf.String()
	}

	out = trimLast(out)
	return
}
