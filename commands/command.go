package commands

type Command struct {
	Text string
}

func NewCommand(text string) *Command {
	return &Command{Text: text}
}

func (c *Command) Type() string {}

type A_Command interface {
	symbol() string
}

type C_Command interface {
	dest() string
	comp() string
	jump() string
}

type L_Command interface {
	symbol() string
}

func (c *Command) symbol() string {}

func (c *Command) dest() string {}

func (c *Command) comp() string {}

func (c *Command) jump() string {}
