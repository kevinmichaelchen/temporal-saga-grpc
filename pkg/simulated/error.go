// Package simulated provides capabilities for random, simulated errors.
package simulated

import (
	"errors"
	"math/rand"

	"github.com/sirupsen/logrus"
)

var errInternalFailure = errors.New("oh no internal failure")

// ErrorProbability - Indicates how probable a simulated error should be.
type ErrorProbability int

const (
	// Low - Low probability.
	Low ErrorProbability = iota
	// MediumLow - Medium/low probability.
	MediumLow
	// Medium - Medium probability.
	Medium
	// MediumHigh - Medium/high probability.
	MediumHigh
	// High - High probability.
	High
)

// Int - Returns the probability as a number.
func (p ErrorProbability) Int() int {
	switch p {
	case Low:
		return 25
	case MediumLow:
		return 35
	case Medium:
		return 45
	case MediumHigh:
		return 65
	case High:
		return 75
	}

	return 50
}

// PossibleError - Returns an error (or none) at random.
func PossibleError(ep ErrorProbability) error {
	p := ep.Int()
	logrus.WithField("error_probability", p).Info("Rolling dice...")
	if rand.Intn(100) < p {
		return errInternalFailure
	}

	return nil
}
