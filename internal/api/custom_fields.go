package api

import "fmt"

// CustomField is defined in tasks.go

type CustomFieldsResponse struct {
	Fields []CustomField `json:"fields"`
}

type SetCustomFieldRequest struct {
	Value interface{} `json:"value"`
}

func (c *Client) GetListCustomFields(listID string) (*CustomFieldsResponse, error) {
	var resp CustomFieldsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/list/%s/field", listID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetFolderCustomFields(folderID string) (*CustomFieldsResponse, error) {
	var resp CustomFieldsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/folder/%s/field", folderID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetSpaceCustomFields(spaceID string) (*CustomFieldsResponse, error) {
	var resp CustomFieldsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/space/%s/field", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetWorkspaceCustomFields(teamID string) (*CustomFieldsResponse, error) {
	var resp CustomFieldsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/team/%s/field", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SetCustomFieldValue(taskID, fieldID string, req *SetCustomFieldRequest) error {
	return c.Do("POST", fmt.Sprintf("/v2/task/%s/field/%s", taskID, fieldID), req, nil)
}

func (c *Client) RemoveCustomFieldValue(taskID, fieldID string) error {
	return c.Do("DELETE", fmt.Sprintf("/v2/task/%s/field/%s", taskID, fieldID), nil, nil)
}
