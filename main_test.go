package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(2,2)
	if got !=4 {
		t.Errorf("Expected 4, got %d", got)
	}
}