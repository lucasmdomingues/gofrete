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
	r.ResponseWriter.Header().Add("Access-Control-Allow-Credentials", "true")

	req, err := http.NewRequest(r.Method, r.Route, nil)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}
