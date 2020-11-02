package version

import (
	"testing"

	"github.com/ericcornelissen/gofunchecks/internal/utils"
)

func TestPrintVersion(t *testing.T) {
	p := utils.NoopLogger
	Print(p, "v3.1.4")
}
