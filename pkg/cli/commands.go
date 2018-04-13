package cli

import (
	"errors"
	"fmt"
	"strings"
)

type commandType uint

const (
	Help commandType = iota
	AddResource
	RemoveResource
	Info
)

/*
	cliCommand provides default structure of cli option
	Args:
	ct    - Shows the type of the command
	short - (optional) short name of an option
	long  - long name of an option
	expArgs - provide an expected input args format
		The syntax for expected args:
			%s - single string value
			%t - true or false
			%d - decimal int
	gotArgs - list of parsed args
*/
type cliCommand struct {
	ct      commandType
	short   uint8
	long    string
	expArgs string
	// forbidden to declare on creation
	gotArgs []interface{}
}

func newCliCommandRegistry() []cliCommand {
	return []cliCommand{
		{ct: Help, long: "help", short: 'h'},
		{ct: AddResource, long: "add", expArgs: "%s"},
		{ct: RemoveResource, long: "remove", expArgs: "%s"},
		{ct: Info, long: "info", short: 'i'}}
}

var cliCommands = newCliCommandRegistry()

func (c *cliCommand) ParseCommand(word string) error {
	switch len(word) {
	case 1:
		return errors.New("can't parse empty option")
	case 2:
		for _, cmd := range cliCommands {
			if cmd.short != 0 && word[1] == cmd.short {
				c.ct, c.expArgs = cmd.ct, cmd.expArgs
				return nil
			}
		}
		return errors.New("there is no such option")
	default:
		trmWord := strings.ToLower(strings.TrimSuffix(word, "--"))
		for _, cmd := range cliCommands {
			if cmd.long != "" && trmWord == cmd.long {
				c.ct, c.expArgs = cmd.ct, cmd.expArgs
				return nil
			}
		}
		return errors.New("there is no such option")
	}
}

/*
	ParseArg adding next argument (if it's possible)
	to cliCommand gotArgs slice
*/
func (c *cliCommand) ParseArg(word string) error {
	if strings.Contains(c.expArgs, "...") {
		// todo: implement support of multiple args
		return errors.New("sorry, multiple arguments are not implemented yet")
	}
	if len(c.gotArgs) != 0 {
		return errors.New("more than one argument was sent")
	}
	switch c.expArgs {
	case "%s":
		c.gotArgs = append(c.gotArgs, word)
	case "%t":
		var s bool
		_, err := fmt.Fscanf(strings.NewReader(word), "%t", s)
		if err != nil {
			return err
		}
		c.gotArgs = append(c.gotArgs, s)
	case "%d":
		var s int
		_, err := fmt.Fscanf(strings.NewReader(word), "%d", s)
		if err != nil {
			return err
		}
		c.gotArgs = append(c.gotArgs, s)
	}
	return nil
}

func parseCommand(s string) ([]*cliCommand, error) {
	parsed := make([]*cliCommand, 0, 3)
	for _, word := range strings.Split(s, " ") {
		var err error
		cmd := &cliCommand{}

		if word[0] == '-' {
			err = cmd.ParseCommand(word)
		} else {
			err = cmd.ParseArg(word)
		}
		if err != nil {
			return nil, err
		}

		parsed = append(parsed, cmd)
	}
	return parsed, nil
}
