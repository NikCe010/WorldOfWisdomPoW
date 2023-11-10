package proof_of_work

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_CheckNonce(t *testing.T) {
	service := NewGenerator(2)

	testCases := []struct {
		name     string
		nonce    []byte
		data     []byte
		expected bool
	}{
		{
			name:     "Success. Nonce solving the challenge. Should return true",
			nonce:    []byte{118, 53, 0, 0, 0, 0, 0, 0},
			data:     []byte{115, 212, 200, 67, 99, 124, 31, 222},
			expected: true,
		},
		{
			name:     "Error. Incorrect answer. Should return false",
			nonce:    []byte{118, 54, 0, 0, 0, 0, 0, 0},
			data:     []byte{115, 212, 200, 67, 99, 124, 31, 222},
			expected: false,
		},
	}

	for _, tc := range testCases {
		res := service.CheckNonce(tc.nonce, tc.data)
		assert.Equal(t, tc.expected, res)
	}
}
