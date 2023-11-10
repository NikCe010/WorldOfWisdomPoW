package proof_of_work

import (
	"context"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Generate(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name               string
		complexity         byte
		expectedComplexity byte
	}{
		{
			name:               "Success. Generate challenge with complexity = 2",
			complexity:         2,
			expectedComplexity: 2,
		},
		{
			name:               "Success. Generate challenge with complexity = 5",
			complexity:         5,
			expectedComplexity: 5,
		},
		{
			name:               "Success. Generate challenge with complexity = 0",
			complexity:         0,
			expectedComplexity: 0,
		},
		{
			name:               "Success. Generate challenge with complexity = math.MaxInt8",
			complexity:         math.MaxInt8,
			expectedComplexity: math.MaxInt8,
		},
	}

	for _, tc := range testCases {
		service := NewGenerator(tc.complexity)
		challenge, complexity, err := service.Generate(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, challenge)
		assert.NotEmpty(t, challenge)
		assert.Equal(t, tc.expectedComplexity, complexity)
	}
}
