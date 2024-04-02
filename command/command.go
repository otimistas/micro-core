// Package command provides a simple implementation of the Command message pattern.
package command

// Command is a command base template.
type Command interface {
	GetMsgID() string
	GetVersion() int
}
