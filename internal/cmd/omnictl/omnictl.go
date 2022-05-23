package omnictl

import (
	"github.com/edsonmichaque/omni/internal/cmd"
	"github.com/spf13/cobra"
)

func New(name string) error {
	cmd := cmd.New(
		cmd.WithHandler(func() *cobra.Command {
			return &cobra.Command{
				Use: name,
			}
		}),
		cmd.WithChildren(
			cmd.New(cmd.WithHandler(CmdLogin)),
		),
	)

	cmd.Apply()
	return cmd.Execute()
}
