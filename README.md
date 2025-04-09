# IAC-Scanner Documentation

## Overview

**IAC-Scanner** is a security analysis tool written in Golang that detects misconfigurations in Infrastructure-as-Code (IaC) files. It currently supports:

- **Dockerfiles**
- **Kubernetes YAML** files
- **Terraform** files

The scanner performs pattern-based static analysis to catch common mistakes and insecure practices during the development lifecycle.

---

## Features

- Static analysis for Docker, Kubernetes, and Terraform
- Severity-based classification: LOW, MEDIUM, HIGH
- Flexible rule engine with regex support
- CLI output and export to log file
- GitHub Actions CI/CD integration
- Scan results uploaded as pipeline artifacts
- Slack notifications with scan summaries
- Upload scan logs to AWS S3 and share object link to Slack

---

## Directory Structure

```
IAC-Scanner/
├── main.go
├── go.mod
├── scanner/
│   ├── scanner.go
│   ├── docker.go
│   ├── k8s.go
│   └── terraform.go
├── rules/
│   ├── docker_rules.go
│   ├── k8s_rules.go
│   └── terraform_rules.go
├── test-files/
│   ├── Dockerfile
│   ├── sample.yaml
│   └── main.tf
```

---

## Installation and Setup

### Prerequisites

- Go 1.22 or higher
- Git

### Installation Steps

```bash
git clone https://github.com/your-org/IAC-Scanner.git
cd IAC-Scanner
go build -o iac-scan main.go
```

### Usage Examples

#### Scan a Single File

```bash
./iac-scan path/to/file.yaml
```

#### Scan a Directory

```bash
./iac-scan ./test-files/
```

#### Output to File

```bash
./iac-scan ./test-files/ > scan-results.log
```

---

## GitHub Actions Integration

A sample GitHub Actions workflow is provided to:

- Set up the Go environment
- Build the scanner
- Run the scan on a target folder
- Upload the scan result as an artifact
- Upload the report file to an AWS S3 bucket
- Fail the job on HIGH severity issues
- Notify a Slack channel with a summary and link to the S3 object

Slack integration is configured using a webhook secret stored in repository secrets. AWS credentials for S3 access should also be configured as GitHub secrets.

---

## Defining Custom Rules

Each rule object includes:

- `Name`: Identifier of the rule
- `Category`: Category or domain (e.g., Hardening, Secrets)
- `Severity`: LOW, MEDIUM, HIGH
- `Pattern`: Go regex for scanning
- `Message`: Description to show
- `Enabled`: Boolean flag

### Example Rule (Docker)

```go
{
  Name: "Avoid Latest Tag",
  Category: "Best Practices",
  Severity: "LOW",
  Pattern: regexp.MustCompile(`(?i)FROM\s+.+:latest`),
  Message: "Avoid using 'latest' tag in base images",
  Enabled: true,
}
```

To define new rules, edit the appropriate file in the `/rules/` directory and recompile the scanner.

---

## Sample Output

```
Scanning Docker file: test-files/Dockerfile
[LOW] Avoid using 'latest' tag in base images
[HIGH] Sensitive data exposed in ENV variable

Scanning Kubernetes file: test-files/sample.yaml
[HIGH] Privileged container detected
[LOW] Security context not properly defined
```

---

## Artifacts and Reports

Scan results can be redirected to a file (e.g., `scan-results.log`) and uploaded as an artifact in your CI/CD workflow. This makes it easy to audit security issues post-deployment.

Artifacts are downloadable from the GitHub Actions summary tab. The report is also uploaded to a configured AWS S3 bucket, and a public or signed link to this report is included in the Slack notification.

---

## Roadmap

- [ ] Add JSON/HTML output format
- [ ] Expand Terraform rule coverage
- [ ] Web dashboard for visualization

---

## Contributing

- Fork the repo and clone it locally
- Group rules by file type
- Test with sample files in `/test-files/`
- Submit PRs with a clear explanation of changes

---

## License

MIT License © 2025 — Kenechukwu



