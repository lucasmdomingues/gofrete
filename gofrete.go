package gofrete

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (frete *Frete) CalcFrete() (*Resultado, error) {

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
		return nil, fmt.Errorf("Erro: %s", resp.Status)
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

	if resultado.Servicos.Servico[0].Erro != "0" && resultado.Servicos.Servico[0].Erro != "011" {
		return nil, fmt.Errorf("Erro: %s", resultado.Servicos.Servico[0].MsgErro)
	}

	return &resultado, nil
}

func (r *Resultado) JSON() (string, error) {

	bytes, err := json.Marshal(r.Servicos.Servico[0])
	if err != nil {
		return "Error on parse json", err
	}

	return string(bytes), nil
}
