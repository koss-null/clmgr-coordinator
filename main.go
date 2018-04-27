package main

import (
	"flag"

	"fmt"
	"myproj.com/clmgr-coordinator/pkg/cli"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	// todo: we do need some fixes here to make main as easy to read as possible

	fmt.Println("Starting clmgr-coordinator")

	if len(args) > 0 {
		c := cli.NewCLI()
		if args[0] == "interact" {
			err, control := c.Start()
			for {
				select {
				case cErr := <-err:
					fmt.Printf("The error have been accured, err %s\n", cErr.Error())
					close(control)
				case <-control:
					fmt.Println("CLI if going off")
					os.Exit(0)
				}
			}
		} else {
			err := c.Exec(args)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
