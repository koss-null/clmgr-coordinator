package cli

import (
	"bufio"
	"fmt"
	"os"
)

type (
	commandLineInterface struct{}

	CLI interface {
		Start() (<-chan error, chan interface{})
		Exec([]string) error
	}
)

func NewCLI() CLI {
	return &commandLineInterface{}
}

// todo: remove this stub func asap
func PerformCommand(cmd *cliCommand, done chan interface{}) error {
	fmt.Printf("Performing command <stub>: %s\n", cmd.long)
	if cmd.ct == Exit {
		close(done)
		return nil
	}
	return nil
}

func (cli *commandLineInterface) Start() (<-chan error, chan interface{}) {
	errChan, done := make(chan error), make(chan interface{})
	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("ClusterManager CLI started\nTo exit press ctrl+c or type '--exit'")
		for {
			select {
			case <-done:
				return
			default:
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Printf(string(inputErr), err.Error())
					continue
				}
				// removing "\n"
				line = line[:len(line)-1]
				if len(line) == 0 {
					continue
				}
				commands, err := parseCommandLine(line)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				// todo: implement command performer
				for i := range commands {
					err := PerformCommand(commands[i], done)
					if err != nil {
						errChan <- err
						close(done)
					}
				}
			}
		}
	}()

	return errChan, done
}

func (cli *commandLineInterface) Exec(args []string) error {
	commands, err := parseCommandList(args)
	if err != nil {
		return err
	}
	for i := range commands {
		err := PerformCommand(commands[i], make(chan interface{}))
		if err != nil {
			return err
		}
	}
	return nil
}
