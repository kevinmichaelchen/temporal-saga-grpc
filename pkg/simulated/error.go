package simulated

import (
	"errors"
	"math/rand"
)

func PossibleError() error {
	p := 34
	if rand.Intn(100) < p {
		return errors.New("oh no internal failure")
	}
	return nil
}
