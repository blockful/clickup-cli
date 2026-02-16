package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// Doc represents a ClickUp Doc (v3 API).
type Doc struct {
	ID          string      `json:"id"`
	Name        string      `json:"name,omitempty"`
	WorkspaceID json.Number `json:"workspace_id,omitempty"`
	Parent      interface{} `json:"parent,omitempty"`
	Creator     interface{} `json:"creator,omitempty"`
	DateCreated json.Number `json:"date_created,omitempty"`
	Deleted     bool        `json:"deleted,omitempty"`
	Visibility  string      `json:"visibility,omitempty"`
}

type DocsResponse struct {
	Docs []Doc `json:"docs"`
}

type CreateDocRequest struct {
	Name       string     `json:"name"`
	Parent     *DocParent `json:"parent,omitempty"`
	Visibility string     `json:"visibility,omitempty"`
}

type DocParent struct {
	ID   string `json:"id"`
	Type int    `json:"type"`
}

type DocPage struct {
	ID          string      `json:"id"`
	Name        string      `json:"name,omitempty"`
	Content     string      `json:"content,omitempty"`
	ContentHtml string      `json:"content_html,omitempty"`
	OrderIndex  interface{} `json:"orderindex,omitempty"`
	DateCreated json.Number `json:"date_created,omitempty"`
	DateUpdated json.Number `json:"date_updated,omitempty"`
	Archived    bool        `json:"archived,omitempty"`
	Protected   bool        `json:"protected,omitempty"`
	CreatorID   interface{} `json:"creator_id,omitempty"`
}

type DocPagesResponse struct {
	Pages []DocPage `json:"pages"`
}

type CreatePageRequest struct {
	Name         string `json:"name"`
	Content      string `json:"content,omitempty"`
	ContentHtml  string `json:"content_html,omitempty"`
	OrderIndex   *int   `json:"orderindex,omitempty"`
	ParentPageID string `json:"parent_page_id,omitempty"`
}

type EditPageRequest struct {
	Name        string `json:"name,omitempty"`
	Content     string `json:"content,omitempty"`
	ContentHtml string `json:"content_html,omitempty"`
	Archived    *bool  `json:"archived,omitempty"`
	Protected   *bool  `json:"protected,omitempty"`
}

type SearchDocsOptions struct {
	WorkspaceID string
}

func (c *Client) CreateDoc(ctx context.Context, workspaceID string, req *CreateDocRequest) (*Doc, error) {
	var resp Doc
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v3/workspaces/%s/docs", workspaceID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SearchDocs(ctx context.Context, workspaceID string) (*DocsResponse, error) {
	var resp DocsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v3/workspaces/%s/docs", workspaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetDoc(ctx context.Context, workspaceID, docID string) (*Doc, error) {
	var resp Doc
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v3/workspaces/%s/docs/%s", workspaceID, docID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreatePage(ctx context.Context, workspaceID, docID string, req *CreatePageRequest) (*DocPage, error) {
	var resp DocPage
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v3/workspaces/%s/docs/%s/pages", workspaceID, docID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPage(ctx context.Context, workspaceID, docID, pageID string) (*DocPage, error) {
	var resp DocPage
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v3/workspaces/%s/docs/%s/pages/%s", workspaceID, docID, pageID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) EditPage(ctx context.Context, workspaceID, docID, pageID string, req *EditPageRequest) (*DocPage, error) {
	var resp DocPage
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v3/workspaces/%s/docs/%s/pages/%s", workspaceID, docID, pageID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetDocPageListing(ctx context.Context, workspaceID, docID string) (*DocPagesResponse, error) {
	var resp DocPagesResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v3/workspaces/%s/docs/%s/page_listing", workspaceID, docID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
