package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	builtins := make(map[string]bool)
	builtins["echo"] = true
	builtins["exit"] = true
	builtins["type"] = true
	builtins["cd"] = true
	builtins["pwd"] = true


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
			case "type":
				if builtins[splitOutput[1]] {
					fmt.Printf("%v is a shell builtin\n", splitOutput[1])
				} else {
					fmt.Printf("%v: not found\n", splitOutput[1])
				}
			default:
				
				fmt.Printf("%v: command not found\n", output[:len(output)-1])
		}
	}
}
