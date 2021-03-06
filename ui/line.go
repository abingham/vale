package ui

import (
	"fmt"
	"strings"

	"github.com/ValeLint/vale/core"
)

// PrintLineAlerts prints Alerts in <path>:<line>:<col>:<check>:<message> format.
func PrintLineAlerts(linted []*core.File) bool {
	var base string

	alertCount := 0
	for _, f := range linted {
		// If vale is run from a parent directory of f, we use a shorter file
		// path -- e.g., if run from the directory 'vale', we use
		// 'testdata/test.cc: ...' instead of
		// /Users/.../.../.../vale/testdata/test.cc: ...'.
		if strings.Contains(f.Path, core.ExeDir) {
			base = strings.Split(f.Path, core.ExeDir)[1]
		} else {
			base = f.Path
		}

		for _, a := range f.SortedAlerts() {
			if a.Severity == "error" {
				alertCount++
			}
			a.Message = fixOutputSpacing(a.Message)
			fmt.Print(fmt.Sprintf("%s:%d:%d:%s:%s\n",
				base, a.Line, a.Span[0], a.Check, a.Message))
		}
	}
	return alertCount != 0
}
