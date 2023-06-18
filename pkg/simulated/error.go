package simulated

import (
	"errors"
	"math/rand"

	"github.com/sirupsen/logrus"
)

var errInternalFailure = errors.New("oh no internal failure")

type ErrorProbability int

const (
	Low ErrorProbability = iota
	MediumLow
	Medium
	MediumHigh
	High
)

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

func PossibleError(ep ErrorProbability) error {
	p := ep.Int()
	logrus.WithField("error_probability", p).Info("Rolling dice...")
	if rand.Intn(100) < p {
		return errInternalFailure
	}

	return nil
}
