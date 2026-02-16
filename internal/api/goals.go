package api

import (
	"context"
	"fmt"
)

type Goal struct {
	ID               string      `json:"id"`
	PrettyID         string      `json:"pretty_id,omitempty"`
	Name             string      `json:"name"`
	TeamID           string      `json:"team_id"`
	Creator          interface{} `json:"creator,omitempty"`
	Owner            interface{} `json:"owner,omitempty"`
	Color            string      `json:"color,omitempty"`
	DateCreated      string      `json:"date_created,omitempty"`
	StartDate        interface{} `json:"start_date,omitempty"`
	DueDate          string      `json:"due_date,omitempty"`
	Description      string      `json:"description,omitempty"`
	Private          bool        `json:"private,omitempty"`
	Archived         bool        `json:"archived,omitempty"`
	MultipleOwners   bool        `json:"multiple_owners,omitempty"`
	FolderID         interface{} `json:"folder_id,omitempty"`
	Members          interface{} `json:"members,omitempty"`
	Owners           interface{} `json:"owners,omitempty"`
	KeyResults       interface{} `json:"key_results,omitempty"`
	PercentCompleted float64     `json:"percent_completed,omitempty"`
}

type GoalsResponse struct {
	Goals   []Goal        `json:"goals"`
	Folders []interface{} `json:"folders,omitempty"`
}

type GoalResponse struct {
	Goal Goal `json:"goal"`
}

type CreateGoalRequest struct {
	Name           string `json:"name"`
	DueDate        int64  `json:"due_date"`
	Description    string `json:"description"`
	MultipleOwners bool   `json:"multiple_owners"`
	Owners         []int  `json:"owners"`
	Color          string `json:"color"`
}

type UpdateGoalRequest struct {
	Name        string `json:"name,omitempty"`
	DueDate     *int64 `json:"due_date,omitempty"`
	Description string `json:"description,omitempty"`
	RemOwners   []int  `json:"rem_owners,omitempty"`
	AddOwners   []int  `json:"add_owners,omitempty"`
	Color       string `json:"color,omitempty"`
}

type KeyResult struct {
	ID               string      `json:"id"`
	GoalID           string      `json:"goal_id"`
	Name             string      `json:"name"`
	Type             string      `json:"type"`
	Unit             string      `json:"unit,omitempty"`
	Creator          interface{} `json:"creator,omitempty"`
	DateCreated      string      `json:"date_created,omitempty"`
	GoalPrettyID     string      `json:"goal_pretty_id,omitempty"`
	PercentCompleted float64     `json:"percent_completed,omitempty"`
	Completed        bool        `json:"completed,omitempty"`
	Owners           interface{} `json:"owners,omitempty"`
}

type KeyResultResponse struct {
	KeyResult KeyResult `json:"key_result"`
}

type CreateKeyResultRequest struct {
	Name       string   `json:"name"`
	Owners     []int    `json:"owners"`
	Type       string   `json:"type"`
	StepsStart int      `json:"steps_start"`
	StepsEnd   int      `json:"steps_end"`
	Unit       string   `json:"unit"`
	TaskIDs    []string `json:"task_ids,omitempty"`
	ListIDs    []string `json:"list_ids,omitempty"`
}

func (c *Client) GetGoals(ctx context.Context, teamID string, includeCompleted bool) (*GoalsResponse, error) {
	path := fmt.Sprintf("/v2/team/%s/goal", teamID)
	if includeCompleted {
		path += "?include_completed=true"
	}
	var resp GoalsResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetGoal(ctx context.Context, goalID string) (*GoalResponse, error) {
	var resp GoalResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/goal/%s", goalID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateGoal(ctx context.Context, teamID string, req *CreateGoalRequest) (*GoalResponse, error) {
	var resp GoalResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/goal", teamID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateGoal(ctx context.Context, goalID string, req *UpdateGoalRequest) (*GoalResponse, error) {
	var resp GoalResponse
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/goal/%s", goalID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteGoal(ctx context.Context, goalID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/goal/%s", goalID), nil, nil)
}

func (c *Client) CreateKeyResult(ctx context.Context, goalID string, req *CreateKeyResultRequest) (*KeyResultResponse, error) {
	var resp KeyResultResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/goal/%s/key_result", goalID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
