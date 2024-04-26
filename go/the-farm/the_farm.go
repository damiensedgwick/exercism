package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	deadCows int
	message  string
}

func (haveACow InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", haveACow.deadCows, haveACow.message)
}

func DivideFood(weightFodder FodderCalculator, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatFactor, err := weightFodder.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodder * fatFactor / float64(cows), nil
}

func ValidateInputAndDivideFood(weightFodder FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(weightFodder, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}

	if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	}

	return nil
}
