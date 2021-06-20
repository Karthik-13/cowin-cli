package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetApiData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
		log.Printf("Could not request a CoWin API. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "CoWin CLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}
