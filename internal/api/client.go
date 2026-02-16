package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"
)

const DefaultBaseURL = "https://api.clickup.com/api"

const (
	defaultMaxRetries    = 3
	defaultRetryBaseWait = 1 * time.Second
	maxRetryWait         = 30 * time.Second
)

// ClientInterface defines all ClickUp API operations.
type ClientInterface interface {
	// Auth
	GetUser(ctx context.Context) (*UserResponse, error)

	// Workspaces
	ListWorkspaces(ctx context.Context) (*WorkspacesResponse, error)

	// Spaces
	ListSpaces(ctx context.Context, workspaceID string) (*SpacesResponse, error)
	GetSpace(ctx context.Context, spaceID string) (*Space, error)
	CreateSpace(ctx context.Context, workspaceID string, req *CreateSpaceRequest) (*Space, error)
	UpdateSpace(ctx context.Context, spaceID string, req *UpdateSpaceRequest) (*Space, error)
	DeleteSpace(ctx context.Context, spaceID string) error

	// Folders
	ListFolders(ctx context.Context, spaceID string) (*FoldersResponse, error)
	GetFolder(ctx context.Context, folderID string) (*Folder, error)
	CreateFolder(ctx context.Context, spaceID string, req *CreateFolderRequest) (*Folder, error)
	UpdateFolder(ctx context.Context, folderID string, req *UpdateFolderRequest) (*Folder, error)
	DeleteFolder(ctx context.Context, folderID string) error

	// Lists
	ListLists(ctx context.Context, folderID string) (*ListsResponse, error)
	ListFolderlessLists(ctx context.Context, spaceID string) (*ListsResponse, error)
	GetList(ctx context.Context, listID string) (*List, error)
	CreateList(ctx context.Context, folderID string, req *CreateListRequest) (*List, error)
	CreateFolderlessList(ctx context.Context, spaceID string, req *CreateListRequest) (*List, error)
	UpdateList(ctx context.Context, listID string, req *UpdateListRequest) (*List, error)
	DeleteList(ctx context.Context, listID string) error

	// Tasks
	ListTasks(ctx context.Context, listID string, opts *ListTasksOptions) (*TasksResponse, error)
	GetTask(ctx context.Context, taskID string, opts ...GetTaskOptions) (*Task, error)
	CreateTask(ctx context.Context, listID string, req *CreateTaskRequest) (*Task, error)
	UpdateTask(ctx context.Context, taskID string, req *UpdateTaskRequest, opts ...UpdateTaskOptions) (*Task, error)
	DeleteTask(ctx context.Context, taskID string) error
	SearchTasks(ctx context.Context, teamID string, opts *SearchTasksOptions) (*TasksResponse, error)

	// Comments
	ListComments(ctx context.Context, taskID string) (*CommentsResponse, error)
	ListListComments(ctx context.Context, listID string) (*CommentsResponse, error)
	CreateComment(ctx context.Context, taskID string, req *CreateCommentRequest) (*CreateCommentResponse, error)
	CreateListComment(ctx context.Context, listID string, req *CreateCommentRequest) (*CreateCommentResponse, error)
	UpdateComment(ctx context.Context, commentID string, req *UpdateCommentRequest) error
	DeleteComment(ctx context.Context, commentID string) error

	// Custom Fields
	GetListCustomFields(ctx context.Context, listID string) (*CustomFieldsResponse, error)
	GetFolderCustomFields(ctx context.Context, folderID string) (*CustomFieldsResponse, error)
	GetSpaceCustomFields(ctx context.Context, spaceID string) (*CustomFieldsResponse, error)
	GetWorkspaceCustomFields(ctx context.Context, teamID string) (*CustomFieldsResponse, error)
	SetCustomFieldValue(ctx context.Context, taskID, fieldID string, req *SetCustomFieldRequest) error
	RemoveCustomFieldValue(ctx context.Context, taskID, fieldID string) error

	// Tags
	GetSpaceTags(ctx context.Context, spaceID string) (*TagsResponse, error)
	CreateSpaceTag(ctx context.Context, spaceID string, req *CreateTagRequest) error
	UpdateSpaceTag(ctx context.Context, spaceID, tagName string, req *UpdateTagRequest) error
	DeleteSpaceTag(ctx context.Context, spaceID, tagName string) error
	AddTagToTask(ctx context.Context, taskID, tagName string) error
	RemoveTagFromTask(ctx context.Context, taskID, tagName string) error

	// Checklists
	CreateChecklist(ctx context.Context, taskID string, req *CreateChecklistRequest) (*ChecklistResponse, error)
	EditChecklist(ctx context.Context, checklistID string, req *EditChecklistRequest) error
	DeleteChecklist(ctx context.Context, checklistID string) error
	CreateChecklistItem(ctx context.Context, checklistID string, req *CreateChecklistItemRequest) (*ChecklistResponse, error)
	EditChecklistItem(ctx context.Context, checklistID, checklistItemID string, req *EditChecklistItemRequest) (*ChecklistResponse, error)
	DeleteChecklistItem(ctx context.Context, checklistID, checklistItemID string) error

	// Docs (v3)
	CreateDoc(ctx context.Context, workspaceID string, req *CreateDocRequest) (*Doc, error)
	SearchDocs(ctx context.Context, workspaceID string) (*DocsResponse, error)
	GetDoc(ctx context.Context, workspaceID, docID string) (*Doc, error)
	CreatePage(ctx context.Context, workspaceID, docID string, req *CreatePageRequest) (*DocPage, error)
	GetPage(ctx context.Context, workspaceID, docID, pageID string) (*DocPage, error)
	EditPage(ctx context.Context, workspaceID, docID, pageID string, req *EditPageRequest) (*DocPage, error)
	GetDocPageListing(ctx context.Context, workspaceID, docID string) (*DocPagesResponse, error)

	// Time Tracking
	GetTimeEntries(ctx context.Context, teamID string, opts *ListTimeEntriesOptions) (*TimeEntriesResponse, error)
	CreateTimeEntry(ctx context.Context, teamID string, req *CreateTimeEntryRequest) (*TimeEntry, error)
	GetTimeEntry(ctx context.Context, teamID, timerID string) (*SingleTimeEntryResponse, error)
	UpdateTimeEntry(ctx context.Context, teamID, timerID string, req *UpdateTimeEntryRequest) error
	DeleteTimeEntry(ctx context.Context, teamID, timerID string) error
	StartTimer(ctx context.Context, teamID string, req *StartTimerRequest) (*SingleTimeEntryResponse, error)
	StopTimer(ctx context.Context, teamID string) (*SingleTimeEntryResponse, error)
	GetRunningTimer(ctx context.Context, teamID string, assignee string) (*SingleTimeEntryResponse, error)
	GetTimeEntryTags(ctx context.Context, teamID string) (*TimeEntryTagsResponse, error)

	// Webhooks
	GetWebhooks(ctx context.Context, teamID string) (*WebhooksResponse, error)
	CreateWebhook(ctx context.Context, teamID string, req *CreateWebhookRequest) (*CreateWebhookResponse, error)
	UpdateWebhook(ctx context.Context, webhookID string, req *UpdateWebhookRequest) (*UpdateWebhookResponse, error)
	DeleteWebhook(ctx context.Context, webhookID string) error

	// Views
	GetTeamViews(ctx context.Context, teamID string) (*ViewsResponse, error)
	GetSpaceViews(ctx context.Context, spaceID string) (*ViewsResponse, error)
	GetFolderViews(ctx context.Context, folderID string) (*ViewsResponse, error)
	GetListViews(ctx context.Context, listID string) (*ViewsResponse, error)
	GetView(ctx context.Context, viewID string) (*ViewResponse, error)
	CreateTeamView(ctx context.Context, teamID string, req *CreateViewRequest) (*ViewResponse, error)
	CreateSpaceView(ctx context.Context, spaceID string, req *CreateViewRequest) (*ViewResponse, error)
	CreateFolderView(ctx context.Context, folderID string, req *CreateViewRequest) (*ViewResponse, error)
	CreateListView(ctx context.Context, listID string, req *CreateViewRequest) (*ViewResponse, error)
	UpdateView(ctx context.Context, viewID string, req *UpdateViewRequest) (*ViewResponse, error)
	DeleteView(ctx context.Context, viewID string) error
	GetViewTasks(ctx context.Context, viewID string, page int) (*ViewTasksResponse, error)

	// Goals
	GetGoals(ctx context.Context, teamID string, includeCompleted bool) (*GoalsResponse, error)
	GetGoal(ctx context.Context, goalID string) (*GoalResponse, error)
	CreateGoal(ctx context.Context, teamID string, req *CreateGoalRequest) (*GoalResponse, error)
	UpdateGoal(ctx context.Context, goalID string, req *UpdateGoalRequest) (*GoalResponse, error)
	DeleteGoal(ctx context.Context, goalID string) error
	CreateKeyResult(ctx context.Context, goalID string, req *CreateKeyResultRequest) (*KeyResultResponse, error)

	// Members
	GetListMembers(ctx context.Context, listID string) (*MembersResponse, error)
	GetTaskMembers(ctx context.Context, taskID string) (*MembersResponse, error)

	// Groups
	GetGroups(ctx context.Context, teamID string) (*GroupsResponse, error)
	CreateGroup(ctx context.Context, teamID string, req *CreateGroupRequest) (*Group, error)
	UpdateGroup(ctx context.Context, groupID string, req *UpdateGroupRequest) (*Group, error)
	DeleteGroup(ctx context.Context, groupID string) error

	// Guests
	InviteGuest(ctx context.Context, teamID string, req *InviteGuestRequest) error
	GetGuest(ctx context.Context, teamID, guestID string) (*GuestResponse, error)
	EditGuest(ctx context.Context, teamID, guestID string, req *EditGuestRequest) (*GuestResponse, error)
	RemoveGuest(ctx context.Context, teamID, guestID string) error
}

