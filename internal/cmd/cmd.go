package cmd

import "github.com/spf13/cobra"

type Command struct {
	Handler     func() *cobra.Command
	Subcommands []*Command
	cmd         *cobra.Command
}

func (c *Command) Apply() {
	c.cmd = c.Handler()

	for _, child := range c.Subcommands {
		if child.cmd == nil {
			child.Apply()
		}

		c.cmd.AddCommand(child.cmd)
	}
}

func (c *Command) Execute() error {
	return c.cmd.Execute()
}

type CommandOption func(*Command)

func WithHandler(f func() *cobra.Command) CommandOption {
	return func(c *Command) {
		c.Handler = f
	}
}

func WithSubcommands(sub ...*Command) CommandOption {
	return func(c *Command) {
		if c.Subcommands == nil {
			c.Subcommands = make([]*Command, 0)
		}

		c.Subcommands = append(c.Subcommands, sub...)
	}

}

func New(opts ...CommandOption) *Command {
	c := &Command{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
