package storage

import "context"

// ObjectStorage is the temporary processing buffer screenshots pass
// through during indexing. It is never used as permanent photo storage:
// callers must delete an object once indexing succeeds or permanently fails.
type ObjectStorage interface {
	GetUploadURL(ctx context.Context, key string) (string, error)
	Download(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}
