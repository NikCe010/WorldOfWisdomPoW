package proof_of_work

import (
	"context"
	"crypto/rand"
)

// Generate - method used to generate new challenge with selected complexity
func (s *Service) Generate(ctx context.Context) ([]byte, byte, error) {
	const numOfDataBytes byte = 8

	data := make([]byte, numOfDataBytes)
	_, err := rand.Read(data)
	if err != nil {
		return nil, 0, err
	}

	return data, s.complexity, nil
}
