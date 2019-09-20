# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/startup ./
rm -f dist/startup.zip
zip -j dist/startup.zip dist/startup
rm -f dist/startup
