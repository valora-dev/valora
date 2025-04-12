package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/version"
)

func main() {
	// Initialize the application
	cdc := codec.New()

	// Use the Cosmos SDK's server to start your application
	rootCmd := server.NewRootCmd()

	// Add subcommands for start, migrate, etc.
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}