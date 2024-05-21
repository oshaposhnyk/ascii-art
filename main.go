package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oshaposhnyk/ascii-art/ascii"
)

func main() {
	flag.Parse()
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}
	input := args[0]
	config := ascii.NewConfig(input)

	if len(args) > 1 {
		config.Template = args[1]
	}
	err := ascii.Run(config, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
