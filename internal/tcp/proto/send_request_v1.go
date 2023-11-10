package proto

// SendRequestV1 - model represent various types of request
type SendRequestV1 struct {
	Operation  Operation
	Complexity byte
	Content    []byte
}

// NewInitiateRequest - new request with Initiate Operation and empty payload
func NewInitiateRequest() *SendRequestV1 {
	return &SendRequestV1{
		Operation:  Initiate,
		Complexity: 0,
		Content:    nil,
	}
}

// NewSendChallengeRequest - new request with SendChallenge Operation. Contains complexity and challenge
func NewSendChallengeRequest(complexity byte, content []byte) *SendRequestV1 {
	return &SendRequestV1{
		Operation:  SendChallenge,
		Complexity: complexity,
		Content:    content,
	}
}

// NewSendDataRequest - new request with SendData Operation. Contains data
func NewSendDataRequest(content []byte) *SendRequestV1 {
	return &SendRequestV1{
		Operation:  SendData,
		Complexity: 0,
		Content:    content,
	}
}

// NewSolvedChallengeRequest - new request with SendNonce Operation. Contains solved challenge
func NewSolvedChallengeRequest(nonce []byte) *SendRequestV1 {
	return &SendRequestV1{
		Operation:  SendNonce,
		Complexity: 0,
		Content:    nonce,
	}
}

// ToMessage - conversion method to Message type
func (r *SendRequestV1) ToMessage() *Message {
	return &Message{
		Operation:  r.Operation,
		Complexity: r.Complexity,
		Length:     byte(len(r.Content)),
		Content:    r.Content,
	}
}
