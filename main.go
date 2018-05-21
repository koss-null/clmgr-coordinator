package main

import (
	"flag"
	"fmt"

	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/config"
	"myproj.com/clmgr-coordinator/pkg/cli"
	"myproj.com/clmgr-coordinator/pkg/cluster"
	"myproj.com/clmgr-coordinator/pkg/common"
	"myproj.com/clmgr-coordinator/pkg/rest"
	"os"
)

func startCLI(exit chan interface{}) {
	flag.Parse()
	args := flag.Args()

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
					close(exit)
					return
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

func startCluster(exit chan interface{}) {
	cl := cluster.New()
	errChan := make(chan error)
	go cl.Start(errChan)

	err := <-errChan
	fmt.Println(err)
	close(exit)
}

func main() {
	config.InitConfig()
	if err := common.InitLogger(); err != nil {
		fmt.Println("can't init logger")
		os.Exit(-1)
	}

	exit := make(chan interface{})
	go startCluster(exit)
	go startCLI(exit)
	clnt := rest.NewClient()
	clnt.Start()

	<-exit
	logger.Info("Finishing...")
}
