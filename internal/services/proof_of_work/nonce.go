package proof_of_work

import "encoding/binary"

// Nonce ...
type Nonce uint64

// NonceFromBytes ...
func NonceFromBytes(nonceBytes []byte) Nonce {
	return Nonce(binary.LittleEndian.Uint64(nonceBytes))
}

// ToBytes ...
func (n Nonce) ToBytes() []byte {
	nonceBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceBytes, uint64(n))

	return nonceBytes
}
