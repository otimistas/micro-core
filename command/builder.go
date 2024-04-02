package command

type CommandBuilder struct {
	Command Command
	Channel CommandChannel
}

// NewCommandBuilder returns a new CommandBuilder.
func NewCommandBuilder() *CommandBuilder {
	return &CommandBuilder{}
}

func (c *CommandBuilder) ForCommand(command Command) *CommandBuilder {
	c.Command = command
	return c
}

func (c *CommandBuilder) WithChannel(channel CommandChannelKey) (*CommandBuilder, error) {
	ch, err := ChannelManagerInstance.ResolveChannel(channel)
	if err != nil {
		return nil, err
	}
	c.Channel = ch
	return c, nil
}

func (c *CommandBuilder) Send() error {
	return nil
}
