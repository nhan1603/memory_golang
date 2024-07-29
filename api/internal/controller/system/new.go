package system

import (
	"context"

	"memory_golang/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// CheckReadiness checks if the system is ready for operation or not
	CheckReadiness(ctx context.Context) error
	MonitorMap()
	MonitorChannel()
	MonitorGoroutine()
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
