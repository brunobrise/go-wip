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

// CreateTodo creates a todo (authenticated user)
func (c *Client) CreateTodo(todo Todo) (Todo, error) {
	completedAt := ""
	if todo.CompletedAt != "" {
		completedAt = fmt.Sprintf(`completed_at: "%s"`,
			time.Now().UTC().Format("2006-01-02T15:04:05.999Z"))
	}
	req := graphql.NewRequest(fmt.Sprintf(`
	mutation {
		createTodo (input: { body: "%s", %s,}) { 
			id
			body
			created_at
			updated_at
			completed_at
		}
	}`, todo.Body, completedAt))

	var res struct {
		Todo Todo `json:"createTodo,omitempty"`
	}
	err := c.do(req, &res)
	if err != nil {
		log.Println(err)
		return Todo{}, err
	}

	return res.Todo, nil
}

// CompleteTodo completes a todo (authenticated user)
func (c *Client) CompleteTodo(todo Todo) (Todo, error) {
	req := graphql.NewRequest(fmt.Sprintf(`
	mutation {
		completeTodo (id: %s) {
			id
			body
			created_at
			updated_at
			completed_at
		}
	}`, todo.ID))

	var res struct {
		Todo Todo `json:"completeTodo,omitempty"`
	}
	err := c.do(req, &res)
	if err != nil {
		log.Println(err)
		return Todo{}, err
	}

	return res.Todo, nil
}

// CompleteTodo completes a todo (authenticated user)
func (c *Client) UncompleteTodo(todo Todo) (Todo, error) {
	req := graphql.NewRequest(fmt.Sprintf(`
	mutation {
		uncompleteTodo (id: %s) {
			id
			body
			created_at
			updated_at
			completed_at
		}
	}`, todo.ID))

	var res struct {
		Todo Todo `json:"uncompleteTodo,omitempty"`
	}
	err := c.do(req, &res)
	if err != nil {
		log.Println(err)
		return Todo{}, err
	}

	return res.Todo, nil
}
