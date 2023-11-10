package proof_of_work

// Service ...
type Service struct {
	complexity byte
}

// NewGenerator ...
func NewGenerator(complexity byte) *Service {
	return &Service{complexity: complexity}
}

// NewSolver ...
func NewSolver() *Service {
	return &Service{}
}
