package wip

import (
	"fmt"
	"log"
	"time"

	"github.com/machinebox/graphql"
)

// Todo is a pending or completed task
type Todo struct {
	ID          string       `json:"id,omitempty"`
	Body        string       `json:"body,omitempty"`
	CreatedAt   string       `json:"created_at,omitempty"`
	UpdatedAt   string       `json:"updated_at,omitempty"`
	CompletedAt string       `json:"completed_at,omitempty"`
	CreatedBy   User         `json:"user"`
	Product     Product      `json:"product"`
	Attachments []Attachment `json:"attachments"`
}

// CreateTodoResult is the resulted from creating a todo
type CreateTodoResult struct {
	Todo Todo `json:"createTodo,omitempty"`
}

// CreateTodo creates a todo for authenticated user
func (c *Client) CreateTodo(todo Todo) (Todo, error) {
	completedAt := ""
	if todo.CompletedAt != "" {
		completedAt = fmt.Sprintf(`completed_at: "%s"`,
			time.Now().UTC().Format("2006-01-02T15:04:05.999Z"))
	}
	req := graphql.NewRequest(fmt.Sprintf(`mutation {
		createTodo (input: { 
			body: "%s",
			%s,}) {
				id
				body
				created_at
				updated_at
				completed_at
	} }`, todo.Body, completedAt))

	var res CreateTodoResult
	err := c.do(req, &res)
	if err != nil {
		log.Println(err)
		return Todo{}, err
	}

	return res.Todo, nil
}
