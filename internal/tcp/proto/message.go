package proto

// Message - represent data transferred throw protocol
type Message struct {
	Operation  Operation
	Complexity byte
	Length     byte
	Content    []byte
}

// ToBytes - method to pack message to slice of bytes
func (p *Message) ToBytes() []byte {
	res := make([]byte, 0, 3+p.Length)
	res = append(res, byte(p.Operation), p.Complexity, p.Length)
	res = append(res, p.Content...)
	return res
}
