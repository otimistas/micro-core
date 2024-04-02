package command

// CommandChannel is a command channel.
type CommandChannel string

// CommandChannelKey is a command channel key.
type CommandChannelKey string

// ChannelManager is a channel manager.
type ChannelManager struct {
	Channels map[CommandChannelKey]CommandChannel
}

// RegisterChannel registers a new channel.
func (c *ChannelManager) RegisterChannel(key CommandChannelKey, channel CommandChannel) error {
	if _, ok := c.Channels[key]; ok {
		return ErrAlreadyExistChannelKey
	}
	c.Channels[key] = channel
	return nil
}

// ResolveChannel returns a channel by key.
func (c *ChannelManager) ResolveChannel(key CommandChannelKey) (CommandChannel, error) {
	channel, ok := c.Channels[key]
	if !ok {
		return "", ErrChannelNotFound
	}
	return channel, nil
}

// NewChannelManager returns a new ChannelManager.
func NewChannelManager() *ChannelManager {
	return &ChannelManager{
		Channels: make(map[CommandChannelKey]CommandChannel),
	}
}

// channelManagerInstance is a channel manager instance.
var channelManagerInstance = NewChannelManager()
