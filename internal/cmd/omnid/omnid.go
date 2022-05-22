package omnid

import (
	"github.com/edsonmichaque/omni/internal/cmd"
	"github.com/spf13/cobra"
)

func New(name string) error {
	cmd := cmd.New(
		cmd.Run(func() *cobra.Command {
			return &cobra.Command{
				Use: name,
			}
		}),
		cmd.WithSubcommands(
			cmd.New(cmd.Run(CmdMigrate)),
			cmd.New(cmd.Run(CmdRun)),
			cmd.New(cmd.Run(CmdRestore)),
		),
	)

	cmd.Apply()
	return cmd.Execute()
}
