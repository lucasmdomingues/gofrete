package types

import "encoding/xml"

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
