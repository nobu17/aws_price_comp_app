# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/starup ./
rm -f dist/starup.zip
zip -j dist/starup.zip dist/starup
rm -f dist/starup
