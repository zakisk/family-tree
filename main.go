package main

import (
	"fmt"
	"os"

	"github.com/zakisk/family-tree/cmd"
)

func main() {
	// root of cobra commands
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
