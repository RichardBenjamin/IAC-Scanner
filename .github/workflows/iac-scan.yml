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
