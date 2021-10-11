package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if(len(os.Args) < 2) {
		fmt.Println("no command provided")
		os.Exit(2)
	}
	
	cmd := os.Args[1]
	
	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI BASICS - REMINDERS CLI", "msg for greet command")
		err := greetCmd.Parse((os.Args[2:]))
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Hello & Welcome : %s", *msgFlag)
	case "help":
		fmt.Println("Some help message")
	default:
		fmt.Printf("unknown command : %s\n", cmd)
	
	}
}