// Client implements ClientInterface using HTTP requests to the ClickUp API.
type Client struct {
	BaseURL       string
	Token         string
	HTTPClient    *http.Client
	MaxRetries    int
	RetryBaseWait time.Duration
}

var _ ClientInterface = (*Client)(nil)

type APIError struct {
	Err     string `json:"err"`
	ECODE   string `json:"ECODE"`
	Message string `json:"message,omitempty"`
}

// ClientError is a typed error returned by all API methods.
// StatusCode is the HTTP status (0 for non-HTTP errors).
// Code is a machine-readable error code (e.g. UNAUTHORIZED, RATE_LIMITED).
// Retryable indicates whether the caller should retry.
type ClientError struct {
	StatusCode int
	Code       string
	Message    string
	Retryable  bool
}

func (e *ClientError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("[%s] %s (HTTP %d)", e.Code, e.Message, e.StatusCode)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewClient(token string) *Client {
	return &Client{
		BaseURL: DefaultBaseURL,
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		MaxRetries:    defaultMaxRetries,
		RetryBaseWait: defaultRetryBaseWait,
	}
}

// Do executes an HTTP request with automatic retry for 429 and 5xx responses.
func (c *Client) Do(ctx context.Context, method, path string, body, result interface{}) error {
	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return &ClientError{Code: "MARSHAL_ERROR", Message: fmt.Sprintf("failed to marshal request body: %v", err)}
		}
	}

	var lastErr error
	maxAttempts := c.MaxRetries + 1
	if maxAttempts < 1 {
		maxAttempts = 1
	}

	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			wait := c.retryWait(attempt)
			select {
			case <-ctx.Done():
				return &ClientError{Code: "CANCELLED", Message: "request cancelled during retry wait"}
			case <-time.After(wait):
			}
		}

		err := c.doOnce(ctx, method, path, bodyBytes, result)
		if err == nil {
			return nil
		}

		lastErr = err

		// Only retry on retryable errors
		if clientErr, ok := err.(*ClientError); ok && clientErr.Retryable && attempt < maxAttempts-1 {
			continue
		}
		return err
	}

	return lastErr
}

