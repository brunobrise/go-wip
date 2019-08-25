package wip

// Attachment is an uploaded file (seems to define an image)
type Attachment struct {
	ID          string  `json:"id"`
	URL         string  `json:"url"`
	AspectRatio float32 `json:"aspect_ratio"`
	Filename    string  `json:"filename"`
	Size        uint32  `json:"size"`
	MimeType    string  `json:"mime_type"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	CreatedBy   string  `json:"user"`
}
