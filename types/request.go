package types

import (
	"bytes"
	"io/ioutil"
	"log"
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

	var values *bytes.Buffer
	if r.Values != nil {
		values = bytes.NewBuffer(r.Values)
	}

	request, err := http.NewRequest(r.Method, r.Route, values)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body, nil
}
