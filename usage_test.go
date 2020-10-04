package main

import "testing"

func TestPrintUsage(t *testing.T) {
	var callCount uint
	p := mockPrinter{callCount: &callCount}

	printUsage(p)
	if callCount != 1 {
		t.Errorf("Expected printer to be called once (called %d times)", callCount)
	}
}
