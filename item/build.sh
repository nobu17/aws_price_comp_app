# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/item ./
rm -f dist/item.zip
zip -j dist/item.zip dist/item
rm -f dist/item
