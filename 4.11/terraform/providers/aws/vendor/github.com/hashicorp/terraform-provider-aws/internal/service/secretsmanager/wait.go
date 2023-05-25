package secretsmanager

import (
	"time"
)

const (
	// Maximum amount of time to wait for Secrets Manager changes to propagate
	PropagationTimeout = 2 * time.Minute
)
