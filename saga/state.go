package saga

// CorrelationID is a unique identifier for a saga.
type CorrelationID any

// SagaState is the interface that wraps the basic SagaState method.
type SagaState interface {
	GetCorrelationID() CorrelationID
}
