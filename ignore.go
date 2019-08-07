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

	if f, err = os.Open(".branch-releaser-ignore"); err != nil {
		return
	}
	defer f.Close()

	ignored = stringset.New()
	scn := bufio.NewScanner(f)

	for scn.Scan() {
		bs := scn.Bytes()
		if len(bs) == 0 {
			continue
		}

		ignored.Set(string(bs))
	}

	return
}
