package gofrete

import (
	"encoding/xml"
)

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
	XMLName  xml.Name `xml:"Servicos"`
	Servicos Servico  `xml:"cServico"`
}

type Servico struct {
	Codigo                string `xml:"Codigo"`
	Valor                 string `xml:"Valor"`
	PrazoEntrega          string `xml:"PrazoEntrega"`
	ValorSemAdicionais    string `xml:"ValorSemAdicionais"`
	ValorMaoPropria       string `xml:"ValorMaoPropria"`
	ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento"`
	ValorValorDeclarado   string `xml:"ValorValorDeclarado"`
	EntregaDomiciliar     string `xml:"EntregaDomiciliar"`
	EntregaSabado         string `xml:"EntregaSabado"`
	Erro                  string `xml:"Erro"`
	MsgErro               string `xml:"MsgErro"`
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
