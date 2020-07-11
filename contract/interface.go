package contract

import (
	"context"
)

// SagaResource Imperfect Compensation (Saga) distributed transaction resource
type SagaResource interface {
	// ID of this resource in Saga Service Comb
	ID () string
	// Transaction original transaction operation in local
	Transaction(ctx context.Context) error
	// Compensation do the forward or backward recovery
	Compensation(ctx context.Context) error
}
