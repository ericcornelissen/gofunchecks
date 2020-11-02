package version

import (
	"fmt"

	"github.com/ericcornelissen/gofunchecks/internal/utils"
)

const versionMessageTemplate = "version %s"

// Print the version message for gofunchecks programs.
func Print(p utils.Printer, version string) {
	versionMessage := fmt.Sprintf(versionMessageTemplate, version)
	p.Print(versionMessage)
}
