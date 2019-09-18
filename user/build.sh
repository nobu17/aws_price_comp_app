# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/user ./
rm -f dist/user.zip
zip -j dist/user.zip dist/user
rm -f dist/user
