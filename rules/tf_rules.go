package rules

import (
	"fmt"
	"os"
	"regexp"
)

type  tfRule struct {
	Name     string
	Category string
	Severity string
	Pattern  *regexp.Regexp
	Message  string
	Enabled  bool
}

var terraformRules = [] tfRule{
	{
		Name:     "Overly Permissive IAM Policy",
		Category: "Permissions",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)"Action":\s*"\*".*"Resource":\s*"\*"`),
		Message:  "IAM policy grants wildcard permissions (Action: * / Resource: *)",
		Enabled:  true,
	},
	{
		Name:     "Exposed Secret in DB Password",
		Category: "Secrets",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)password\s*=\s*"[^"]+"`),
		Message:  "Hardcoded password found in Terraform",
		Enabled:  true,
	},
	{
		Name:     "Open CIDR Block (0.0.0.0/0)",
		Category: "Network",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)cidr_blocks\s*=\s*\[.*0\.0\.0\.0\/0.*\]`),
		Message:  "Open CIDR block detected (0.0.0.0/0)",
		Enabled:  true,
	},
	// {
	// 	Name:     "Missing S3 Encryption",
	// 	Category: "Hardening",
	// 	Severity: "MEDIUM",
	// 	Pattern:  regexp.MustCompile(`(?i)resource\s+"aws_s3_bucket"[\s\S]+?{(?![\s\S]*server_side_encryption_configuration)`),
	// 	Message:  "S3 bucket is missing encryption configuration",
	// 	Enabled:  true,
	// },
	// {
	// 	Name:     "Missing S3 Versioning",
	// 	Category: "Backup",
	// 	Severity: "LOW",
	// 	Pattern:  regexp.MustCompile(`(?i)resource\s+"aws_s3_bucket"[\s\S]+?{(?![\s\S]*versioning\s+{)`),
	// 	Message:  "S3 bucket missing versioning block",
	// 	Enabled:  true,
	// },
	// {
	// 	Name:     "Missing Tags",
	// 	Category: "BestPractices",
	// 	Severity: "LOW",
	// 	Pattern:  regexp.MustCompile(`(?i)resource\s+\"aws_.+?\"\s+\".+?\"\s+{(?![\s\S]*tags\s+=)`),
	// 	Message:  "Resource is missing tags",
	// 	Enabled:  true,
	// },
	{
		Name:     "Default Credentials Usage",
		Category: "Secrets",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)username\s*=\s*"admin"|password\s*=\s*"admin"`),
		Message:  "Default credentials used in configuration",
		Enabled:  true,
	},
	{
		Name:     "Missing Logging",
		Category: "Observability",
		Severity: "MEDIUM",
		Pattern:  regexp.MustCompile(`(?i)cloudtrail|flow_log|log_group`),
		Message:  "Missing logging configuration for critical services",
		Enabled:  true,
	},
	{
		Name:     "Unused Variables",
		Category: "BestPractices",
		Severity: "LOW",
		Pattern:  regexp.MustCompile(`(?i)variable\s+"[a-zA-Z0-9_]+"`),
		Message:  "Declared variable might be unused",
		Enabled:  true,
	},
	{
		Name:     "Outdated Provider Version",
		Category: "Dependencies",
		Severity: "LOW",
		Pattern:  regexp.MustCompile(`(?i)provider\s+\"[a-zA-Z0-9_]+\"\s+{[\s\S]*?version\s*=\s*\"[0-9]+\.[0-9]+\.[0-9]+\"`),
		Message:  "Hardcoded or outdated provider version",
		Enabled:  true,
	},
	{
		Name:     "State File Insecure",
		Category: "Storage",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)backend\s+\"local\"`),
		Message:  "State file is stored locally and may not be secure",
		Enabled:  true,
	},
	{
		Name:     "Missing TLS Configuration",
		Category: "Hardening",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)protocol\s*=\s*"http"`),
		Message:  "TLS not enforced, 'http' protocol used",
		Enabled:  true,
	},
	{
		Name:     "Unnecessary Public Access",
		Category: "Network",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)acl\s*=\s*"public-read"|"public-read-write"`),
		Message:  "Public access to resource detected (e.g., S3)",
		Enabled:  true,
	},
	{
		Name:     "Immutable Infrastructure Violation",
		Category: "BestPractices",
		Severity: "MEDIUM",
		Pattern:  regexp.MustCompile(`(?i)lifecycle\s+{[\s\S]*?create_before_destroy\s*=\s*false`),
		Message:  "Resource is not using immutable infrastructure practices",
		Enabled:  true,
	},
	{
		Name:     "Insufficient Resource Limits",
		Category: "Performance",
		Severity: "MEDIUM",
		Pattern:  regexp.MustCompile(`(?i)instance_type\s*=\s*"t2\.micro"`),
		Message:  "Low resource limits set that could cause outages",
		Enabled:  true,
	},
}

func CheckTerraform(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", file, err)
		return
	}
	content := string(data)

	for _, rule := range terraformRules {
		if rule.Enabled && rule.Pattern.MatchString(content) {
			ReportIssue(file, rule.Message, rule.Severity)
		}
	}
}