// doOnce performs a single HTTP request attempt.
func (c *Client) doOnce(ctx context.Context, method, path string, bodyBytes []byte, result interface{}) error {
	var reqBody io.Reader
	if bodyBytes != nil {
		reqBody = bytes.NewReader(bodyBytes)
	}

	url := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return &ClientError{Code: "REQUEST_ERROR", Message: fmt.Sprintf("failed to create request: %v", err)}
	}

	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return &ClientError{Code: "CANCELLED", Message: "request cancelled"}
		}
		return &ClientError{Code: "NETWORK_ERROR", Message: fmt.Sprintf("request failed: %v", err), Retryable: true}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ClientError{Code: "READ_ERROR", Message: fmt.Sprintf("failed to read response: %v", err)}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		code := errorCodeFromStatus(resp.StatusCode)
		retryable := resp.StatusCode == 429 || resp.StatusCode >= 500
		var apiErr APIError
		if json.Unmarshal(respBody, &apiErr) == nil && (apiErr.Err != "" || apiErr.Message != "") {
			msg := apiErr.Err
			if msg == "" {
				msg = apiErr.Message
			}
			return &ClientError{StatusCode: resp.StatusCode, Code: code, Message: msg, Retryable: retryable}
		}
		return &ClientError{StatusCode: resp.StatusCode, Code: code, Message: fmt.Sprintf("API returned status %d", resp.StatusCode), Retryable: retryable}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return &ClientError{Code: "UNMARSHAL_ERROR", Message: fmt.Sprintf("failed to parse response: %v", err)}
		}
	}

	return nil
}

// retryWait computes wait duration with exponential backoff + jitter.
func (c *Client) retryWait(attempt int) time.Duration {
	base := c.RetryBaseWait
	if base <= 0 {
		base = defaultRetryBaseWait
	}
	wait := time.Duration(float64(base) * math.Pow(2, float64(attempt-1)))
	if wait > maxRetryWait {
		wait = maxRetryWait
	}
	// Add jitter: 0 to 50% of wait
	if wait > 0 {
		jitter := time.Duration(rand.Int63n(int64(wait) / 2))
		wait += jitter
	}
	return wait
}

func errorCodeFromStatus(status int) string {
	switch status {
	case 401:
		return "UNAUTHORIZED"
	case 403:
		return "FORBIDDEN"
	case 404:
		return "NOT_FOUND"
	case 429:
		return "RATE_LIMITED"
	default:
		return "API_ERROR"
	}
}

// Helper pointer functions for building optional request fields.
func StringPtr(s string) *string    { return &s }
func IntPtr(i int) *int             { return &i }
func Int64Ptr(i int64) *int64       { return &i }
func BoolPtr(b bool) *bool          { return &b }
func Float64Ptr(f float64) *float64 { return &f }
