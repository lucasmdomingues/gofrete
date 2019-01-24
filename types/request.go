package types

import (
	"io/ioutil"
	"net/http"
)

type Request struct {
	ResponseWriter http.ResponseWriter
	Route          string
	Method         string
	Values         []byte
}

func (r *Request) SendRequest() ([]byte, error) {

	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	r.ResponseWriter.Header().Add("Access-Control-Allow-Credentials", "true")

	client := http.Client{}

	resp, err := client.Get(r.Route)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
