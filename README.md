# alta3research-go-cert

This project is to demonstrate proficiency in the GO language. A REST client written in Go will make a HTTP GET call to a NASA REST API to retrieve the current NEO's (Near Earth Objects) Ids and urls which be used to retrieve more details from the NASA site. The NEO ids and urls will be collected in a map and printed. A follow up HTTP GET call will be made for an NEO and its details will be printed. The code can be easily modified to print more pertinent data as well as re-packaged so that func GetNeoDetails can be called by other Go programs.
The project demonstrates how GO can simplify the processing of complex JSON in the responses of REST HTTP calls.   

## Getting Started

These instructions will get you a copy of the project up and running on your local machine
for development and testing purposes. 

-Make          directory: mdir simple
-Change into   directory: cd simple
-Git cmd to install code: git clone https://github.com/developer37/alta3research-go-cert.git
-Get into project dir   : cd alta3research-go-cert
-Run program            : go run alta3research-gocert01.go



### Prerequisites

Make sure Go is installed on your machine: go version
If it is not install go to https://go.dev/dl/ and follow instructions there to install on specific operating system

## Built With

* GO - The coding language used

## Authors

* **Developer37** - *Initial work* - (https://example.com/)

## References
https://alta3.com/overview-golang-training
https://api.nasa.gov
https://go.dev/doc/
https://transform.tools/json-to-go
https://blog.logrocket.com/using-json-go-guide/
https://www.sohamkamani.com/golang/json/

