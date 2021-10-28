package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LaunchGUI(interpreter *Interpreter) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.ReplaceAll(input, "\r\n", "")
		response := interpreter.ProcessCommand(input)
		if response.Success {
			fmt.Println("Operation Successful: " + response.Message)
		} else {
			fmt.Println("Operation Failed: " + response.Message)
		}

	}
}
