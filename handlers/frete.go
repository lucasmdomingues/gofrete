package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gofrete/types"
	"gofrete/utils"
	"net/http"
)

func FreteHandler(w http.ResponseWriter, r *http.Request) {

	log := utils.Log{}

	if err := r.ParseForm(); err != nil {
		log.Error(w, err)
		return
	}

	cdServico := r.FormValue("CdServico")
	if cdServico == "" {
		log.Info(w, nil, "Código de serviço inválido.")
		return
	}

	cepOrigem := r.FormValue("CepOrigem")
	if cepOrigem == "" {
		log.Info(w, nil, "CEP de origem inválido.")
		return
	}

	CepDestino := r.FormValue("CepDestino")
	if CepDestino == "" {
		log.Info(w, nil, "CEP de destino inválido.")
		return
	}

	vlPeso := r.FormValue("VlPeso")
	if vlPeso == "" {
		log.Info(w, nil, "Valor do peso inválido.")
		return
	}

	vlComprimento := r.FormValue("VlComprimento")
	if vlComprimento == "" {
		log.Info(w, nil, "Valor do comprimento inválido.")
		return
	}

	vlAltura := r.FormValue("VlAltura")
	if vlAltura == "" {
		log.Info(w, nil, "Valor da altura inválido.")
		return
	}

	vlLargura := r.FormValue("VlLargura")
	if vlLargura == "" {
		log.Info(w, nil, "Valor da largura inválido.")
		return
	}

	vlDiametro := r.FormValue("VlDiametro")
	if vlDiametro == "" {
		log.Info(w, nil, "Valor do diametro inválido.")
		return
	}

	cdMaoPropria := r.FormValue("CdMaoPropria")
	if cdMaoPropria == "" {
		log.Info(w, nil, "Código de mão própria inválido")
		return
	}

	vlValorDeclarado := r.FormValue("VlValorDeclarado")
	if vlValorDeclarado == "" {
		log.Info(w, nil, "Valor do declarado inválido.")
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
		log.Error(w, err)
		return
	}

	resultado := types.Resultado{}

	err = xml.Unmarshal(data, &resultado)
	if err != nil {
		log.Error(w, err)
		return
	}

	json, err := json.Marshal(resultado)
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
