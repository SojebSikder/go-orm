package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sojebsikder/go-orm/parser"
)

func main() {
	src, err := parser.ReadAllFromFileOrStdin(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading schema: %v\n", err)
		os.Exit(2)
	}

	ast, err := parser.ParseSchema(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(3)
	}

	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", "  ")
	// if err := enc.Encode(ast); err != nil {
	// 	fmt.Fprintf(os.Stderr, "json encode error: %v\n", err)
	// 	os.Exit(4)
	// }

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
