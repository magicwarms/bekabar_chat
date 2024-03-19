package main

import (
	"testing"
)

func TestApplicationRunning(t *testing.T) {
	err := RunApplication("testing")
	if err != nil {
		t.Fatal(err)
	}
}
