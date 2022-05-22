package omnictl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdLogin() *cobra.Command {
	var params struct {
		debug  bool
		config string
	}

	cmd := &cobra.Command{
		Use: "login",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("login")
			fmt.Println("debug:", params.debug)

			return nil
		},
	}

	cmd.Flags().BoolVar(&params.debug, "debug", false, "")
	cmd.Flags().StringVar(&params.config, "config", "", "")

	return cmd
}
