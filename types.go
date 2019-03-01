package gofrete

import (
	"encoding/xml"
	"fmt"
)

const WSCORREIOS_FRETE = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx/CalcPrecoPrazo?"

type Frete struct {
	CdEmpresa          string `xml:"nCdEmpresa"`
	DsSenha            string `xml:"sDsSenha"`
	CdServico          string `xml:"nCdServico"`
	CepOrigem          string `xml:"sCepOrigem"`
	CepDestino         string `xml:"sCepDestino"`
	VlPeso             string `xml:"nVlPeso"`
	CdFormato          string `xml:"nCdFormato"`
	VlComprimento      string `xml:"nVlComprimento"`
	VlAltura           string `xml:"nVlAltura"`
	VlLargura          string `xml:"nVlLargura"`
	VlDiametro         string `xml:"nVlDiametro"`
	CdMaoPropria       string `xml:"sCdMaoPropria"`
	VlValorDeclarado   string `xml:"nVlValorDeclarado"`
	CdAvisoRecebimento string `xml:"sCdAvisoRecebimento"`
	StrRetorno         string `xml:"StrRetorno"`
}

type Resultado struct {
	XMLName  xml.Name `xml:"cResultado"`
	Servicos Servicos `xml:"Servicos"`
}

type Servicos struct {
	XMLName xml.Name  `xml:"Servicos"`
	Servico []Servico `xml:"cServico"`
}

type Servico struct {
	Codigo                string `xml:"Codigo"`
	Valor                 string `xml:"Valor"`
	PrazoEntrega          string `xml:"PrazoEntrega"`
	ValorMaoPropria       string `xml:"ValorMaoPropria"`
	ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento"`
	ValorValorDeclarado   string `xml:"ValorValorDeclarado"`
	EntregaDomiciliar     string `xml:"EntregaDomiciliar"`
	EntregaSabado         string `xml:"EntregaSabado"`
	Erro                  string `xml:"Erro"`
	MsgErro               string `xml:"MsgErro"`
	ValorSemAdicionais    string `xml:"ValorSemAdicionais"`
	ObsFim                string `xml:"obsFim"`
}

func NewFrete(cdEmpresa, dsSenha, cdServico, cepOrigem, cepDestino, vlPeso, cdFormato, vlComprimento, vlAltura,
	vlLargura, vlDiametro, cdMaoPropria, vlValorDeclarado, cdAvisoRecebimento string) *Frete {

	return &Frete{
		CdEmpresa:          cdEmpresa,
		DsSenha:            dsSenha,
		CdServico:          cdServico,
		CepOrigem:          cepOrigem,
		CepDestino:         cepDestino,
		VlPeso:             vlPeso,
		CdFormato:          cdFormato,
		VlComprimento:      vlComprimento,
		VlAltura:           vlAltura,
		VlLargura:          vlLargura,
		VlDiametro:         vlDiametro,
		CdMaoPropria:       cdMaoPropria,
		VlValorDeclarado:   vlValorDeclarado,
		CdAvisoRecebimento: cdAvisoRecebimento,
	}
}

func (f *Frete) NewURL() string {

	values := make(map[string]string)

	values["nCdEmpresa"] = f.CdEmpresa
	values["sDsSenha"] = f.DsSenha
	values["nCdServico"] = f.CdServico
	values["sCepOrigem"] = f.CepOrigem
	values["sCepDestino"] = f.CepDestino
	values["nVlPeso"] = f.VlPeso
	values["nCdFormato"] = f.CdFormato
	values["nVlComprimento"] = f.VlComprimento
	values["nVlAltura"] = f.VlAltura
	values["nVlLargura"] = f.VlLargura
	values["nVlDiametro"] = f.VlDiametro
	values["sCdMaoPropria"] = f.CdMaoPropria
	values["nVlValorDeclarado"] = f.VlValorDeclarado
	values["sCdAvisoRecebimento"] = f.CdAvisoRecebimento
	values["StrRetorno"] = "xml"

	url := WSCORREIOS_FRETE
	for key, value := range values {
		url += fmt.Sprintf("%s=%s&", key, value)
	}

	return url
}
