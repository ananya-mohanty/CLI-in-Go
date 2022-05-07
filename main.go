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

func make_get_request(url string, custom_headers *[]string) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < len(*custom_headers); i++ {
		curr_header := strings.Split((*custom_headers)[i], "=")
		req.Header.Set(curr_header[0], curr_header[1])
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func make_post_request(url string, json_file_path string, custom_headers *[]string) {

	// Read and parse JSON File for POST Body
	content, err := ioutil.ReadFile(json_file_path)
	if err != nil {
		log.Fatal(err)
	}
	user := map[string]string{}
	json.Unmarshal(content, &user)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	postBody, _ := json.Marshal(user)
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < len(*custom_headers); i++ {
		curr_header := strings.Split((*custom_headers)[i], "=")
		req.Header.Set(curr_header[0], curr_header[1])
	}

	if err != nil {
		log.Fatal(err)
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
	sb := string(body)
	log.Printf(sb)

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
		make_get_request(*url, custom_headers)
	} else if *request_type == "POST" {
		make_post_request(*url, *json_file, custom_headers)
	}

}
