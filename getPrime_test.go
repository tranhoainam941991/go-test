package main

import "testing"

func TestGetPrime(t *testing.T) {
	lessThenZero, lessThenZeroError := getPrime(0)

	if lessThenZeroError == nil && lessThenZero != 0 {
		t.Errorf("failed 1")
	}

	oneRes, oneErr := getPrime(1)

	if oneErr != nil || oneRes != 1 {
		t.Errorf("fail 2")
	}

	fivefiveRes, fivefiveErr := getPrime(55)

	if fivefiveErr != nil || fivefiveRes != 53 {
		t.Errorf("fail 3")
	}
}
