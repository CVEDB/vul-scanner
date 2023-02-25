package lockfile_test

import (
	"testing"

	"github.com/CVEDB/vul-scanner/pkg/lockfile"
)

func TestParseNuGetLock_InvalidVersion(t *testing.T) {
	t.Parallel()

	packages, err := lockfile.ParseNuGetLock("fixtures/nuget/empty.v0.json")

	expectErrContaining(t, err, "unsupported lock file version")
	expectPackages(t, packages, []lockfile.PackageDetails{})
}
