package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		//read user input
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		isEcho := strings.HasPrefix(command, "echo ")
		if isEcho {
			fmt.Println(command[5:])
			continue
		}
		switch command {
		case "exit":
			os.Exit(0)
		}

		fmt.Println(command + ": command not found")
	}
}
