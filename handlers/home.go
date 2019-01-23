package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) {

	api := make(map[string]string)
	api["Title"] = "GoFrete (Correios)"
	api["Version"] = "1.0"

	json, err := json.Marshal(api)
	if err != nil {
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, string(json))
}
