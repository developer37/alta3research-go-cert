package main

import (
        "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
        "os"
        "strconv"
)

type Neos struct {
     Neos [] Neo `json:"neo_reference_id"`
 }

type Neo struct {
     Magnitude float32  `json:"absolute_magnitude_h"`
     Hazardous bool     `json:"is_potentially_hazardous_asteroid"`
}	 
func main() {

	//keyvalue := "LpSn62kOWJcOsS0lulxJ831BffGpQSrkzCZhP22f"
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
	var neos  Neos
	json.Unmarshal(body, &neos)
	// print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(neos.Neos); i++ {
        fmt.Println("Neo Type: " + neos.Neos[i].Type)
        fmt.Println("Neo Magnitude: " + strconv.Itoa(neos.Neos[i].Magnitude))
        fmt.Println("Neo Hazardous: " + strconv.Itoa(neos.Neos[i].Hazardous))
         
    }
	
}

