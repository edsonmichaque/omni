package main

import (
	"fmt"
	"os"

	"github.com/edsonmichaque/omni/internal/cmd/omnictl"
)

func main() {
	if err := omnictl.New("omnictl"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
