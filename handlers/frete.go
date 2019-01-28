package handlers

import (
	"GoFrete/types"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func FreteHandler(w http.ResponseWriter, r *http.Request) {

	log := types.Log{}

	if err := r.ParseForm(); err != nil {
		log.Error(w, err)
		return
	}

	cdEmpresa := r.FormValue("CdEmpresa")

	dsSenha := r.FormValue("DsSenha")

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

	cepDestino := r.FormValue("CepDestino")
	if cepDestino == "" {
		log.Info(w, nil, "CEP de destino inválido.")
		return
	}

	vlPeso := r.FormValue("VlPeso")
	if vlPeso == "" {
		log.Info(w, nil, "Valor do peso inválido.")
		return
	}

	cdFormato := r.FormValue("CdFormato")
	if cdFormato == "" {
		log.Info(w, nil, "Código de formato inválido.")
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
		log.Info(w, nil, "Valor declarado inválido.")
		return
	}

	cdAvisoRecebimento := r.FormValue("CdAvisoRecebimento")
	if cdAvisoRecebimento == "" {
		log.Info(w, nil, "Aviso de recebimento inválido.")
		return
	}

	correios := types.Correios{}

	frete := &types.Frete{
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

	correios.Frete = frete

	route := correios.Frete.MakeURL()

	request := &types.Request{
		ResponseWriter: w,
		Route:          route,
		Method:         "GET",
		Values:         nil,
	}

	response, err := request.SendRequest()
	if err != nil {
		log.Error(w, err)
		return
	}

	resultado := types.Resultado{}

	err = xml.Unmarshal(response, &resultado)
	if err != nil {
		log.Error(w, err)
		return
	}

	json, err := json.Marshal(resultado.Servicos.Servico[0])
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
