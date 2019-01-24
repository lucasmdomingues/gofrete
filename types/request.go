package types

import (
	"bytes"
	"io"
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

	client := http.Client{}

	var values io.Reader
	if r.Values != nil {
		values = bytes.NewBuffer(r.Values)
	}

	request, err := http.NewRequest(r.Method, r.Route, values)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
