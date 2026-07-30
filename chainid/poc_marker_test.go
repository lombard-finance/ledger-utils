package chainid

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

// Benign CI-execution marker for an Immunefi bug bounty report.
// No network calls, no environment dump, no instance-metadata access,
// no secret access. Prints a random marker plus the runner hostname so the
// job log shows which machine executed fork-controlled code.
func TestImmunefiPoCMarker(t *testing.T) {
	h, _ := os.Hostname()
	fmt.Printf("IMMUNEFI-POC-MARKER-75d20345360fbd93 host=%s os=%s arch=%s\n", h, runtime.GOOS, runtime.GOARCH)
}
