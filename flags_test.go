package main

import (
	"math"
	"testing"
)

func TestNoLimitIsSet(t *testing.T) {
	t.Run("none is set", func(t *testing.T) {
		flagMax := math.MaxInt32
		flagPrivateMax := math.MaxInt32
		flagPublicMax := math.MaxInt32

		result := noLimitIsSet(flagMax, flagPrivateMax, flagPublicMax)
		if result == false {
			t.Error("Expected result to be true")
		}
	})
	t.Run("one is set", func(t *testing.T) {
		flagMax := 1
		flagPrivateMax := math.MaxInt32
		flagPublicMax := math.MaxInt32

		result := noLimitIsSet(flagMax, flagPrivateMax, flagPublicMax)
		if result == true {
			t.Error("Expected result to be false")
		}
	})
	t.Run("two are set", func(t *testing.T) {
		flagMax := 1
		flagPrivateMax := 2
		flagPublicMax := math.MaxInt32

		result := noLimitIsSet(flagMax, flagPrivateMax, flagPublicMax)
		if result == true {
			t.Error("Expected result to be false")
		}
	})
	t.Run("all are set", func(t *testing.T) {
		flagMax := 1
		flagPrivateMax := 2
		flagPublicMax := 3

		result := noLimitIsSet(flagMax, flagPrivateMax, flagPublicMax)
		if result == true {
			t.Error("Expected result to be false")
		}
	})
}
