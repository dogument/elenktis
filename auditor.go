package elenktis

import (
	"context"
	"time"
)

// Auditor is the interface that defines the operations possible for an auditor
type Auditor interface {
	Create(context.Context, AuditLog) error
}

type AuditLog struct {
	ApplicationID string
	Event         string
	EventTime     time.Time
}
