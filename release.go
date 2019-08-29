package main

import "fmt"

func release(source, destination string) (err error) {
	var origin string
	// Get the branch set before the process started, will revert to this branch after
	if origin, err = getBranchName(); err != nil {
		err = fmt.Errorf("error getting branch name: %v", err)
		return
	}
	defer gitCheckout(origin)

	// Sync with origin
	if err = gitPull(); err != nil {
		err = fmt.Errorf("error getting branch name: %v", err)
		return
	}

	out.Success("Synced with origin")

	// Checkout the destination branch
	if err = gitCheckout(destination); err != nil {
		err = fmt.Errorf("error encountered while switching to branch \"%s\": %v", destination, err)
		return
	}

	out.Success("Switched to destination branch \"%s\"", destination)

	var updated bool
	// Merge the source branch INTO the destination branch
	if updated, err = gitMerge(source); err != nil {
		err = fmt.Errorf("error encountered while merging with branch \"%s\": %v", source, err)
		return
	}

	if !updated {
		// No update occurred, leave note and return
		out.Notification("Destination branch \"%s\" already up to date", destination)
		return
	}

	out.Success("Destination branch \"%s\" synced with source branch \"%s\"", destination, source)

	// Push updated changes to origin
	if err = gitPush(); err != nil {
		err = fmt.Errorf("error encountered while pushing changes to origin: %v", err)
		return
	}

	out.Success("Pushed changes to origin")
	return
}
