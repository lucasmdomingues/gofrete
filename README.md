### Instalação
```go
go get github.com/lucasmdomingues/gofrete
```

### Exemplo
```go
package main

import (
	"fmt"

	"github.com/lucasmdomingues/gofrete"
)

func main() {

	frete := gofrete.NewFrete("", "", "40010", "07748415", "02019010",
		"0.100", "1", "16", "11", "11", "11", "S", "100", "S")

	resultado, err := gofrete.CalcFrete(frete)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", resultado)
}
```
