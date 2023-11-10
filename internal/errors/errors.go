package errors

import "fmt"

var (
	MaxIterationErr            = fmt.Errorf("max iterations reached, no solution found")
	ChallengeSolvingTimeoutErr = fmt.Errorf("challenge solving timeout exceeded")
)
