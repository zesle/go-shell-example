package main

import (
	"testing"
)

func TestRunCMD(t *testing.T) {
	err := execLine("echo Hello World")
	if err != nil {
		t.Errorf("Failed to run 'echo' command.")
	}
}
