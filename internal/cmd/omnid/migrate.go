package omnid

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdMigrate() *cobra.Command {
	var params struct {
		debug  bool
		config string
	}

	cmd := &cobra.Command{
		Use: "migrate",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("migrate")
			fmt.Println("debug:", params.debug)

			return nil
		},
	}

	cmd.Flags().BoolVar(&params.debug, "debug", false, "")
	cmd.Flags().StringVar(&params.config, "config", "", "")

	return cmd
}
