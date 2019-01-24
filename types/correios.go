package types

import (
	"encoding/xml"
	"fmt"
)

const CORREIOS_URL = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx/CalcPrecoPrazo?"

type Correios struct {
	Frete *Frete
}

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

func (f *Frete) PopulateURL() string {

	url := fmt.Sprintf(CORREIOS_URL+"nCdEmpresa=%s&sDsSenha=%s&nCdServico=%s&sCepOrigem=%s&sCepDestino=%s&nVlPeso=%s&nCdFormato=%s&"+
		"nVlComprimento=%s&nVlAltura=%s&nVlLargura=%s&nVlDiametro=%s&sCdMaoPropria=%s&nVlValorDeclarado=%s&sCdAvisoRecebimento=%s&StrRetorno=%s",
		f.CdEmpresa, f.DsSenha, f.CdServico, f.CepOrigem, f.CepDestino, f.VlPeso, f.CdFormato, f.VlComprimento,
		f.VlAltura, f.VlLargura, f.VlDiametro, f.CdMaoPropria, f.VlValorDeclarado, f.CdAvisoRecebimento, f.StrRetorno,
	)

	return url
}
