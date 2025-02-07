package main

import (
	"testing"
)

const numberOfCommands = 18

func TestCommand(t *testing.T) {
	if len(CommandTypes) != numberOfCommands {
		t.Errorf("Expected %d commands, got %d", numberOfCommands, len(CommandTypes))
	}

	if len(CommandFuncs) != numberOfCommands {
		t.Errorf("Expected %d commands, got %d", numberOfCommands, len(CommandTypes))
	}
}
