package omnid

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
			cmd.New(cmd.WithHandler(CmdMigrate)),
			cmd.New(cmd.WithHandler(CmdRun)),
			cmd.New(cmd.WithHandler(CmdRestore)),
		),
	)

	cmd.Apply()
	return cmd.Execute()
}
