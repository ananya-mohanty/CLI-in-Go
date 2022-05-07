package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/akamensky/argparse"
)

func make_get_request(url string) *http.Request {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return req
}

func make_post_request(url string, json_file_path string) *http.Request {

	// Read and parse JSON File for POST Body
	content, err := ioutil.ReadFile(json_file_path)
	if err != nil {
		log.Fatal(err)
	}
	user := map[string]string{}
	json.Unmarshal(content, &user)

	// POST Request

	postBody, _ := json.Marshal(user)
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		log.Fatalln(err)
	}
	return req
}

func set_custom_headers(req *http.Request, custom_headers *[]string) {

	// Adding custom headers passed by user
	for i := 0; i < len(*custom_headers); i++ {
		curr_header := strings.Split((*custom_headers)[i], "=")
		req.Header.Set(curr_header[0], curr_header[1])
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	output_body := string(body)
	log.Printf(output_body)

}
func main() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")

	request_type := parser.String("r", "request_type", &argparse.Options{Required: true, Help: "Type of API Request"})
	url := parser.String("u", "url", &argparse.Options{Required: false, Default: "http://headers.jsontest.com/", Help: "URL of API Request"})
	json_file := parser.String("j", "json_file", &argparse.Options{Required: false, Default: "users.json", Help: "JSON File Path for POST request"})
	custom_headers := parser.StringList("c", "custom_header", &argparse.Options{Required: false, Help: "List of custom headers"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *request_type == "GET" {
		req := make_get_request(*url)
		set_custom_headers(req, custom_headers)
	} else if *request_type == "POST" {
		req := make_post_request(*url, *json_file)
		set_custom_headers(req, custom_headers)
	}
}
