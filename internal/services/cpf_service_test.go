package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpfService_Generate(t *testing.T) {
	t.Parallel()
	service := NewCpfService(false)

	t.Run("Generate a unformatted CPF", func(t *testing.T) {
		cpf, _ := service.Generate(false, "")
		assert.Equal(t, 11, len(*cpf))
	})

	t.Run("Generate a formatted CPF", func(t *testing.T) {
		cpf, _ := service.Generate(true, "")
		assert.Equal(t, 14, len(*cpf))
	})

	// 9th digit represents the state
	/* 1: DF, GO, MS, MT e TO;
	2: AC, AM, AP, PA, RO e RR;
	3: CE, MA e PI;
	4: AL, PB, PE, RN;
	5: BA e SE;
	6: MG;
	7: ES e RJ;
	8: SP;
	*/
	t.Run("Generate a CPF for SP address state", func(t *testing.T) {
		cpf, _ := service.Generate(false, "SP")
		assert.Equal(t, 11, len(*cpf))
		assert.Equal(t, (*cpf)[8:9], "8")
	})

	t.Run("Generate a CPF for RJ address state", func(t *testing.T) {
		cpf, _ := service.Generate(false, "RJ")
		assert.Equal(t, 11, len(*cpf))
		assert.Equal(t, (*cpf)[8:9], "7")
	})
}
