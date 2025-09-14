package version

import (
	"strings"

	"github.com/sagan/ptool/constants"
)

var (
	VersionSuffix = "DEV-cancelpt" // e.g. DEV
	VersionTag    = "v0.1.11"
	Version       = ""
	Commit        = constants.NONE
	Date          = "unknown"
)

func init() {
	if Version == "" {
		if VersionSuffix == "" {
			Version = VersionTag
		} else {
			Version = VersionTag + "-" + VersionSuffix
		}
	} else if !strings.HasPrefix(Version, "v") {
		Version = "v" + Version
	}
}
