package wip

// Product is as maker baby
type Product struct {
	ID         string `json:"id,omitempty"`
	Hashtag    string `json:"hashtag,omitempty"`
	Name       string `json:"name,omitempty"`
	Pitch      string `json:"pitch,omitempty"`
	URL        string `json:"url,omitempty"`
	WebsiteURL string `json:"website_url,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	Todos      []Todo `json:"todos,omitempty"`
	Makers     []User `json:"makers,omitempty"`
}
