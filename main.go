package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/amisadman/unfinished-dream/engine"
)

func main() {
	fmt.Println("SQL Engine v0.0.1")
	fmt.Println(strings.Repeat("-", 65))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("somoynosto> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		text = strings.TrimSpace(text)

		if text == "exit" || text == "quit" {
			fmt.Println("Waking up... Goodbye.")
			break
		}

		if text == "" {
			continue
		}

		stmt, err := engine.Parse(text)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		result := engine.Execute(stmt)
		fmt.Printf("Result: %s\n", result)
	}
}
