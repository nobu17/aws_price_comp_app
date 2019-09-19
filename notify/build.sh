# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/notify ./
rm -f dist/notify.zip
zip -j dist/notify.zip dist/notify
rm -f dist/notify
