package gofrete

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CalcFrete(frete *Frete) (*Resultado, error) {

	url := frete.NewURL()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/xml")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: %s, Try again.", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resultado := Resultado{}

	err = xml.Unmarshal(body, &resultado)
	if err != nil {
		return nil, err
	}

	if resultado.Servicos.Servico[0].MsgErro != "" {
		return nil, fmt.Errorf("Error: %s", resultado.Servicos.Servico[0].MsgErro)
	}

	return &resultado, nil
}
