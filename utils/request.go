package utils

import (
	"io/ioutil"
	"net/http"
)

func MakeRequest(w http.ResponseWriter, route string) ([]byte, error) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

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
