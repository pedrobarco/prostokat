package version

import "runtime"

var (
	version   = ""
	metadata  = ""
	gitCommit = ""
)

type BuildInfo struct {
	Version   string
	GitCommit string
	GoVersion string
}

func GetVersion() string {
	if metadata == "" {
		return version
	}
	return version + "+" + metadata
}

func Get() BuildInfo {
	return BuildInfo{
		Version:   GetVersion(),
		GitCommit: gitCommit,
		GoVersion: runtime.Version(),
	}
}
