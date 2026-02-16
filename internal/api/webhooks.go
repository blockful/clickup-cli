package api

import (
	"context"
	"fmt"
)

type Webhook struct {
	ID       string      `json:"id"`
	UserID   int         `json:"userid"`
	TeamID   int         `json:"team_id"`
	Endpoint string      `json:"endpoint"`
	ClientID string      `json:"client_id"`
	Events   interface{} `json:"events"`
	TaskID   interface{} `json:"task_id"`
	ListID   interface{} `json:"list_id"`
	FolderID interface{} `json:"folder_id"`
	SpaceID  interface{} `json:"space_id"`
	Health   interface{} `json:"health"`
	Secret   string      `json:"secret"`
}

type WebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type CreateWebhookRequest struct {
	Endpoint string   `json:"endpoint"`
	Events   []string `json:"events"`
	SpaceID  *int     `json:"space_id,omitempty"`
	FolderID *int     `json:"folder_id,omitempty"`
	ListID   *int     `json:"list_id,omitempty"`
	TaskID   *string  `json:"task_id,omitempty"`
}

type CreateWebhookResponse struct {
	ID      string  `json:"id"`
	Webhook Webhook `json:"webhook"`
}

type UpdateWebhookRequest struct {
	Endpoint string `json:"endpoint"`
	Events   string `json:"events"`
	Status   string `json:"status"`
}

type UpdateWebhookResponse struct {
	ID      string  `json:"id"`
	Webhook Webhook `json:"webhook"`
}

func (c *Client) GetWebhooks(ctx context.Context, teamID string) (*WebhooksResponse, error) {
	var resp WebhooksResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/webhook", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateWebhook(ctx context.Context, teamID string, req *CreateWebhookRequest) (*CreateWebhookResponse, error) {
	var resp CreateWebhookResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/webhook", teamID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateWebhook(ctx context.Context, webhookID string, req *UpdateWebhookRequest) (*UpdateWebhookResponse, error) {
	var resp UpdateWebhookResponse
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/webhook/%s", webhookID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteWebhook(ctx context.Context, webhookID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/webhook/%s", webhookID), nil, nil)
}
