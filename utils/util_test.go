package utils

import "testing"

func TestRandInt(t *testing.T) {
	for i := 0; i < 22; i++ {
		t.Log(RandInt(3))
	}
}
