package saga

// SagaManager manages all saga definitions.
type SagaManager struct {
	SagaDefinitions map[string]*SagaDefinition
}

// NewSagaManager returns a new SagaManager.
func NewSagaManager() *SagaManager {
	return &SagaManager{
		SagaDefinitions: make(map[string]*SagaDefinition),
	}
}

// RegisterDefinition registers a new saga definition.
func (s *SagaManager) RegisterDefinition(key string, definition *SagaDefinition) error {
	if _, ok := s.SagaDefinitions[key]; ok {
		return ErrAlreadyExistSagaDefinitionKey
	}
	s.SagaDefinitions[key] = definition
	return nil
}

// ResolveDefinition returns a saga definition by key.
func (s *SagaManager) ResolveDefinition(key string) (*SagaDefinition, error) {
	definition, ok := s.SagaDefinitions[key]
	if !ok {
		return &SagaDefinition{}, ErrSagaDefinitionNotFound
	}
	return definition, nil
}
