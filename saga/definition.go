package saga

// SagaDefinition is the interface that wraps the basic SagaDefinition method.
type SagaDefinition struct{}

func (s *SagaDefinition) startSagaOrchestrator(state *SagaState) error {
	return nil
}

type DefinitionBuilder struct{}

// NewDefinitionBuilder returns a new DefinitionBuilder.
func NewDefinitionBuilder() *DefinitionBuilder {
	return &DefinitionBuilder{}
}

func (d *DefinitionBuilder) Step() *DefinitionBuilder {
	return d
}

func (d *DefinitionBuilder) WithCompensation() *DefinitionBuilder {
	return d
}

func (d *DefinitionBuilder) InvokeParticipant() *DefinitionBuilder {
	return d
}

func (d *DefinitionBuilder) OnReply() *DefinitionBuilder {
	return d
}

// Build returns a new SagaDefinition.
func (d *DefinitionBuilder) Build() *SagaDefinition {
	return nil
}
