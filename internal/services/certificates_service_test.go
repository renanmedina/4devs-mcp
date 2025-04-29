package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCertificatesService_Generate(t *testing.T) {
	t.Parallel()

	t.Run("Generate a birth certificate number", func(t *testing.T) {
		certificatesService := NewCertificatesService(false)
		certificateNumber, err := certificatesService.Generate(false, BIRTHDATE_CERTIFICATE)
		assert.NoError(t, err, "Expected no error when generating certificate")
		assert.Equal(t, len(*certificateNumber), 32, "Expected certificate to be 32 characters long")
	})

	t.Run("Generate a formatted birth certificate number", func(t *testing.T) {
		certificatesService := NewCertificatesService(false)
		certificateNumber, err := certificatesService.Generate(true, BIRTHDATE_CERTIFICATE)
		assert.NoError(t, err, "Expected no error when generating certificate")
		assert.Equal(t, len(*certificateNumber), 40, "Expected certificate to be 40 characters long")
	})

	t.Run("Generate a marriage certificate number", func(t *testing.T) {
		certificatesService := NewCertificatesService(false)
		certificateNumber, err := certificatesService.Generate(false, MARRIAGE_CERTIFICATE)
		assert.NoError(t, err, "Expected no error when generating certificate")
		assert.Equal(t, len(*certificateNumber), 32, "Expected certificate to be 32 characters long")
	})

	t.Run("Generate a formatted marriage certificate number", func(t *testing.T) {
		certificatesService := NewCertificatesService(false)
		certificateNumber, err := certificatesService.Generate(true, MARRIAGE_CERTIFICATE)
		assert.NoError(t, err, "Expected no error when generating certificate")
		assert.Equal(t, len(*certificateNumber), 40, "Expected certificate to be 40 characters long")
	})
}
