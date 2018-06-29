package main

import (
	"os"
	"testing"
)

func TestExecInput_noArg(t *testing.T) {
	err := execInput("ls")
	if err != nil {
		t.Errorf("Failed to run 'ls' command: %v", err)
	}
}

func TestExecLine_withArg(t *testing.T) {
	err := execInput("ls -l")
	if err != nil {
		t.Errorf("Failed to run 'ls -l' command: %v", err)
	}
}

func TestCD_noArg(t *testing.T) {
	err := execInput("cd")
	if err == ErrNoPath {
		return
	}
	if err != nil {
		t.Errorf("Failed to run 'cd' command: %v", err)
	}
}
func TestCD_withArg(t *testing.T) {
	oldDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed to get curent directory: %v", err)
	}

	err = execInput("cd ..")
	if err != nil {
		t.Errorf("Failed to run 'cd ..' command: %v", err)
	}

	newDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed to get new directory: %v", err)
	}
	if oldDir == newDir {
		t.Errorf("Failed to change directory.")
	}
}
