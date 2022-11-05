package main

import (
	"testing"
)

// Testing testing in Go
func TestInputType(t *testing.T) {
	got := input()

	if got == -1 {
		t.Errorf("An invalid number was inputted!")
	}
}
