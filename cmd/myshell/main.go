package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		output, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		splitOutput := strings.Fields(output)
		if len(splitOutput) == 0 {
			return
		}
		command := splitOutput[0]
		switch command {
			case "exit":
				return
			case "echo":
				fmt.Printf("%v\n", strings.Join(splitOutput[1:], " "))
			default:
				fmt.Printf("%v: command not found\n", output[:len(output)-1])
		}
	}
}
