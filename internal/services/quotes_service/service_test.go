package quotes_service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	service := New()

	assert.NotNil(t, service)
	assert.NotEmpty(t, service)
}

func TestService_GetRandomQuote(t *testing.T) {
	service := New()
	for i := 0; i < 3; i++ {
		quote := service.GetRandomQuote()
		assert.NotNil(t, quote)
		assert.NotEmpty(t, quote)
	}
}
