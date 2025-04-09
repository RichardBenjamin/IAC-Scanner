package rules

import(
	"fmt"
	"regexp"
	"os"
)


type Rule struct {
	Name     string
	Category string
	Severity string
	Pattern  *regexp.Regexp
	Message  string
	Enabled  bool
}

func ReportIssue(file string, message string, severity string) {
	fmt.Printf("[%s] %s: %s\n", severity, file, message)


	logFile, err := os.OpenFile("scan-results.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	logMsg := fmt.Sprintf("[%s] %s: %s\n", severity, file, message)
	logFile.WriteString(logMsg)
}

