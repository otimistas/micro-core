package saga

import "errors"

var (
	// ErrSagaNotFound is returned when a saga is not found.
	ErrSagaNotFound = errors.New("saga not found")
	// ErrSagaAlreadyExists is returned when a saga already exists.
	ErrSagaAlreadyExists = errors.New("saga already exists")
	// ErrAlreadyExistSagaDefinitionKey is returned when a saga definition key already exists.
	ErrAlreadyExistSagaDefinitionKey = errors.New("saga definition key already exists")
	// ErrSagaDefinitionNotFound is returned when a saga definition is not found.
	ErrSagaDefinitionNotFound = errors.New("saga definition not found")
)
