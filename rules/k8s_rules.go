package rules

import (
	"fmt"
	"os"
	"regexp"
)

type K8sRule struct {
	Name     string
	Category string
	Severity string
	Pattern  *regexp.Regexp
	Message  string
	Enabled  bool
}

var k8sRules = []K8sRule{
	{
		Name:     "Privileged Container",
		Category: "Permissions",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)privileged:\s*true`),
		Message:  "Privileged container detected",
		Enabled:  true,
	},
	{
		Name:     "Host Network Access",
		Category: "Network",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)hostNetwork:\s*true|hostPID:\s*true|hostIPC:\s*true`),
		Message:  "Pod uses host network, PID, or IPC namespace",
		Enabled:  true,
	},
	{
		Name:     "Hardcoded Secret in Env Var",
		Category: "Secrets",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)env:\s*-\s*name:.*value:\s*".*(key|secret|password).*"`),
		Message:  "Hardcoded secret found in environment variables",
		Enabled:  true,
	},
	{
		Name:     "Unrestricted Capabilities",
		Category: "Permissions",
		Severity: "MEDIUM",
		Pattern:  regexp.MustCompile(`(?i)capabilities:\s*\n\s*add:`),
		Message:  "Container adds extra Linux capabilities",
		Enabled:  true,
	},
	// {
	// 	Name:     "Missing Resource Limits",
	// 	Category: "Performance",
	// 	Severity: "MEDIUM",
	// 	Pattern:  regexp.MustCompile(`(?i)resources:\s*\n(?![\s\S]*limits:)`),
	// 	Message:  "Container missing resource limits",
	// 	Enabled:  true,
	// },
	{
		Name:     "Default Service Account Used",
		Category: "Permissions",
		Severity: "LOW",
		Pattern:  regexp.MustCompile(`(?i)serviceAccountName:\s*default`),
		Message:  "Pod is using the default service account",
		Enabled:  true,
	},
	{
		Name:     "Excessive RBAC Permissions",
		Category: "Permissions",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)rules:\s*-\s*apiGroups:.*\*.*resources:.*\*`),
		Message:  "RBAC role grants wildcard permissions",
		Enabled:  true,
	},
	{
		Name:     "Missing Network Policy",
		Category: "Network",
		Severity: "MEDIUM",
		Pattern:  regexp.MustCompile(`(?i)kind:\s*NetworkPolicy`),
		Message:  "No NetworkPolicy found in namespace",
		Enabled:  false,
	},
	{
		Name:     "Missing TLS in Ingress",
		Category: "Hardening",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)kind:\s*Ingress[\s\S]*?tls:\s*\[\]`),
		Message:  "Ingress lacks TLS configuration",
		Enabled:  true,
	},
	{
		Name:     "Writable Container Filesystem",
		Category: "Hardening",
		Severity: "MEDIUM",
		Pattern:  regexp.MustCompile(`(?i)readOnlyRootFilesystem:\s*false`),
		Message:  "Container filesystem is not read-only",
		Enabled:  true,
	},
	// {
	// 	Name:     "Missing Liveness/Readiness Probes",
	// 	Category: "Observability",
	// 	Severity: "LOW",
	// 	Pattern:  regexp.MustCompile(`(?i)containers:([\s\S]*?)(?!livenessProbe|readinessProbe)`),
	// 	Message:  "Missing liveness or readiness probes",
	// 	Enabled:  true,
	// },
	{
		Name:     "HostPath Volume Use",
		Category: "Storage",
		Severity: "HIGH",
		Pattern:  regexp.MustCompile(`(?i)hostPath:\s*path:`),
		Message:  "HostPath volume usage detected",
		Enabled:  true,
	},
	{
		Name:     "Missing Security Context",
		Category: "Hardening",
		Severity: "LOW",
		Pattern:  regexp.MustCompile(`(?i)securityContext:\s*{?\s*}?`),
		Message:  "Security context not properly defined",
		Enabled:  true,
	}
}

func CheckKubernetesYAML(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", file, err)
		return
	}
	content := string(data)

	for _, rule := range k8sRules {
		if rule.Enabled && rule.Pattern.MatchString(content) {
			ReportIssue(file, rule.Message, rule.Severity)
		}
	}
}
