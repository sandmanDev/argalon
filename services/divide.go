package services

import "errors"

func DivideTwoByEachOther(a, b float64) (float64, float64, error) {

	if a == 0 || b == 0 {
		return 0, 0, errors.New("cannot divide by 0")
	}
	var result1, result2 float64

	result1 = a / b
	result2 = b / a

	return result1, result2, nil
}
