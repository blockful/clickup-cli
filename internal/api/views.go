package api

import "fmt"

type View struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Parent   interface{} `json:"parent,omitempty"`
	Grouping interface{} `json:"grouping,omitempty"`
	Divide   interface{} `json:"divide,omitempty"`
	Sorting  interface{} `json:"sorting,omitempty"`
	Filters  interface{} `json:"filters,omitempty"`
	Columns  interface{} `json:"columns,omitempty"`
	Settings interface{} `json:"settings,omitempty"`
}

type ViewsResponse struct {
	Views []View `json:"views"`
}

type ViewResponse struct {
	View View `json:"view"`
}

type ViewTasksResponse struct {
	Tasks    []interface{} `json:"tasks"`
	LastPage bool          `json:"last_page"`
}

type CreateViewRequest struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Grouping    interface{} `json:"grouping,omitempty"`
	Divide      interface{} `json:"divide,omitempty"`
	Sorting     interface{} `json:"sorting,omitempty"`
	Filters     interface{} `json:"filters,omitempty"`
	Columns     interface{} `json:"columns,omitempty"`
	TeamSidebar interface{} `json:"team_sidebar,omitempty"`
	Settings    interface{} `json:"settings,omitempty"`
}

type UpdateViewRequest struct {
	Name        string      `json:"name,omitempty"`
	Type        string      `json:"type,omitempty"`
	Parent      interface{} `json:"parent,omitempty"`
	Grouping    interface{} `json:"grouping,omitempty"`
	Divide      interface{} `json:"divide,omitempty"`
	Sorting     interface{} `json:"sorting,omitempty"`
	Filters     interface{} `json:"filters,omitempty"`
	Columns     interface{} `json:"columns,omitempty"`
	TeamSidebar interface{} `json:"team_sidebar,omitempty"`
	Settings    interface{} `json:"settings,omitempty"`
}

func (c *Client) GetTeamViews(teamID string) (*ViewsResponse, error) {
	var resp ViewsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/team/%s/view", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetSpaceViews(spaceID string) (*ViewsResponse, error) {
	var resp ViewsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/space/%s/view", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetFolderViews(folderID string) (*ViewsResponse, error) {
	var resp ViewsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/folder/%s/view", folderID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetListViews(listID string) (*ViewsResponse, error) {
	var resp ViewsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/list/%s/view", listID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetView(viewID string) (*ViewResponse, error) {
	var resp ViewResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/view/%s", viewID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateTeamView(teamID string, req *CreateViewRequest) (*ViewResponse, error) {
	var resp ViewResponse
	if err := c.Do("POST", fmt.Sprintf("/v2/team/%s/view", teamID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateSpaceView(spaceID string, req *CreateViewRequest) (*ViewResponse, error) {
	var resp ViewResponse
	if err := c.Do("POST", fmt.Sprintf("/v2/space/%s/view", spaceID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateFolderView(folderID string, req *CreateViewRequest) (*ViewResponse, error) {
	var resp ViewResponse
	if err := c.Do("POST", fmt.Sprintf("/v2/folder/%s/view", folderID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateListView(listID string, req *CreateViewRequest) (*ViewResponse, error) {
	var resp ViewResponse
	if err := c.Do("POST", fmt.Sprintf("/v2/list/%s/view", listID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateView(viewID string, req *UpdateViewRequest) (*ViewResponse, error) {
	var resp ViewResponse
	if err := c.Do("PUT", fmt.Sprintf("/v2/view/%s", viewID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteView(viewID string) error {
	return c.Do("DELETE", fmt.Sprintf("/v2/view/%s", viewID), nil, nil)
}

func (c *Client) GetViewTasks(viewID string, page int) (*ViewTasksResponse, error) {
	path := fmt.Sprintf("/v2/view/%s/task", viewID)
	if page > 0 {
		path += fmt.Sprintf("?page=%d", page)
	}
	var resp ViewTasksResponse
	if err := c.Do("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
