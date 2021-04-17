package version

var (
	// Version is the version filled in by the compiler
	Version string
	// GitCommit is the git commit filled in by the compiler
	GitCommit string
)

func GetVersion() string {
	return Version
}

func GetGitCommit() string {
	return GitCommit
}
