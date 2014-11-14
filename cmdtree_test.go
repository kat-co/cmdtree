package cmdtree

import (
	"fmt"
	"testing"
)

func TestRoot(t *testing.T) {

	helpRan := false
	root := Root(" ")
	root.SubCmd("help", func(args string) error {
		helpRan = true
		return nil
	})

	if err := root.Execute("help"); err != nil {
		t.Errorf("Received unexpected error: %v", err)
	} else if !helpRan {
		t.Error(`"help" command did not run.`)
	}
}

func TestExecuteCommand(t *testing.T) {

	rootRun := false
	subRun := false

	help := NewCmd(" ", "help", func(args string) error {
		rootRun = true
		return nil
	})
	help.SubCmd("cmdtree", func(args string) error {
		subRun = true
		return nil
	})

	// Test executing the root
	if err := help.Execute("help"); err != nil {
		t.Errorf("Unexpected error received: %v", err)
	} else if !rootRun {
		t.Error("Root was not run.")
	}

	// Test executing sub commands.
	if err := help.Execute("help cmdtree"); err != nil {
		t.Errorf("Unexpected error received: %v", err)
	} else if !subRun {
		t.Error("Command was not run.")
	}
}

func TestExecuteEmptySubCmd(t *testing.T) {

	help := NewCmd(" ", "help", nil)
	help.SubCmd("cmdtree", nil)

	if err := help.Execute("help"); err == nil {
		t.Error("Expected a usage error.")
	} else if expected := fmt.Sprintf("Command usage:\n%s", help); err.Error() != expected {
		t.Logf("Got: %v", err)
		t.Errorf("Expected: %s", expected)
	}
}

func TestExecuteCommandWhenArgumentsNotMatch(t *testing.T) {

	var helpArgsPassed string

	help := NewCmd(" ", "help", func(passed string) error {
		helpArgsPassed = passed
		return nil
	})
	help.SubCmd("cmdtree", nil)

	if err := help.Execute("help 1 2 3"); err != nil {
		t.Error("Unexpected error.")
	} else if helpArgsPassed == "" {
		t.Error("help was not run")
	} else if helpArgsPassed != "1 2 3" {
		t.Errorf("help not passed correct arguments: %s", helpArgsPassed)
	}
}

func TestCommandUsage(t *testing.T) {

	help := NewCmd(" ", "help", nil)
	help.SubCmd("cmdtree", nil).SubCmd("usage", nil)
	help.SubCmd("golang", nil)

	const expected = "help\n\tcmdtree\n\t\tusage\n\tgolang"

	if got := help.String(); got != expected {
		t.Logf("Got: %s", got)
		t.Errorf("Expected output: %s", expected)
	}
}
