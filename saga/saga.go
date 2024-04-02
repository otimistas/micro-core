// Package saga provides a simple implementation of the Saga pattern.
package saga

// Saga is a saga.
type Saga struct {
	SagaState
	SagaDefinition
}

// ExecFirstStep returns the first step of the saga.
func (s *Saga) ExecFirstStep() error {
	return s.SagaDefinition.startSagaOrchestrator(&s.SagaState)
}
