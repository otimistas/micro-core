package microcore

// Event is the interface that wraps the basic Event method.
type Event interface {
	// Event returns the event name.
	Event() string
}

// DomainEventPublisher is the interface that wraps the basic DomainEventPublisher method.
type DomainEventPublisher interface {
	Publish(event Event)
}

// DomainEventSubscriber is the interface that wraps the basic DomainEventSubscriber method.
type DomainEventSubscriber interface {
	Subscribe(event Event)
}

// DomainEventBus is the interface that wraps the basic DomainEventBus method.
type DomainEventBus interface {
	Publish(event Event)
	Subscribe(event Event)
}

// DomainEvent is the struct that wraps the basic DomainEvent method.
type DomainEvent struct {
	// The event name
	Name string
}
