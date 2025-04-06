package rules

import(
	"fmt"
	"regexp"
)

// Rule is the standard format used across all scanners
type Rule struct {
	Name     string
	Category string
	Severity string
	Pattern  *regexp.Regexp
	Message  string
	Enabled  bool
}

// ReportIssue prints a scan issue in a standard format
func ReportIssue(file string, message string, severity string) {
	fmt.Printf("[%s] %s: %s\n", severity, file, message)
}
