package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/akamensky/argparse"
)

func make_get_request(url string) {
	// GET Request
	if url == "" {
		url = "https://jsonplaceholder.typicode.com/posts"
	}
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func make_post_request(url string, json_file_path string) {

	if url == "" {
		url = "https://postman-echo.com/post"
	}

	if json_file_path == "" {
		json_file_path = "users.json"
	}
	content, err := ioutil.ReadFile(json_file_path)
	if err != nil {
		log.Fatal(err)
	}
	user := map[string]string{}
	json.Unmarshal(content, &user)
	if err != nil {
		log.Fatal(err)
	}

	postBody, _ := json.Marshal(user)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
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
	// Create new parser object
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	request_type := parser.String("r", "request_type", &argparse.Options{Required: true, Help: "Type of API Request"})
	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL of API Request"})
	json_file := parser.String("j", "json_file", &argparse.Options{Required: false, Help: "JSON File Path for POST request"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *request_type == "GET" {
		make_get_request(*url)
	} else if *request_type == "POST" {
		make_post_request(*url, *json_file)
	}

}
