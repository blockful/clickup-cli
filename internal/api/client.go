package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const DefaultBaseURL = "https://api.clickup.com/api"

// ClientInterface defines all ClickUp API operations.
type ClientInterface interface {
	// Auth
	GetUser() (*UserResponse, error)

	// Workspaces
	ListWorkspaces() (*WorkspacesResponse, error)

	// Spaces
	ListSpaces(workspaceID string) (*SpacesResponse, error)
	GetSpace(spaceID string) (*Space, error)
	CreateSpace(workspaceID string, req *CreateSpaceRequest) (*Space, error)
	UpdateSpace(spaceID string, req *UpdateSpaceRequest) (*Space, error)
	DeleteSpace(spaceID string) error

	// Folders
	ListFolders(spaceID string) (*FoldersResponse, error)
	GetFolder(folderID string) (*Folder, error)
	CreateFolder(spaceID string, req *CreateFolderRequest) (*Folder, error)
	UpdateFolder(folderID string, req *UpdateFolderRequest) (*Folder, error)
	DeleteFolder(folderID string) error

	// Lists
	ListLists(folderID string) (*ListsResponse, error)
	ListFolderlessLists(spaceID string) (*ListsResponse, error)
	GetList(listID string) (*List, error)
	CreateList(folderID string, req *CreateListRequest) (*List, error)
	CreateFolderlessList(spaceID string, req *CreateListRequest) (*List, error)
	UpdateList(listID string, req *UpdateListRequest) (*List, error)
	DeleteList(listID string) error

	// Tasks
	ListTasks(listID string, opts *ListTasksOptions) (*TasksResponse, error)
	GetTask(taskID string, opts ...GetTaskOptions) (*Task, error)
	CreateTask(listID string, req *CreateTaskRequest) (*Task, error)
	UpdateTask(taskID string, req *UpdateTaskRequest, opts ...UpdateTaskOptions) (*Task, error)
	DeleteTask(taskID string) error
	SearchTasks(teamID string, opts *SearchTasksOptions) (*TasksResponse, error)

	// Comments
	ListComments(taskID string) (*CommentsResponse, error)
	ListListComments(listID string) (*CommentsResponse, error)
	CreateComment(taskID string, req *CreateCommentRequest) (*CreateCommentResponse, error)
	CreateListComment(listID string, req *CreateCommentRequest) (*CreateCommentResponse, error)
	UpdateComment(commentID string, req *UpdateCommentRequest) error
	DeleteComment(commentID string) error

	// Custom Fields
	GetListCustomFields(listID string) (*CustomFieldsResponse, error)
	GetFolderCustomFields(folderID string) (*CustomFieldsResponse, error)
	GetSpaceCustomFields(spaceID string) (*CustomFieldsResponse, error)
	GetWorkspaceCustomFields(teamID string) (*CustomFieldsResponse, error)
	SetCustomFieldValue(taskID, fieldID string, req *SetCustomFieldRequest) error
	RemoveCustomFieldValue(taskID, fieldID string) error

	// Tags
	GetSpaceTags(spaceID string) (*TagsResponse, error)
	CreateSpaceTag(spaceID string, req *CreateTagRequest) error
	UpdateSpaceTag(spaceID, tagName string, req *UpdateTagRequest) error
	DeleteSpaceTag(spaceID, tagName string) error
	AddTagToTask(taskID, tagName string) error
	RemoveTagFromTask(taskID, tagName string) error

	// Checklists
	CreateChecklist(taskID string, req *CreateChecklistRequest) (*ChecklistResponse, error)
	EditChecklist(checklistID string, req *EditChecklistRequest) error
	DeleteChecklist(checklistID string) error
	CreateChecklistItem(checklistID string, req *CreateChecklistItemRequest) (*ChecklistResponse, error)
	EditChecklistItem(checklistID, checklistItemID string, req *EditChecklistItemRequest) (*ChecklistResponse, error)
	DeleteChecklistItem(checklistID, checklistItemID string) error

	// Docs (v3)
	CreateDoc(workspaceID string, req *CreateDocRequest) (*Doc, error)
	SearchDocs(workspaceID string) (*DocsResponse, error)
	GetDoc(workspaceID, docID string) (*Doc, error)
	CreatePage(workspaceID, docID string, req *CreatePageRequest) (*DocPage, error)
	GetPage(workspaceID, docID, pageID string) (*DocPage, error)
	EditPage(workspaceID, docID, pageID string, req *EditPageRequest) (*DocPage, error)
	GetDocPageListing(workspaceID, docID string) (*DocPagesResponse, error)

	// Time Tracking
	GetTimeEntries(teamID string, opts *ListTimeEntriesOptions) (*TimeEntriesResponse, error)
	CreateTimeEntry(teamID string, req *CreateTimeEntryRequest) (*TimeEntry, error)
	GetTimeEntry(teamID, timerID string) (*SingleTimeEntryResponse, error)
	UpdateTimeEntry(teamID, timerID string, req *UpdateTimeEntryRequest) error
	DeleteTimeEntry(teamID, timerID string) error
	StartTimer(teamID string, req *StartTimerRequest) (*SingleTimeEntryResponse, error)
	StopTimer(teamID string) (*SingleTimeEntryResponse, error)
	GetRunningTimer(teamID string, assignee string) (*SingleTimeEntryResponse, error)
	GetTimeEntryTags(teamID string) (*TimeEntryTagsResponse, error)

	// Webhooks
	GetWebhooks(teamID string) (*WebhooksResponse, error)
	CreateWebhook(teamID string, req *CreateWebhookRequest) (*CreateWebhookResponse, error)
	UpdateWebhook(webhookID string, req *UpdateWebhookRequest) (*UpdateWebhookResponse, error)
	DeleteWebhook(webhookID string) error

	// Views
	GetTeamViews(teamID string) (*ViewsResponse, error)
	GetSpaceViews(spaceID string) (*ViewsResponse, error)
	GetFolderViews(folderID string) (*ViewsResponse, error)
	GetListViews(listID string) (*ViewsResponse, error)
	GetView(viewID string) (*ViewResponse, error)
	CreateTeamView(teamID string, req *CreateViewRequest) (*ViewResponse, error)
	CreateSpaceView(spaceID string, req *CreateViewRequest) (*ViewResponse, error)
	CreateFolderView(folderID string, req *CreateViewRequest) (*ViewResponse, error)
	CreateListView(listID string, req *CreateViewRequest) (*ViewResponse, error)
	UpdateView(viewID string, req *UpdateViewRequest) (*ViewResponse, error)
	DeleteView(viewID string) error
	GetViewTasks(viewID string, page int) (*ViewTasksResponse, error)

	// Goals
	GetGoals(teamID string, includeCompleted bool) (*GoalsResponse, error)
	GetGoal(goalID string) (*GoalResponse, error)
	CreateGoal(teamID string, req *CreateGoalRequest) (*GoalResponse, error)
	UpdateGoal(goalID string, req *UpdateGoalRequest) (*GoalResponse, error)
	DeleteGoal(goalID string) error
	CreateKeyResult(goalID string, req *CreateKeyResultRequest) (*KeyResultResponse, error)

	// Members
	GetListMembers(listID string) (*MembersResponse, error)
	GetTaskMembers(taskID string) (*MembersResponse, error)

	// Groups
	GetGroups(teamID string) (*GroupsResponse, error)
	CreateGroup(teamID string, req *CreateGroupRequest) (*Group, error)
	UpdateGroup(groupID string, req *UpdateGroupRequest) (*Group, error)
	DeleteGroup(groupID string) error

	// Guests
	InviteGuest(teamID string, req *InviteGuestRequest) error
	GetGuest(teamID, guestID string) (*GuestResponse, error)
	EditGuest(teamID, guestID string, req *EditGuestRequest) (*GuestResponse, error)
	RemoveGuest(teamID, guestID string) error
}

