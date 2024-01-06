// Package simulated provides capabilities for random, simulated errors.
package simulated

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

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

// SleepAndPossibleError - Sleeps and returns an error at random.
func SleepAndPossibleError(ep ErrorProbability) error {
	Sleep()

	return PossibleError(ep)
}

// PossibleError - Returns an error (or none) at random.
func PossibleError(ep ErrorProbability) error {
	probability := ep.Int()

	roll, err := generateRandomNumber(100)
	if err != nil {
		return err
	}

	logrus.WithField("error_probability", probability).Info("Rolling dice...")

	if roll < probability {
		return errInternalFailure
	}

	return nil
}

func generateRandomNumber(ceiling int) (int, error) {
	// Generate a random number between "[0, ceiling)".
	// That is, 0 (inclusive) and ceiling (exclusive).
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(ceiling)))
	if err != nil {
		return 0, fmt.Errorf("unable to generate random number: %w", err)
	}

	return int(randomNum.Int64()), nil
}
