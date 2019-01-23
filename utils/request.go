package utils

import (
	"io/ioutil"
	"net/http"
)

func MakeRequest(w http.ResponseWriter, route string) ([]byte, error) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	resp, err := http.Get(route)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
