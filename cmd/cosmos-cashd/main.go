package main

import (
	"os"

	"github.com/allinbits/cosmos-cash/cmd/cosmos-cashd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
