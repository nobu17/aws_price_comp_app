# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/alert ./
rm -f dist/alert.zip
zip -j dist/alert.zip dist/alert
rm -f dist/alert
