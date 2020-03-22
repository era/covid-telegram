package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAll() (*AllCasesResponse, error) {
	var all AllCasesResponse
	err := call("all", &all)

	if err != nil {
		return nil, err
    }

	return &all, nil
}

func ByCountry (country string) (*CountryResponse, error) {
	var response CountryResponse
	err := call("countries/" + country, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func call(path string, responseObject interface{}) error {
	log.Print("Requesting ", path)
	endpoint := "https://corona.lmao.ninja/"
	response, err := http.Get(endpoint + path)
	if err != nil {
		return err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	log.Print("Response: ", string(responseData))
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return err
	}

	return nil
}
