# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/compare ./
rm -f dist/compare.zip
zip -j dist/compare.zip dist/compare
rm -f dist/compare
