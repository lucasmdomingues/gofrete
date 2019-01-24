package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gofrete/types"
	"gofrete/utils"
	"net/http"
)

var l *utils.Log

func FreteHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		l.Error(w, err)
		return
	}

	cdServico := r.FormValue("CdServico")
	if len(cdServico) == 0 {
		l.Info(w, nil, "Código de serviço inválido.")
		return
	}

	cepOrigem := r.FormValue("CepOrigem")
	if len(cepOrigem) == 0 {
		l.Info(w, nil, "CEP de origem inválido.")
		return
	}

	CepDestino := r.FormValue("CepDestino")
	if len(CepDestino) == 0 {
		l.Info(w, nil, "CEP de destino inválido.")
		return
	}

	vlPeso := r.FormValue("VlPeso")
	if len(vlPeso) == 0 {
		l.Info(w, nil, "Valor do peso inválido.")
		return
	}

	vlComprimento := r.FormValue("VlComprimento")
	if len(vlComprimento) == 0 {
		l.Info(w, nil, "Valor do comprimento inválido.")
		return
	}

	vlAltura := r.FormValue("VlAltura")
	if len(vlAltura) == 0 {
		l.Info(w, nil, "Valor da altura inválido.")
		return
	}

	vlLargura := r.FormValue("VlLargura")
	if len(vlLargura) == 0 {
		l.Info(w, nil, "Valor da largura inválido.")
		return
	}

	vlDiametro := r.FormValue("VlDiametro")
	if len(vlDiametro) == 0 {
		l.Info(w, nil, "Valor do diametro inválido.")
		return
	}

	cdMaoPropria := r.FormValue("CdMaoPropria")
	if len(cdMaoPropria) == 0 {
		l.Info(w, nil, "Código de mão própria inválido")
		return
	}

	vlValorDeclarado := r.FormValue("VlValorDeclarado")
	if len(cdServico) == 0 {
		l.Info(w, nil, "Valor do declarado inválido.")
		return
	}

	correios := &types.Correios{
		Frete: &types.Frete{
			CdEmpresa:          "",
			DsSenha:            "",
			CdServico:          cdServico,
			CepOrigem:          cepOrigem,
			CepDestino:         CepDestino,
			VlPeso:             vlPeso,
			CdFormato:          "1",
			VlComprimento:      vlComprimento,
			VlAltura:           vlAltura,
			VlLargura:          vlLargura,
			VlDiametro:         vlDiametro,
			CdMaoPropria:       cdMaoPropria,
			VlValorDeclarado:   vlValorDeclarado,
			CdAvisoRecebimento: "N",
			StrRetorno:         "xml",
		},
	}

	route := correios.Frete.PopulateURL()

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
