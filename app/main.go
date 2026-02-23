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

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)
		parts := strings.Fields(command)

		if len(parts) == 0 {
			continue
		}

		implementedCommands := []string{"echo", "exit", "type"}

		switch parts[0] {
		case "echo":
			fmt.Println(strings.Join(parts[1:], " "))
		case "exit":
			os.Exit(0)
		case "type":
			if len(parts) < 2 {
				fmt.Println("Usage: type <command>")
				continue
			}
			cmd := parts[1]
			if contains(implementedCommands, cmd) {
				fmt.Printf("%s is a shell builtin\n", cmd)
			}
			fmt.Println(cmd + ": command not found")

		default:
			fmt.Println(command + ": command not found")
		}
	}
}

func contains(implementedCommands []string, cmd string) bool {
	for _, c := range implementedCommands {
		if c == cmd {
			return true
		}
	}
	return false
}
