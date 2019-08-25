package wip

// Todo is a pending or completed task
type Todo struct {
	ID          string       `json:"id,omitempty"`
	Content     string       `json:"body,omitempty"`
	CreatedAt   string       `json:"created_at,omitempty"`
	UpdatedAt   string       `json:"updated_at,omitempty"`
	CompletedAt string       `json:"completed_at,omitempty"`
	CreatedBy   User         `json:"user"`
	Product     Product      `json:"product"`
	Attachments []Attachment `json:"attachments"`
}
