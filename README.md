### Instalação
```go
go get github.com/lucasmdomingues/gofrete
```

### Códigos de servico dos correios populares
* 41106 PAC
* 40010 SEDEX
* 40045 SEDEX a Cobrar
* 40215 SEDEX 10

### Exemplo
```go
package main

import (
	"fmt"

	"github.com/lucasmdomingues/gofrete"
)

func main() {

	frete := gofrete.NewFrete("09146920", "123456", "40010", "07748415", "02019010",
		"0.100", "1", "16", "11", "11", "11", "S", "100", "S")

	resultado, err := frete.CalcFrete()
	if err != nil {
		fmt.Println(err)
		return
	}

	json, err := resultado.JSON()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(json)
}
```
```json
{"Codigo":"40010","Valor":"33,96","PrazoEntrega":"1","ValorMaoPropria":"6,80","ValorAvisoRecebimento":"5,75","ValorValorDeclarado":"1,61","EntregaDomiciliar":"S","EntregaSabado":"S","Erro":"0","MsgErro":"","ValorSemAdicionais":"19,80","ObsFim":""}
```
