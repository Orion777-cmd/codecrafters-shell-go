package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"os/exec"
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
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("%v\n", dir)
			}
		case "cd":
			if len(splitOutput) < 2 {
				fmt.Println("cd: missing argument")
				break
			}
			err := os.Chdir(splitOutput[1])
			if err != nil {
				fmt.Printf("%v: No such file or directory\n", splitOutput[1])
			}
		default:
			pathVar := os.Getenv("PATH")
			paths := strings.Split(pathVar, ":")
			found := false
			for _, path := range paths {
				fullPath := filepath.Join(path, command)
				info, err := os.Stat(fullPath)
				if err == nil && info.Mode().IsRegular() && info.Mode().Perm()&0111 != 0 {
					found = true
					cmd := exec.Command(fullPath, splitOutput[1:]...)
					cmd.Stdin = os.Stdin
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil {
						fmt.Printf("%v\n", err)
					}
					break
				}
			}
			if !found {
				fmt.Printf("%v: command not found\n", output[:len(output)-1])
			}				
		}
	}
}
