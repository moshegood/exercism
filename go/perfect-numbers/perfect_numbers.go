package perfect

import (
	"fmt"
)

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = fmt.Errorf(">=0 !!")

func Classify(input int64) (Classification, error) {
	if input <= 0 {
		return ClassificationPerfect, ErrOnlyPositive
	}

	factors := []int64{}
	var i int64
	for i = 1; i < input; i++ {
		if input%i == 0 {
			factors = append(factors, i)
		}
	}

	var sum int64 = 0
	for _, factor := range factors {
		sum += factor
	}

	if sum == input {
		return ClassificationPerfect, nil

	}
	if sum < input {
		return ClassificationDeficient, nil
	}
	return ClassificationAbundant, nil
}
