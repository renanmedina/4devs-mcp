package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCnhService_Generate(t *testing.T) {
	t.Parallel()

	t.Run("Generate a random CNH number", func(t *testing.T) {
		cnhService := NewCnhService(false)
		cnh, err := cnhService.Generate()
		assert.NoError(t, err, "Expected no error when generating CNH")
		assert.NotNil(t, cnh, "Expected a non-nil CNH")
		assert.Equal(t, 11, len(*cnh), "Expected CNH to be 11 characters long")
	})
}
