package proof_of_work

import (
	"context"
	"math"

	"worldofwisdom.com/m/internal/errors"
)

// SolveChallenge - method used to solve challenge
func (s *Service) SolveChallenge(ctx context.Context, complexity byte, data []byte) (Nonce, error) {
	var nonce Nonce = 0
	for ; nonce < math.MaxUint64; nonce++ {
		select {
		case <-ctx.Done():
			return 0, errors.ChallengeSolvingTimeoutErr
		default:

		}
		if s.check(complexity, nonce, data) {
			return nonce, nil
		}
	}

	return 0, errors.MaxIterationErr // practically unreachable
}
