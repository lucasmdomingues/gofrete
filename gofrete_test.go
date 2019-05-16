package gofrete

import (
	"fmt"
	"testing"
)

func TestCalcFrete(t *testing.T) {

	frete := NewFrete("09146920", "123456", "40010", "07748415", "02019010",
		"0.100", "1", "16", "11", "11", "11", "S", "100", "S")

	resultado, err := frete.CalcFrete()
	if err != nil {
		t.Error(err)
		return
	}

	json, err := resultado.JSON()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(json)
}
