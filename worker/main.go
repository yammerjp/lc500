package main

import (
	"bufio"
	"fmt"
	"os"

	v8 "rogchap.com/v8go"
)

func main() {
	iso := v8.NewIsolate()
	ctx := v8.NewContext(iso)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Text()

		script, _ := iso.CompileUnboundScript(line, "main.js", v8.CompileOptions{})
		val, err := script.Run(ctx)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}

		fmt.Println(val)

		fmt.Print("> ")
	}
}
