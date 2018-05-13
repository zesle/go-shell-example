package main

import (
	"testing"
)

func TestRunCMD_noArg(t *testing.T) {
	err := execLine("ls")
	if err != nil {
		t.Errorf("Failed to run 'ls' command.")
	}
}

func TestRunCMD_withArg(t *testing.T) {
	err := execLine("ls -l")
	if err != nil {
		t.Errorf("Failed to run 'ls -l' command.")
	}
}
