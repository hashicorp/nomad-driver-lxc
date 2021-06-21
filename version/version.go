package version

import (
	"fmt"
	"strings"
)

var (
	// GitCommit and GitDescribe are the commit that was compiled. These will
	// be filled in by the compiler.
	GitCommit   string
	GitDescribe string

	// Version must conform to the format expected by
	// github.com/hashicorp/go-version for tests to work.
	Version = "0.3.1"

	// VersionPrerelease is a pre-release marker for the version. If this is ""
	// (empty string) then it means that it is a final release. Otherwise, this
	// is a pre-release such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = "dev"

	// VersionMetadata is metadata further describing the build type.
	VersionMetadata = ""
)

// GetHumanVersion composes the parts of the version in a way that's suitable
// for displaying to humans.
func GetHumanVersion() string {
	version := Version
	release := VersionPrerelease

	if GitDescribe != "" {
		version = GitDescribe
	} else {
		if release == "" {
			release = "dev"
		}

		if release != "" && !strings.HasSuffix(version, "-"+release) {
			// if we tagged a prerelease version then the release is in the version
			// already.
			version += fmt.Sprintf("-%s", release)
		}

		if VersionMetadata != "" {
			version += fmt.Sprintf("+%s", VersionMetadata)
		}
	}

	// Add the commit hash at the very end of the version.
	if GitCommit != "" {
		version += fmt.Sprintf(" (%s)", GitCommit)
	}

	// Add v as prefix if not present
	if !strings.HasPrefix(version, "v") {
		version = fmt.Sprintf("v%s", version)
	}

	// Strip off any single quotes added by the git information.
	return strings.Replace(version, "'", "", -1)
}
