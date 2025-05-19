package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"simplewal/internal/wal"
)

// usage prints the CLI usage instructions and exits.
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <a> <b>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Adds two integers using the WAL package.\n")
	fmt.Fprintf(os.Stderr, "Example: %s 2 3\n", os.Args[0])
	os.Exit(1)
}

func main() {
	// Define flags
	help := flag.Bool("help", false, "Show usage information")
	flag.Parse()

	// Show usage if -help is specified or args are invalid
	if *help || len(flag.Args()) != 2 {
		usage()
	}

	// Parse command-line arguments
	a, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid first argument %q: %v\n", flag.Args()[0], err)
		usage()
	}
	b, err := strconv.Atoi(flag.Args()[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid second argument %q: %v\n", flag.Args()[1], err)
		usage()
	}

	// Call the Add function from the wal package
	result := wal.Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, result)
}
