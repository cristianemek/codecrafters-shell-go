package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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

		builtins := map[string]bool{
			"echo": true,
			"exit": true,
			"type": true,
		}

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
			if builtins[cmd] {
				fmt.Printf("%s is a shell builtin\n", cmd)
			} else {
				// Check if the command exists in PATH os.PathListSeparator
				found := false
				dirs := strings.SplitSeq(os.Getenv("PATH"), string(os.PathListSeparator))

				for dir := range dirs {
					//existe
					fullPath := filepath.Join(dir, cmd)
					if info, err := os.Stat(fullPath); err == nil {
						//permisos ejecucion
						if info.Mode().Perm()&0111 != 0 {
							fmt.Printf("%s is %s/%s\n", cmd, dir, cmd)
							found = true
							break
						}
					}
				}
				if !found {
					fmt.Printf("%s: not found\n", cmd)
				}
			}

		default:
			fmt.Println(command + ": command not found")
		}
	}
}
