package proof_of_work

import (
	"bytes"
	"crypto/sha256"
)

// CheckNonce - method used to check correctness of the client answer
func (s *Service) CheckNonce(nonceBytes []byte, data []byte) bool {
	return s.check(s.complexity, NonceFromBytes(nonceBytes), data)
}

func (s *Service) check(complexity byte, nonce Nonce, data []byte) bool {
	prefix := make([]byte, complexity)

	sum := append(data, nonce.ToBytes()...)
	hash := sha256.Sum256(sum)
	res := bytes.HasPrefix(hash[:], prefix)
	if res {
		return true
	}
	return false
}
