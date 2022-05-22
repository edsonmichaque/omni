package cmd

import "github.com/spf13/cobra"

type Command struct {
	Run func() *cobra.Command
	Sub []*Command
	cmd *cobra.Command
}

func (c *Command) Apply() {
	c.cmd = c.Run()

	for _, child := range c.Sub {
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

func Run(f func() *cobra.Command) CommandOption {
	return func(c *Command) {
		c.Run = f
	}
}

func WithSubcommands(sub ...*Command) CommandOption {
	return func(c *Command) {
		if c.Sub == nil {
			c.Sub = make([]*Command, 0)
		}

		c.Sub = append(c.Sub, sub...)
	}

}

func New(opts ...CommandOption) *Command {
	c := &Command{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
