package main

import "os"

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
