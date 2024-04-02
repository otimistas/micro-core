package microcore

// Aggregate is the interface that wraps the basic Aggregate method.
type Aggregate struct {
	// The aggregate ID
	ID string
	// The aggregate version
	Version int
	// The aggregate events
	Events []Event
}
