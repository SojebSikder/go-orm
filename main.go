package main

import (
	"fmt"
	"os"

	"github.com/sojebsikder/go-orm/generator"
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

	// Convert to ast
	ast, err := parser.ParseSchema(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(3)
	}

	// ast to sql
	g := generator.NewPostgreSQLGenerator(ast)

	// Open file for writing SQL
	outFile, err := os.Create("schema.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating sql file: %v\n", err)
		os.Exit(6)
	}
	defer outFile.Close()

	if _, err := outFile.WriteString(g.Generate()); err != nil {
		fmt.Fprintf(os.Stderr, "error writing sql file: %v\n", err)
		os.Exit(7)
	}

	fmt.Println("Schema saved to schema.sql")
}
