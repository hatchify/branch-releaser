package main

func release(source, destination string) (err error) {
	var origin string
	// Get the branch set before the process started, will revert to this branch after
	if origin, err = getBranchName(); err != nil {
		out.Error("error getting branch name: %v", err)
		return
	}
	defer gitCheckout(origin)

	// Sync with origin
	if err = gitPull(); err != nil {
		out.Error("error getting branch name: %v", err)
		return
	}

	out.Success("Synced with origin")

	// Checkout the destination branch
	if err = gitCheckout(destination); err != nil {
		out.Error("error encountered while switching to branch \"%s\": %v", destination, err)
		return
	}

	out.Success("Switched to destination branch \"%s\"", destination)

	var updated bool
	// Merge the source branch INTO the destination branch
	if updated, err = gitMerge(source); err != nil {
		out.Error("error encountered while merging with branch \"%s\": %v", source, err)
		return
	}

	if !updated {
		// No update occurred, leave note and return
		out.Notification("Destination branch \"%s\" already up to date", destination)
		return
	}

	out.Notification("Destination branch \"%s\" already up to date", destination)

	// Push updated changes to origin
	if err = gitPush(); err != nil {
		out.Error("error encountered while pushing changes to origin: %v", err)
		return
	}

	out.Success("Pushed changes to origin")
	return
}
