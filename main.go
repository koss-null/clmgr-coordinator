package main

import (
	"flag"

	"fmt"
	"myproj.com/clmgr-coordinator/pkg/cli"
)

func main() {
	flag.Parse()
	// args := flag.Args()
	// todo: we do need some fixes here to make main as easy to read as possible
	if true {
		c := cli.NewCLI()
		err, control := c.Start()
		for {
			select {
			case cErr := <-err:
				fmt.Printf("The error have been accured, err %s\n", cErr.Error())
				close(control)
			}
		}
	}
}
