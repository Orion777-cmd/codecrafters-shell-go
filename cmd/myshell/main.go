package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		output, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		command := output[:len(output)-1]
		switch command {
			case "exit 0":
				return
			default:
				fmt.Printf("%v: command not found\n", output[:len(output)-1])
		}
	}
}
