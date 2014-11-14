package cmdtree

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// CommandExecutor is a function which is run when a command is
// resolved.
type CommandExecutor func(args string) error

// CommandExtender represents something that can extend a Command.
type CommandExtender interface {
	// SubCmd adds a sub-command to this command.
	SubCmd(trigger string, executor CommandExecutor) Command
}

// Command represents a command that can be executed or extended with
// sub-commands.
type Command interface {

	// CommandExtender ensures a Command can extend itself.
	CommandExtender

	// Execute parses the input given with the delimiter defined when
	// constructing the command. It then executes the command or
	// sub-command found, and returns any errors.
	Execute(input string) error

	// String returns the command tree in a human readable format.
	String() string
}

// Root creates a new command with no top-level commands. This is
// useful when you have multiple top-level commands.
func Root(delimiter string) Command {
	return NewCmd(delimiter, "", nil)
}

// NewCmd creates a new command with a top-level command filled out.
// This is useful for when you are going to extend this command.
func NewCmd(subCmdDelim, trigger string, executor CommandExecutor) Command {
	return &command{
		delimiter: subCmdDelim,
		trigger:   trigger,
		executor:  executor,
	}
}

type command struct {
	delimiter, trigger string
	subCmds            []*command
	executor           CommandExecutor
}

// SubCmd implements SubCmd on the CommandExtender interface.
func (tree *command) SubCmd(trigger string, executor CommandExecutor) Command {
	subCmd := &command{delimiter: tree.delimiter, trigger: trigger, executor: executor}
	tree.subCmds = append(tree.subCmds, subCmd)
	return subCmd
}

// Execute implements Execute on the Command interface.
func (tree *command) Execute(input string) error {

	cmdStrs := strings.Split(input, tree.delimiter)

	if tree.trigger != "" {

		if cmdStrs[0] != tree.trigger {
			// It all starts here. Need to at least match self.
			return fmt.Errorf(`could not match command "%s"`, input)
		} else if len(tree.subCmds) == 0 || len(cmdStrs) == 1 {
			// No sub-commands, or no requested sub-commands; just
			// execute self.

			if tree.executor == nil {
				return fmt.Errorf("command usage:\n%s", tree)
			}

			return tree.executor(strings.Join(cmdStrs[1:], tree.delimiter))
		}

		// Pop ourselves off so we can match the sub-commands.
		cmdStrs = cmdStrs[1:]
	}

	for _, subCmd := range tree.subCmds {
		if subCmd.trigger == cmdStrs[0] {
			// TODO(kate): PoC; splitting/concat on every subcmd is awful.
			return subCmd.Execute(strings.Join(cmdStrs, tree.delimiter))
		}
	}

	if tree.executor != nil {
		return tree.executor(strings.Join(cmdStrs, tree.delimiter))
	}

	return fmt.Errorf(`could not match sub-commands "%s"`, cmdStrs[0])
}

// String implements String on the Command interface.
func (tree *command) String() string {
	outputBuff := new(bytes.Buffer)
	tree.recurseString(outputBuff, 0)
	return outputBuff.String()
}

func (tree *command) recurseString(buff io.Writer, indentLevel int) {
	if indentLevel > 0 {
		fmt.Fprintf(buff, "\n")
		for i := 0; i < indentLevel; i++ {
			fmt.Fprintf(buff, "\t")
		}
	}

	fmt.Fprintf(buff, tree.trigger)

	for _, subCmd := range tree.subCmds {
		subCmd.recurseString(buff, indentLevel+1)
	}
}
