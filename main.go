package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sojebsikder/go-orm/parser"
)

var version = "0.1.0"
var appName = "go-orm"

func showUsage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s run <file>\n\n", appName)
	fmt.Printf("  %s help\n", appName)
	fmt.Printf("  %s version\n", appName)
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "run":
		run(os.Args[2:])
	case "help":
		showUsage()
	case "version":
		fmt.Printf("%s version %s\n", appName, version)
	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Printf("Use '%s help' to see available commands.", appName)
		os.Exit(1)
	}
}

func run(args []string) {
	src, err := parser.ReadAllFromFileOrStdin(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading schema: %v\n", err)
		os.Exit(2)
	}

	ast, err := parser.ParseSchema(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(3)
	}

	// Open file for writing JSON
	outFile, err := os.Create("schema.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating json file: %v\n", err)
		os.Exit(5)
	}
	defer outFile.Close()

	enc := json.NewEncoder(outFile)
	enc.SetIndent("", "  ")
	if err := enc.Encode(ast); err != nil {
		fmt.Fprintf(os.Stderr, "json encode error: %v\n", err)
		os.Exit(4)
	}

	fmt.Println("Schema saved to schema.json")
}
