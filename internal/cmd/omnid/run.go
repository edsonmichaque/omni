package omnid

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdRun() *cobra.Command {
	var params struct {
		debug  bool
		config string
	}

	cmd := &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("run")
			fmt.Println("debug:", params.debug)

			return nil
		},
	}

	cmd.Flags().BoolVar(&params.debug, "debug", false, "")
	cmd.Flags().StringVar(&params.config, "config", "", "")

	return cmd
}
