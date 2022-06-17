# alta3research-go-cert

This project is to demonstrate proficiency in the GO language. A REST client written in Go will make a HTTP GET call to a NASA REST API to retrieve the current NEO's (Near Earth Objects) Ids and urls which be used to retrieve more details from the NASA site. The NEO ids and urls will be collected in a map and printed. A follow up HTTP GET call will be made for an NEO and its details will be printed. The code can be easily modified to print more pertinent data as well as re-packaged so that func GetNeoDetails can be called by other Go programs.
The project demonstrates how GO can simplify the processing of complex JSON in the responses of REST HTTP calls.   

## Getting Started

These instructions will get you a copy of the project up and running on your local machine
for development and testing purposes. 

1. Make          directory: mdir simple
2. Change into   directory: cd simple
3. Git cmd to install code: git clone https://github.com/developer37/alta3research-go-cert.git
4. Get into project dir   : cd alta3research-go-cert
5. Run program            : go run alta3research-gocert01.go



### Prerequisites

1. Make sure Go is installed on your machine: go version
2. If GO is not installed,  go to https://go.dev/dl/ and follow instructions there to install on your specific operating system.

## Built With

* GO - The coding language used

## Authors

* **Developer37** - *Initial work* - (https://example.com/)

## References
1. https://alta3.com/overview-golang-training
2. https://api.nasa.gov
3. https://go.dev/doc/
4. https://transform.tools/json-to-go
5. https://blog.logrocket.com/using-json-go-guide/
6. https://www.sohamkamani.com/golang/json/

