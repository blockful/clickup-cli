package api

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Color          string `json:"color"`
	ProfilePicture string `json:"profilePicture"`
	Initials       string `json:"initials"`
}

type UserResponse struct {
	User User `json:"user"`
}

func (c *Client) GetUser() (*UserResponse, error) {
	var resp UserResponse
	if err := c.Do("GET", "/v2/user", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
