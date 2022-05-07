# CLI-in-Go
Simple CLI in Go that makes GET and POST requests
- URL, custom headers and json file (for POST requests) can be sent using argparse
- Library used for argparse: github.com/akamensky/argparse
- Default values have been provided for url, json file path
- Sample input:
  - ananya$ go run main.go -r GET -u http://headers.jsontest.com/ -c Age=15 Name=User
  - ananya$ go run main.go -r POST -u "https://postman-echo.com/post" -j users.json -c phone=9000

