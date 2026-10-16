package job

import (
	"crypto/rand"
	"time"
)

type Status string

const (
	StatusQueued   Status = "queued"
	StatusRunning  Status = "running"
	StatusDone     Status = "done"
	StatusFailed   Status = "failed"
	StatusCanceled Status = "canceled"
)

func (s Status) Final() bool {
	return s == StatusDone || s == StatusFailed || s == StatusCanceled
}

type Job struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	Status    Status    `json:"status"`
	Attempts  int       `json:"attempts"`
	Result    string    `json:"result,omitempty"`
	Error     string    `json:"error,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(typ, payload string) Job {
	now := time.Now()
	return Job{
		ID:        rand.Text(),
		Type:      typ,
		Payload:   payload,
		Status:    StatusQueued,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
