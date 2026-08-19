package model

import "time"

type IndexingStatus string

const (
	IndexingStatusPending    IndexingStatus = "pending"
	IndexingStatusProcessing IndexingStatus = "processing"
	IndexingStatusIndexed    IndexingStatus = "indexed"
	IndexingStatusFailed     IndexingStatus = "failed"
)

// Screenshot is the server-side record for a user's screenshot.
// The original image is never stored permanently: it passes through
// temporary object storage during indexing and is deleted once
// Description/OCRText/Keywords/Embedding are populated.
type Screenshot struct {
	ID            string
	UserID        string
	DeviceAssetID string
	ContentHash   string

	Description string
	Category    string
	OCRText     string
	Keywords    []string
	Source      string

	Embedding []float32

	IndexingStatus IndexingStatus
	FailureReason  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
