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
        run: |
          curl -X POST "${{ steps.get_upload_url.outputs.upload_url }}" \
            -F filename=@scan-results.log

      - name: Finalize Upload with Slack
        run: |
          curl -X POST https://slack.com/api/files.completeUploadExternal \
            -H "Authorization: Bearer ${{ secrets.SLACK_BOT_TOKEN }}" \
            -H "Content-Type: application/json" \
            --data "$(jq -n --arg id "${{ steps.get_upload_url.outputs.file_id }}" --argjson channel '["${{ secrets.SLACK_CHANNEL_ID }}"]' '{files:[{id:$id}], channel_ids:$channel}')"
