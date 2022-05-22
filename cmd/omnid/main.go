package main

import (
	"fmt"
	"os"

	"github.com/edsonmichaque/omni/internal/cmd/omnid"
)

func main() {
	if err := omnid.New("omnid"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
