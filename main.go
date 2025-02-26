package main

import (
	"fmt"
	"log"
	"os"

	"github.com/warrenb95/cron-parser/internal/formatter"
	"github.com/warrenb95/cron-parser/internal/parser"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Println("Usage: cronparser '*/15 0 1,15 * 1-5 /usr/bin/find'")
		return
	} else if len(args) > 1 {
		log.Fatalln("Error: Only one argument (cron expression) is allowed.")
	}

	input := args[0]

	// Parse the cron fields.
	fields, err := parser.ParseFields(input)
	if err != nil {
		log.Fatalf("parsing fields: %v", err)
	}

	// Format the parsed fields.
	output := formatter.Format(fields)
	fmt.Println(output)
}
