package gofrete

import (
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/net/html/charset"
)

type Service interface {
	Calculate(frete *Frete) (*Resultado, error)
}

type service struct {
	BaseURL url.URL
}

func NewService() Service {
	return &service{
		BaseURL: url.URL{
			Scheme: "https",
			Host:   "ws.correios.com.br",
			Path:   "calculador/",
		},
	}
}

func (s *service) Calculate(frete *Frete) (*Resultado, error) {
	url, err := s.BaseURL.Parse(path.Join("CalcPrecoPrazo.aspx"))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("nCdEmpresa", frete.CdEmpresa)
	q.Add("sDsSenha", frete.DsSenha)
	q.Add("nCdServico", frete.CdServico)
	q.Add("sCepOrigem", frete.CepOrigem)
	q.Add("sCepDestino", frete.CepDestino)
	q.Add("nVlPeso", frete.VlPeso)
	q.Add("nCdFormato", frete.CdFormato)
	q.Add("nVlComprimento", frete.VlComprimento)
	q.Add("nVlAltura", frete.VlAltura)
	q.Add("nVlLargura", frete.VlLargura)
	q.Add("nVlDiametro", frete.VlDiametro)
	q.Add("sCdMaoPropria", frete.CdMaoPropria)
	q.Add("nVlValorDeclarado", frete.VlValorDeclarado)
	q.Add("sCdAvisoRecebimento", frete.CdAvisoRecebimento)
	q.Add("StrRetorno", "xml")

	url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/xml")

	var client http.Client

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var resultado Resultado

	err = decoder.Decode(&resultado)
	if err != nil {
		return nil, err
	}

	if resultado.Servicos.Erro != "0" && resultado.Servicos.Erro != "011" {
		return nil, errors.New(resultado.Servicos.MsgErro)
	}

	return &resultado, nil
}
