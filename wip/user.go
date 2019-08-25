package wip

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

// User is the data struture for a WIP user
type User struct {
	ID                  string    `json:"id,omitempty"`
	Username            string    `json:"username,omitempty"`
	FirstName           string    `json:"first_name,omitempty"`
	LastName            string    `json:"last_name,omitempty"`
	URL                 string    `json:"url,omitempty"`
	AvatarURL           string    `json:"avatar_url,omitempty"`
	TimeZone            string    `json:"time_zone,omitempty"`
	CompletedTodosCount uint32    `json:"completed_todos_count,omitempty"`
	Streaking           bool      `json:"streaking,omitempty"`
	Streak              uint16    `json:"streak,omitempty"`
	BestStreak          uint16    `json:"best_streak,omitempty"`
	Products            []Product `json:"products,omitempty"`
	Todos               []Todo    `json:"todos,omitempty"`
}

// UserData is the data struture for fetching users
type UserData struct {
	User User `json:"user,omitempty"`
}

// ViewerData is the data struture for fetching authenticated user
type ViewerData struct {
	Viewer User `json:"viewer,omitempty"`
}

// GetUser fetches an user with specified id
func (c *Client) GetUser(id string) User {
	req := graphql.NewRequest(fmt.Sprintf(`{ user (id: %s) {
		id
		username
	} }`, id))

	var data UserData
	err := c.do(req, &data)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	return data.User
}

// GetViewer fetches authenticated user
func (c *Client) GetViewer() User {
	req := graphql.NewRequest(`{ viewer {
		id
		username
		first_name
		last_name
		url
		avatar_url
		time_zone
		completed_todos_count
		streaking
		streak
		best_streak
		products {
			id
			hashtag
			name
			pitch
			url
			website_url
			created_at
			updated_at		  
		}
		todos {
			id
			body
			created_at
			updated_at
			completed_at
			attachments {
				id
				url
				aspect_ratio
				filename
				size				
				mime_type
				created_at
				updated_at
			}
			product {
				id
				hashtag
				name
				pitch
				url
			}
		}
	} }`)

	var data ViewerData
	err := c.do(req, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data.Viewer
}
