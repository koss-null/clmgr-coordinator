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
	}
)

func (cli *commandLineInterface) Start() (<-chan error, chan interface{}) {
	errChan, done := make(chan error), make(chan interface{})
	go func() {
		reader := bufio.NewReader(os.Stdin)
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
				commands, err := parseCommand(line)
				if err != nil {
					fmt.Printf(err.Error())
					continue
				}
				// todo: implement command performer
				for i := range commands {
					err := PerformCommand(commands[i])
					if err != nil {
						errChan <- err
					}
				}
			}
		}
	}()

	return errChan, done
}
