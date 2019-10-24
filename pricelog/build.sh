# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/pricelog ./
rm -f dist/pricelog.zip
zip -j dist/pricelog.zip dist/pricelog
rm -f dist/pricelog
