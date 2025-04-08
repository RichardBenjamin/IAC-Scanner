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
- Fail the job on HIGH severity issues

> File location: `.github/workflows/iac-scan.yml`

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
MIT License © 2025 — Your Name or Organization

































































































































































i have added a pipeline to test the files on github. i have added the the feature of sending the logs as artifact and notification link to a slack channel 























name: Infrastructure Security Scan

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  workflow_dispatch:
    runs-on: ubuntu-latest 



    env:
      SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK_URL }} 

    steps:

     - name: Send Slack alert on failure
       if: always()
       uses: slackapi/slack-github-action@v1.24.0
       with:
         payload: |
            {
              "text": " *IAC Scan Failed!*",
              "blocks": [
                {
                  "type": "section",
                  "text": {
                    "type": "mrkdwn",
                    "text": "*IAC Scan failed on repo:* `${{ github.repository }}`\n*Branch:* `${{ github.ref }}`\n*Commit:* `${{ github.sha }}`\n<${{ github.server_url }}/${{ github.repository }}/actions/runs/${{ github.run_id }}|View Logs>"
                  }
                }
              ]
            }


    #  - name: Slack notification
    #    uses: act10ns/slack@v2
    #    with:
    #       channel: '#all-devsecops'
    #       status: ${{ job.status }}
    #       steps: ${{ toJson(steps) }}
    #    if: always()

     - name:  Checkout code
       uses: actions/checkout@v3

     - name:  Set up Go
       uses: actions/setup-go@v4
       with:
        go-version: '1.22'

     - name:  Build IAC-Scanner
       run: |
        go build -o iac-scanner

     - name:  Run IAC-Scanner on target folder
       run: |
        ./iac-scanner ./test-files > scan-results.log

     - name:  Display scan results
       run: |
        cat scan-results.log


     - name: Upload scan results
       uses: actions/upload-artifact@v4
       with:
        name: iac-scan-report
        path: scan-results.log

     - name:  Fail if HIGH severity issues are found
       run: |
        if grep -q "\[HIGH\]" scan-results.log; then
          echo "High severity issues detected. Failing the job!"
          exit 1
        fi




























name: Infrastructure Security Scan

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  workflow_dispatch:
    runs-on: ubuntu-latest

    env:
      SLACK_BOT_TOKEN: ${{ secrets.SLACK_BOT_TOKEN }}
      SLACK_CHANNEL_ID: ${{ secrets.SLACK_CHANNEL_ID }}

    steps:
      - name: Checkout code
        uses: actions/checkout@v3
name: CI Build with Hybrid Logging

on: [push]

