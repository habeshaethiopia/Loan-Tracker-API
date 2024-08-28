package domain

import (
	"context"
	"time"
)

type LogEntry struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	EventType   string    `json:"event_type"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id,omitempty"` // Optional, depending on the event type
}
type LogRepository interface {
	// SaveLog saves a log entry to the repository
	SaveLog(ctx context.Context, log LogEntry) error
	// GetLogsByUserID retrieves log entries for a specific user
	GetLogsByUserID(ctx context.Context, userID string) ([]LogEntry, error)
	// GetAllLogs retrieves all log entries
	GetAllLogs(ctx context.Context) ([]LogEntry, error)
	// GetLogsByEventType retrieves log entries for a specific event type
	GetLogsByEventType(ctx context.Context, eventType string) ([]LogEntry, error)
}
