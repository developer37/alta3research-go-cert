package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	keyvalue := "DEMO_KEY"

	params := "api_key=" + url.QueryEscape(keyvalue)

	path := fmt.Sprintf("https://api.nasa.gov/neo/rest/v1/feed?%s", params)

	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

