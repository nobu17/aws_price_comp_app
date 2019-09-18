# !/bin/bash
GOOS=linux GOARCH=amd64 go build -o dist/priceget ./
rm -f dist/priceget.zip
zip -j dist/priceget.zip dist/priceget
rm -f dist/priceget
