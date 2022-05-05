package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	name  string
	email string
}

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
	user2 := User{}
	json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}

	//Encode the data
	postBody, _ := json.Marshal(user2)
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

}

func main() {
	fmt.Println("Type of request: ")

	// var then variable name then variable type
	var request_type string

	// Taking input from user
	fmt.Scanln(&request_type)
	if request_type == "GET" {
		var url string
		fmt.Scanln(&url)
		make_get_request(url)
	} else if request_type == "POST" {
		var url string
		var json_file_path string
		fmt.Scanln(&url)
		fmt.Scanln(&json_file_path)
		make_post_request(url, json_file_path)
	}

}
