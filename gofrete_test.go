package gofrete

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	service := NewService()

	frete := NewFrete("09146920", "123456", "40010", "07748415", "02019010",
		"0.100", "1", "16", "11", "11", "11", "S", "100", "S")

	t.Run("should be able calculate price", func(t *testing.T) {
		_, err := service.Calculate(frete)
		assert.Nil(t, err)
	})

	t.Run("should be able return error passing invalid service code", func(t *testing.T) {
		frete.CdServico = "123"

		_, err := service.Calculate(frete)
		assert.NotNil(t, err)
	})
}
