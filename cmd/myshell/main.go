package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	output, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%v: command not found\n", output[:len(output)-1])
}
