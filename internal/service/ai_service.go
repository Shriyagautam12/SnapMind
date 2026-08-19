package service

import "context"

type ImageMetadata struct {
	Description string
	Category    string
	OCRText     string
	Keywords    []string
}

// AIProvider is the only boundary the rest of the backend talks to for
// AI work. The mobile app never sees this — it only sees upload → indexed.
// Gemini is the initial concrete implementation; swapping providers later
// should never require changes outside this interface's implementation.
type AIProvider interface {
	AnalyzeImage(ctx context.Context, image []byte) (*ImageMetadata, error)
	EmbedText(ctx context.Context, text string) ([]float32, error)
	EmbedImage(ctx context.Context, image []byte) ([]float32, error)
}