jobs:
  build-and-log:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Run script and capture logs
        run: |
          echo "🚧 Starting build..." > log.txt
          echo "This is a test log line." >> log.txt
          echo "❌ ERROR: Something failed." >> log.txt

      # --- Upload to GitHub Artifacts ---
      - name: Upload log to GitHub Artifacts
        uses: actions/upload-artifact@v4
        with:
          name: build-log
          path: log.txt

      # --- Configure AWS for S3 upload ---
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v2
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ secrets.AWS_REGION }}

      # --- Upload to S3 ---
      - name: Set log filename
        id: logfile
        run: echo "filename=logs/log-${{ github.run_id }}.txt" >> $GITHUB_OUTPUT

      - name: Upload log to S3
        run: |
          aws s3 cp log.txt s3://my-ci-logs-bucket/${{ steps.logfile.outputs.filename }}

      # --- Generate Pre-signed URL ---
      - name: Get S3 Pre-signed URL
        id: signed_url
        run: |
          url=$(aws s3 presign s3://my-ci-logs-bucket/${{ steps.logfile.outputs.filename }} --expires-in 3600)
          echo "url=$url" >> $GITHUB_OUTPUT

      # - name: Debug Signed URL
      #   run: echo "Signed URL: ${{ steps.signed_url.outputs.url }}"

      # --- Send Slack Notification ---

      
      - name: Notify Slack
        run: |
          SHORT_LOG=$(tail -n 10 log.txt | sed 's/"/\\"/g')
          echo "S3 Link: ${{ steps.signed_url.outputs.url }}"
          curl -X POST -H 'Content-type: application/json' \
          --data "{
            \"text\": \"✅ CI Build Complete* for \`main\`\n
          \`\`\`${SHORT_LOG}\`\`\`\n
          🔗 *S3 Logs:* <${{ steps.signed_url.outputs.url }}>
          📦 *GitHub Artifact:* <https://github.com/your-org/your-repo/actions/runs/${{ github.run_id }}>\"
              }" \
              ${{ secrets.SLACK_WEBHOOK_URL2}}

      - name: Run IAC scan and save logs
        run: |
          echo "Running scan..."
          ./scanner/scan.sh > scan_output.log 2>&1 || true
        continue-on-error: true

      - name:  Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.22'

      - name:  Build IAC-Scanner
        run: |
          go build -o iac-scanner

      - name:  Run IAC-Scanner on target folder
        run: |
          ./iac-scanner ./test-files > scan-results.log

      - name:  Display scan results
        run: |
          cat scan-results.log


      - name: Upload scan results
        uses: actions/upload-artifact@v4
        with:
          name: iac-scan-report
          path: scan-results.log

      - name:  Fail if HIGH severity issues are found
        run: |
          if grep -q "\[HIGH\]" scan-results.log; then
          echo "High severity issues detected. Failing the job!"
          exit 1
          fi


      - name: Upload full logs to Slack (if scan fails)
        if: always()
        run: |
          curl -F file=@scan-results.log \
               -F "initial_comment= *IAC Scan Failed!*\nRepo: ${{ github.repository }}\n<${{ github.server_url }}/${{ github.repository }}/actions/runs/${{ github.run_id }}|View Full Workflow>" \
               -F channels=${{ secrets.SLACK_CHANNEL_ID }} \
               -H "Authorization: Bearer ${{ secrets.SLACK_BOT_TOKEN }}" \
               https://slack.com/api/files.upload

























































 Current 

 name: Infrastructure Security Scan

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  workflow_dispatch:
    runs-on: ubuntu-latest

    env:
      SLACK_BOT_TOKEN: ${{ secrets.SLACK_BOT_TOKEN }}
      SLACK_CHANNEL_ID: ${{ secrets.SLACK_CHANNEL_ID }}

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Run IAC scan and save logs
        run: |
          echo "Running scan..."
          ./scanner/scan.sh > scan_output.log 2>&1 || true
        continue-on-error: true

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.22'

      - name: Build IAC-Scanner
        run: |
          go build -o iac-scanner

      - name: Run IAC-Scanner on target folder
        run: |
          ./iac-scanner ./test-files > scan-results.log

      - name: Display scan results
        run: |
          cat scan-results.log

      - name: Upload scan results
        uses: actions/upload-artifact@v4
        with:
          name: iac-scan-report
          path: scan-results.log

      - name: Fail if HIGH severity issues are found
        run: |
          if grep -q "\[HIGH\]" scan-results.log; then
            echo "High severity issues detected. Failing the job!"
            exit 1
          fi
        continue-on-error: true

      - name: Get Slack Upload URL
        id: get_upload_url
        run: |
          FILE_NAME="scan-results.log"
          FILE_SIZE=$(wc -c < "$FILE_NAME" | xargs)

          echo "Uploading $FILE_NAME ($FILE_SIZE bytes)"

          response=$(curl -s -X POST https://slack.com/api/files.getUploadURLExternal \
            -H "Authorization: Bearer ${{ secrets.SLACK_BOT_TOKEN }}" \
            -H "Content-Type: application/json; charset=utf-8" \
            --data "$(jq -n \
              --arg filename "$FILE_NAME" \
              --argjson length "$FILE_SIZE" \
              --arg alt_text "Scan Results Log" \
              '{filename: $filename, length: $length, alt_text: $alt_text}')")

          echo "Slack API response: $response"

          upload_url=$(echo "$response" | jq -r '.upload_url')
          file_id=$(echo "$response" | jq -r '.file_id')

          echo "upload_url=$upload_url" >> $GITHUB_OUTPUT
          echo "file_id=$file_id" >> $GITHUB_OUTPUT


      - name: Upload file to Slack
        run: |name: CI Build with Hybrid Logging

on: [push]

jobs:
  build-and-log:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Run script and capture logs
        run: |
          echo "🚧 Starting build..." > log.txt
          echo "This is a test log line." >> log.txt
          echo "❌ ERROR: Something failed." >> log.txt

      # --- Upload to GitHub Artifacts ---
      - name: Upload log to GitHub Artifacts
        uses: actions/upload-artifact@v4
        with:
          name: build-log
          path: log.txt

      # --- Configure AWS for S3 upload ---
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v2
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ secrets.AWS_REGION }}

      # --- Upload to S3 ---
      - name: Set log filename
        id: logfile
        run: echo "filename=logs/log-${{ github.run_id }}.txt" >> $GITHUB_OUTPUT

      - name: Upload log to S3
        run: |
          aws s3 cp log.txt s3://my-ci-logs-bucket/${{ steps.logfile.outputs.filename }}

      # --- Generate Pre-signed URL ---
      - name: Get S3 Pre-signed URL
        id: signed_url
        run: |
          url=$(aws s3 presign s3://my-ci-logs-bucket/${{ steps.logfile.outputs.filename }} --expires-in 3600)
          echo "url=$url" >> $GITHUB_OUTPUT

      # - name: Debug Signed URL
      #   run: echo "Signed URL: ${{ steps.signed_url.outputs.url }}"

      # --- Send Slack Notification ---

      
      - name: Notify Slack
        run: |
          SHORT_LOG=$(tail -n 10 log.txt | sed 's/"/\\"/g')
          echo "S3 Link: ${{ steps.signed_url.outputs.url }}"
          curl -X POST -H 'Content-type: application/json' \
          --data "{
            \"text\": \"✅ CI Build Complete* for \`main\`\n
          \`\`\`${SHORT_LOG}\`\`\`\n
          🔗 *S3 Logs:* <${{ steps.signed_url.outputs.url }}>
          📦 *GitHub Artifact:* <https://github.com/your-org/your-repo/actions/runs/${{ github.run_id }}>\"
              }" \
              ${{ secrets.SLACK_WEBHOOK_URL2}}

          curl -X POST "${{ steps.get_upload_url.outputs.upload_url }}" \
            -F filename=@scan-results.log

      - name: Finalize Upload with Slack
        run: |
          curl -X POST https://slack.com/api/files.completeUploadExternal \
            -H "Authorization: Bearer ${{ secrets.SLACK_BOT_TOKEN }}" \
            -H "Content-Type: application/json" \
            --data "$(jq -n --arg id "${{ steps.get_upload_url.outputs.file_id }}" --argjson channel '["${{ secrets.SLACK_CHANNEL_ID }}"]' '{files:[{id:$id}], channel_ids:$channel}')"



























































































































































name: CI Build with Hybrid Logging

on: [push]

jobs:
  build-and-log:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Run script and capture logs
        run: |
          echo "🚧 Starting build..." > log.txt
          echo "This is a test log line." >> log.txt
          echo "❌ ERROR: Something failed." >> log.txt

      # --- Upload to GitHub Artifacts ---
      - name: Upload log to GitHub Artifacts
        uses: actions/upload-artifact@v4
        with:
          name: build-log
          path: log.txt

      # --- Configure AWS for S3 upload ---
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v2
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ secrets.AWS_REGION }}

      # --- Upload to S3 ---
      - name: Set log filename
        id: logfile
        run: echo "filename=logs/log-${{ github.run_id }}.txt" >> $GITHUB_OUTPUT

      - name: Upload log to S3
        run: |
          aws s3 cp log.txt s3://my-ci-logs-bucket/${{ steps.logfile.outputs.filename }}

      # --- Generate Pre-signed URL ---
      - name: Get S3 Pre-signed URL
        id: signed_url
        run: |
          url=$(aws s3 presign s3://my-ci-logs-bucket/${{ steps.logfile.outputs.filename }} --expires-in 3600)
          echo "url=$url" >> $GITHUB_OUTPUT

      # - name: Debug Signed URL
      #   run: echo "Signed URL: ${{ steps.signed_url.outputs.url }}"

      # --- Send Slack Notification ---

      
      - name: Notify Slack
        run: |
          SHORT_LOG=$(tail -n 10 log.txt | sed 's/"/\\"/g')
          echo "S3 Link: ${{ steps.signed_url.outputs.url }}"
          curl -X POST -H 'Content-type: application/json' \
          --data "{
            \"text\": \"✅ CI Build Complete* for \`main\`\n
          \`\`\`${SHORT_LOG}\`\`\`\n
          🔗 *S3 Logs:* <${{ steps.signed_url.outputs.url }}>
          📦 *GitHub Artifact:* <https://github.com/your-org/your-repo/actions/runs/${{ github.run_id }}>\"
              }" \
              ${{ secrets.SLACK_WEBHOOK_URL2}}
