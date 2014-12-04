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
	fmt.Printf("Starting %v v. %v - %v (%v) T: %v", App, Build, Branch, Commit, Buildtime)
}

func GetAppVersion() string {
	return App
}

func GetBuildId() string {
	return Build
}

func GetBuildTime() string {
	return Buildtime
}

func GetBranch() string {
	return Branch
}

func GetCommitId() string {
	return Commit
}

func PrintDetails() {
	fmt.Printf("Starting %v v. %v - %v (%v) T: %v", App, Build, Branch, Commit, Buildtime)
	return
}
