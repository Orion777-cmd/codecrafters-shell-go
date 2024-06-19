package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
				pathVar := os.Getenv("PATH")
				paths := strings.Split(pathVar, ":")
				found := false
				for _, path := range paths {
					fullPath := filepath.Join(path, splitOutput[1])
					info, err := os.Stat(fullPath)
					if err == nil && info.Mode().IsRegular() && info.Mode().Perm()&0111 != 0 {
						fmt.Printf("%v is %v\n", splitOutput[1], fullPath)
						found = true
						break
					}
				}
				if !found {

					fmt.Printf("%v: not found\n", splitOutput[1])
				}
			}
		default:

			fmt.Printf("%v: command not found\n", output[:len(output)-1])
		}
	}
}
