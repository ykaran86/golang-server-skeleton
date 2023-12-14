package main

import (
	"golang-server-skeleton/cmd"
	"golang-server-skeleton/config"
	"log"
)

func main() {
	config.LoadConfig()

	cli := cmd.NewCLICommand()
	err := cli.Execute()
	if err != nil {
		log.Fatalf("failed exeucting command, err: %v", err)
	}
}
