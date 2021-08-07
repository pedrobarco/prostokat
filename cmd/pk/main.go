package main

import "fmt"

func main() {
	err := rootCmd.Execute()
	if err != nil {
		panic(fmt.Errorf("Fatal error running cli: %s", err))
	}
}
