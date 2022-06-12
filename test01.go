package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	keyvalue := "LpSn62kOWJcOsS0lulxJ831BffGpQSrkzCZhP22f"

	params := "api_key=" + url.QueryEscape(keyvalue)

	path := fmt.Sprintf("https://api.nasa.gov/planetary/apod?%s", params)

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

