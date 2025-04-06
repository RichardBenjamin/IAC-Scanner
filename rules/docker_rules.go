package rules

import(
	"fmt"
	"os"
	"regexp"
)

type DockerRule struct {
	Name 		string
	Category 	string
	Severity 	string
	Pattern 	*regexp.Regexp
	Message 	string
	Enabled 	bool
}

var dockerRules = []DockerRule{
{
	Name: 		"Using latest Tag",
	Category: 	"ImageManagement",
	Severity: 	"LOW",
	Pattern: 	regexp.MustCompile(`(?i)FROM\s+.*:latest`),
	Message: 	"Avoid using 'latest' tag in base images" ,
	Enabled: 	true,
},
{
	Name:		"Exposed Sensitive Data",
	Category:	"Secrets",
	Severity:	"HIGH",
	Pattern:	regexp.MustCompile(`(?i)ENV\s+.*(KEY|SECRET|PASSWORD)=.+`),
	Message:	"Sensitive data exposed in ENV variable",
	Enabled:	true,
},
{
	Name:		"Missing USER Instruction",
	Category:	"Permissions",
	Severity:	"MEDIUM",
	Pattern:	regexp.MustCompile(`(?i)^USER\s+`),
	Message:	"No USER instruction found - container may run as root",
	Enabled:	true,
},
{
	Name:     "Unnecessary Packages Installed",
	Category: "Optimization",
	Severity: "LOW",
	Pattern:  regexp.MustCompile(`(?i)apt-get install`),
	Message:  "Avoid installing unnecessary packages in image",
	Enabled:  true,
},
{
	Name:     "Missing Healthcheck",
	Category: "Observability",
	Severity: "LOW",
	Pattern:  regexp.MustCompile(`(?i)^HEALTHCHECK`),
	Message:  "No HEALTHCHECK instruction found",
	Enabled:  false,
},
{
	Name:     "Using ADD Instead of COPY",
	Category: "BestPractices",
	Severity: "LOW",
	Pattern:  regexp.MustCompile(`(?i)^ADD\s+`),
	Message:  "Prefer COPY over ADD unless needed",
	Enabled:  true,
},
{
	Name:     "Missing .dockerignore",
	Category: "Hardening",
	Severity: "LOW",
	Pattern:  regexp.MustCompile(`(?i)`),
	Message:  ".dockerignore file missing — check context manually",
	Enabled:  false,
},
{
	Name:     "No Multi-stage Builds",
	Category: "Optimization",
	Severity: "LOW",
	Pattern:  regexp.MustCompile(`(?i)FROM.*AS`),
	Message:  "Multi-stage builds not detected — consider using them to reduce image size",
	Enabled:  false,
},
{
	Name:     "Unpinned Package Versions",
	Category: "BestPractices",
	Severity: "MEDIUM",
	Pattern:  regexp.MustCompile(`(?i)apt-get install\s+[^\n\r]*$`),
	Message:  "Unpinned package versions found",
	Enabled:  true,
},
{
	Name:     "Shell Form of CMD or ENTRYPOINT",
	Category: "BestPractices",
	Severity: "LOW",
	Pattern:  regexp.MustCompile(`(?i)CMD\s+\[?.*\]?|ENTRYPOINT\s+\[?.*\]?`),
	Message:  "Consider using exec form of CMD or ENTRYPOINT",
	Enabled:  true,
},
}

func CheckDockerfile(file string) {
	data, err := os.ReadFile(file) 
		if err != nil {
			fmt.Printf("Error reading file %s: %v/n", file, err)
			return
		}
	
	content := string(data)
	for _, rule := range dockerRules {
		if rule.Enabled == true && rule.Pattern.MatchString(content){
			ReportIssue(file, rule.Severity, rule.Message)
		}
	}
}