// Client implements ClientInterface using HTTP requests to the ClickUp API.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

var _ ClientInterface = (*Client)(nil)

type APIError struct {
	Err     string `json:"err"`
	ECODE   string `json:"ECODE"`
	Message string `json:"message,omitempty"`
}

type ClientError struct {
	StatusCode int
	Code       string
	Message    string
}

func (e *ClientError) Error() string {
	return e.Message
}

func NewClient(token string) *Client {
	return &Client{
		BaseURL: DefaultBaseURL,
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Do(method, path string, body interface{}, result interface{}) error {
	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return &ClientError{Code: "MARSHAL_ERROR", Message: fmt.Sprintf("failed to marshal request body: %v", err)}
		}
		reqBody = bytes.NewReader(data)
	}

	url := c.BaseURL + path
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return &ClientError{Code: "REQUEST_ERROR", Message: fmt.Sprintf("failed to create request: %v", err)}
	}

	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return &ClientError{Code: "NETWORK_ERROR", Message: fmt.Sprintf("request failed: %v", err)}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ClientError{Code: "READ_ERROR", Message: fmt.Sprintf("failed to read response: %v", err)}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		code := errorCodeFromStatus(resp.StatusCode)
		var apiErr APIError
		if json.Unmarshal(respBody, &apiErr) == nil && (apiErr.Err != "" || apiErr.Message != "") {
			msg := apiErr.Err
			if msg == "" {
				msg = apiErr.Message
			}
			return &ClientError{StatusCode: resp.StatusCode, Code: code, Message: msg}
		}
		return &ClientError{StatusCode: resp.StatusCode, Code: code, Message: fmt.Sprintf("API returned status %d", resp.StatusCode)}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return &ClientError{Code: "UNMARSHAL_ERROR", Message: fmt.Sprintf("failed to parse response: %v", err)}
		}
	}

	return nil
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
