package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
			"pwd":  true,
		}

		switch parts[0] {
		case "echo":
			fmt.Println(strings.Join(parts[1:], " "))
		case "exit":
			os.Exit(0)
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error getting current directory:", err)
			} else {
				fmt.Println(dir)
			}
		case "type":
			if len(parts) < 2 {
				fmt.Println("Usage: type <command>")
				continue
			}
			cmd := parts[1]
			if builtins[cmd] {
				fmt.Printf("%s is a shell builtin\n", cmd)
			} else {
				if fullPath, found := searchInPath(cmd); found {
					if isExecutable(fullPath) {
						fmt.Printf("%s is %s\n", cmd, fullPath)
					}
				} else {
					fmt.Printf("%s: not found\n", cmd)
				}
			}

		default:
			cmd, found := searchInPath(parts[0])
			if !found {
				fmt.Printf("%s: command not found\n", parts[0])
				continue
			}
			execCmd := exec.Command(cmd, parts[1:]...)
			execCmd.Args[0] = parts[0]
			execCmd.Stdin = os.Stdin   //a la salida del comando le asignamos la entrada estandar teclado
			execCmd.Stdout = os.Stdout //a la salida del comando le asignamos la salida estandar pantalla
			execCmd.Stderr = os.Stderr //salida de errores
			if err := execCmd.Run(); err != nil {
				fmt.Fprintln(os.Stderr, "Error executing command:", err)
			}

		}
	}
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0111 != 0
}

func searchInPath(cmd string) (string, bool) {
	dirs := strings.SplitSeq(os.Getenv("PATH"), string(os.PathListSeparator))

	for dir := range dirs {
		//existe
		fullPath := filepath.Join(dir, cmd)
		if _, err := os.Stat(fullPath); err == nil {
			if isExecutable(fullPath) {
				return fullPath, true
			}
		}
	}
	return "", false
}
