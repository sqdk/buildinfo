package buildinfo

import (
	"fmt"
)

const (
	App       = "NOT_DEFINED"
	Build     = "DEVELOPMENT"
	Buildtime = "NOT_DEFINED"
	Branch    = "DEVELOPMENT"
	Commit    = "0"
)

func init() {
	fmt.Printf("Starting %v v. %v - %v (%v) T: %v", app, build, branch, commit, buildtime)
}

func GetAppVersion() string {
	return app
}

func GetBuildId() string {
	return build
}

func GetBuildTime() string {
	return buildtime
}

func GetBranch() string {
	return branch
}

func GetCommitId() string {
	return commit
}

func PrintDetails() {
	fmt.Printf("Starting %v v. %v - %v (%v) T: %v", app, build, branch, commit, buildtime)
	return
}
