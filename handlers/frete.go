package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gofrete/types"
	"gofrete/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var l utils.Log

const CORREIOS = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx/CalcPrecoPrazo?"

func FreteHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		l.Error(w, err)
		return
	}

	vars := mux.Vars(r)

	cep := vars["cep"]
	if cep == "" || len(cep) < 8 {
		l.Info(w, nil, "CEP is required")
		return
	}

	frete := &types.Frete{
		CdEmpresa:          "",
		DsSenha:            "",
		CdServico:          "40010",
		CepOrigem:          "01107010",
		CepDestino:         cep,
		VlPeso:             "0.10",
		CdFormato:          "1",
		VlComprimento:      "16",
		VlAltura:           "2",
		VlLargura:          "10",
		VlDiametro:         "0",
		CdMaoPropria:       "N",
		VlValorDeclarado:   "100",
		CdAvisoRecebimento: "N",
		StrRetorno:         "xml",
	}

	route := fmt.Sprintf(CORREIOS+"nCdEmpresa=%s&sDsSenha=%s&nCdServico=%s&sCepOrigem=%s&sCepDestino=%s&nVlPeso=%s&nCdFormato=%s&"+
		"nVlComprimento=%s&nVlAltura=%s&nVlLargura=%s&nVlDiametro=%s&sCdMaoPropria=%s&nVlValorDeclarado=%s&sCdAvisoRecebimento=%s&StrRetorno=%s",
		frete.CdEmpresa, frete.DsSenha, frete.CdServico, frete.CepOrigem, frete.CepDestino, frete.VlPeso, frete.CdFormato, frete.VlComprimento,
		frete.VlAltura, frete.VlLargura, frete.VlDiametro, frete.CdMaoPropria, frete.VlValorDeclarado, frete.CdAvisoRecebimento, frete.StrRetorno,
	)

	data, err := utils.MakeRequest(w, route)
	if err != nil {
		l.Error(w, err)
		return
	}

	resultado := types.Resultado{}

	err = xml.Unmarshal(data, &resultado)
	if err != nil {
		l.Error(w, err)
		return
	}

	json, err := json.Marshal(resultado)
	if err != nil {
		l.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
