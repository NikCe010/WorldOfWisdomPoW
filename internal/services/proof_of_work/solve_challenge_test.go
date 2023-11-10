package proof_of_work

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"worldofwisdom.com/m/internal/errors"
)

func TestService_SolveChallenge(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name          string
		complexity    byte
		data          []byte
		expectedNonce Nonce
		err           error
		timeout       int
	}{
		{
			name:          "Success. Solve challenge with complexity 2",
			complexity:    2,
			data:          []byte{115, 212, 200, 67, 99, 124, 31, 222},
			expectedNonce: 13686,
			err:           nil,
			timeout:       150,
		},
		{
			name:          "Success. Solve challenge with complexity 1",
			complexity:    1,
			data:          []byte{115, 212, 200, 67, 99, 124, 31, 222},
			expectedNonce: 339,
			err:           nil,
			timeout:       150,
		},
		{
			name:          "Error. Solve challenge with timeout = 0",
			complexity:    2,
			data:          []byte{115, 212, 200, 67, 99, 124, 31, 222},
			expectedNonce: 339,
			err:           errors.ChallengeSolvingTimeoutErr,
			timeout:       0,
		},
	}

	for _, tc := range testCases {
		service := NewSolver()
		ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(tc.timeout))
		defer cancel()
		nonce, err := service.SolveChallenge(ctxWithTimeout, tc.complexity, tc.data)
		if tc.err == nil {
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedNonce, nonce)
		} else {
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.err)
		}
	}
}
