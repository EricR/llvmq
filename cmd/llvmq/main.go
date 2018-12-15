package main

import (
	"fmt"
	"github.com/ericr/llvmq/config"
	"github.com/ericr/llvmq/query"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("Usage: llvmq path/to/src.ll \"query\"")
	}
	conf := config.New()

	q, err := query.New(os.Args[1], conf)
	if err != nil {
		return fmt.Errorf("Error creating query: %s", err)
	}

	res, err := q.Exec(os.Args[2])
	if err != nil {
		return fmt.Errorf("Query returned error: %s", err)
	}
	fmt.Printf("\n%s\n", res)

	return nil
}
