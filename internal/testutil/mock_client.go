package testutil

import (
	"context"

	"github.com/blockful/clickup-cli/internal/api"
)

// MockClient implements api.ClientInterface for testing.
type MockClient struct {
	GetUserFn              func(context.Context) (*api.UserResponse, error)
	ListWorkspacesFn       func(context.Context) (*api.WorkspacesResponse, error)
	ListSpacesFn           func(context.Context, string) (*api.SpacesResponse, error)
	GetSpaceFn             func(context.Context, string) (*api.Space, error)
	CreateSpaceFn          func(context.Context, string, *api.CreateSpaceRequest) (*api.Space, error)
	UpdateSpaceFn          func(context.Context, string, *api.UpdateSpaceRequest) (*api.Space, error)
	DeleteSpaceFn          func(context.Context, string) error
	ListFoldersFn          func(context.Context, string) (*api.FoldersResponse, error)
	GetFolderFn            func(context.Context, string) (*api.Folder, error)
	CreateFolderFn         func(context.Context, string, *api.CreateFolderRequest) (*api.Folder, error)
	UpdateFolderFn         func(context.Context, string, *api.UpdateFolderRequest) (*api.Folder, error)
	DeleteFolderFn         func(context.Context, string) error
	ListListsFn            func(context.Context, string) (*api.ListsResponse, error)
	ListFolderlessListsFn  func(context.Context, string) (*api.ListsResponse, error)
	GetListFn              func(context.Context, string) (*api.List, error)
	CreateListFn           func(context.Context, string, *api.CreateListRequest) (*api.List, error)
	CreateFolderlessListFn func(context.Context, string, *api.CreateListRequest) (*api.List, error)
	UpdateListFn           func(context.Context, string, *api.UpdateListRequest) (*api.List, error)
	DeleteListFn           func(context.Context, string) error
	ListTasksFn            func(context.Context, string, *api.ListTasksOptions) (*api.TasksResponse, error)
	GetTaskFn              func(context.Context, string, ...api.GetTaskOptions) (*api.Task, error)
	CreateTaskFn           func(context.Context, string, *api.CreateTaskRequest) (*api.Task, error)
	UpdateTaskFn           func(context.Context, string, *api.UpdateTaskRequest, ...api.UpdateTaskOptions) (*api.Task, error)
	DeleteTaskFn           func(context.Context, string) error
	SearchTasksFn          func(context.Context, string, *api.SearchTasksOptions) (*api.TasksResponse, error)
	ListCommentsFn         func(context.Context, string, string) (*api.CommentsResponse, error)
	ListListCommentsFn         func(context.Context, string, string) (*api.CommentsResponse, error)
	CreateCommentFn        func(context.Context, string, *api.CreateCommentRequest) (*api.CreateCommentResponse, error)
	CreateListCommentFn    func(context.Context, string, *api.CreateCommentRequest) (*api.CreateCommentResponse, error)
	UpdateCommentFn        func(context.Context, string, *api.UpdateCommentRequest) error
	DeleteCommentFn        func(context.Context, string) error

	// Custom Fields
	GetListCustomFieldsFn      func(context.Context, string) (*api.CustomFieldsResponse, error)
	GetFolderCustomFieldsFn    func(context.Context, string) (*api.CustomFieldsResponse, error)
	GetSpaceCustomFieldsFn     func(context.Context, string) (*api.CustomFieldsResponse, error)
	GetWorkspaceCustomFieldsFn func(context.Context, string) (*api.CustomFieldsResponse, error)
	SetCustomFieldValueFn      func(context.Context, string, string, *api.SetCustomFieldRequest) error
	RemoveCustomFieldValueFn   func(context.Context, string, string) error

	// Tags
	GetSpaceTagsFn      func(context.Context, string) (*api.TagsResponse, error)
	CreateSpaceTagFn    func(context.Context, string, *api.CreateTagRequest) error
	UpdateSpaceTagFn    func(context.Context, string, string, *api.UpdateTagRequest) error
	DeleteSpaceTagFn    func(context.Context, string, string) error
	AddTagToTaskFn      func(context.Context, string, string) error
	RemoveTagFromTaskFn func(context.Context, string, string) error

	// Checklists
	CreateChecklistFn     func(context.Context, string, *api.CreateChecklistRequest) (*api.ChecklistResponse, error)
	EditChecklistFn       func(context.Context, string, *api.EditChecklistRequest) error
	DeleteChecklistFn     func(context.Context, string) error
	CreateChecklistItemFn func(context.Context, string, *api.CreateChecklistItemRequest) (*api.ChecklistResponse, error)
	EditChecklistItemFn   func(context.Context, string, string, *api.EditChecklistItemRequest) (*api.ChecklistResponse, error)
	DeleteChecklistItemFn func(context.Context, string, string) error

	// Docs
	CreateDocFn         func(context.Context, string, *api.CreateDocRequest) (*api.Doc, error)
	SearchDocsFn        func(context.Context, string) (*api.DocsResponse, error)
	GetDocFn            func(context.Context, string, string) (*api.Doc, error)
	CreatePageFn        func(context.Context, string, string, *api.CreatePageRequest) (*api.DocPage, error)
	GetPageFn           func(context.Context, string, string, string) (*api.DocPage, error)
	EditPageFn          func(context.Context, string, string, string, *api.EditPageRequest) (*api.DocPage, error)
	GetDocPageListingFn func(context.Context, string, string) (*api.DocPagesResponse, error)

	// Time Tracking
	GetTimeEntriesFn   func(context.Context, string, *api.ListTimeEntriesOptions) (*api.TimeEntriesResponse, error)
	CreateTimeEntryFn  func(context.Context, string, *api.CreateTimeEntryRequest) (*api.TimeEntry, error)
	GetTimeEntryFn     func(context.Context, string, string) (*api.SingleTimeEntryResponse, error)
	UpdateTimeEntryFn  func(context.Context, string, string, *api.UpdateTimeEntryRequest) error
	DeleteTimeEntryFn  func(context.Context, string, string) error
	StartTimerFn       func(context.Context, string, *api.StartTimerRequest) (*api.SingleTimeEntryResponse, error)
	StopTimerFn        func(context.Context, string) (*api.SingleTimeEntryResponse, error)
	GetRunningTimerFn  func(context.Context, string, string) (*api.SingleTimeEntryResponse, error)
	GetTimeEntryTagsFn            func(context.Context, string) (*api.TimeEntryTagsResponse, error)
	GetTimeEntryHistoryFn         func(context.Context, string, string) (*api.TimeEntryHistoryResponse, error)
	AddTagsToTimeEntriesFn        func(context.Context, string, *api.AddTagsToTimeEntriesRequest) error
	RemoveTagsFromTimeEntriesFn   func(context.Context, string, *api.RemoveTagsFromTimeEntriesRequest) error
	ChangeTagNamesFn              func(context.Context, string, *api.ChangeTagNameRequest) error
	GetLegacyTrackedTimeFn        func(context.Context, string, string) (*api.LegacyTimeResponse, error)
	TrackLegacyTimeFn             func(context.Context, string, *api.LegacyTrackTimeRequest) (*api.LegacyTimeResponse, error)
	EditLegacyTimeFn              func(context.Context, string, string, *api.LegacyEditTimeRequest) error
	DeleteLegacyTimeFn            func(context.Context, string, string) error

	// Webhooks
	GetWebhooksFn   func(context.Context, string) (*api.WebhooksResponse, error)
	CreateWebhookFn func(context.Context, string, *api.CreateWebhookRequest) (*api.CreateWebhookResponse, error)
	UpdateWebhookFn func(context.Context, string, *api.UpdateWebhookRequest) (*api.UpdateWebhookResponse, error)
	DeleteWebhookFn func(context.Context, string) error

	// Views
	GetTeamViewsFn     func(context.Context, string) (*api.ViewsResponse, error)
	GetSpaceViewsFn    func(context.Context, string) (*api.ViewsResponse, error)
	GetFolderViewsFn   func(context.Context, string) (*api.ViewsResponse, error)
	GetListViewsFn     func(context.Context, string) (*api.ViewsResponse, error)
	GetViewFn          func(context.Context, string) (*api.ViewResponse, error)
	CreateTeamViewFn   func(context.Context, string, *api.CreateViewRequest) (*api.ViewResponse, error)
	CreateSpaceViewFn  func(context.Context, string, *api.CreateViewRequest) (*api.ViewResponse, error)
	CreateFolderViewFn func(context.Context, string, *api.CreateViewRequest) (*api.ViewResponse, error)
	CreateListViewFn   func(context.Context, string, *api.CreateViewRequest) (*api.ViewResponse, error)
	UpdateViewFn       func(context.Context, string, *api.UpdateViewRequest) (*api.ViewResponse, error)
	DeleteViewFn       func(context.Context, string) error
	GetViewTasksFn     func(context.Context, string, int) (*api.ViewTasksResponse, error)

	// Goals
	GetGoalsFn        func(context.Context, string, bool) (*api.GoalsResponse, error)
	GetGoalFn         func(context.Context, string) (*api.GoalResponse, error)
	CreateGoalFn      func(context.Context, string, *api.CreateGoalRequest) (*api.GoalResponse, error)
	UpdateGoalFn      func(context.Context, string, *api.UpdateGoalRequest) (*api.GoalResponse, error)
	DeleteGoalFn      func(context.Context, string) error
	CreateKeyResultFn func(context.Context, string, *api.CreateKeyResultRequest) (*api.KeyResultResponse, error)
	UpdateKeyResultFn func(context.Context, string, *api.UpdateKeyResultRequest) (*api.KeyResultResponse, error)
	DeleteKeyResultFn func(context.Context, string) error

	// Members
	GetListMembersFn func(context.Context, string) (*api.MembersResponse, error)
	GetTaskMembersFn func(context.Context, string) (*api.MembersResponse, error)

	// Groups
	GetGroupsFn   func(context.Context, string, []string) (*api.GroupsResponse, error)
	CreateGroupFn func(context.Context, string, *api.CreateGroupRequest) (*api.Group, error)
	UpdateGroupFn func(context.Context, string, *api.UpdateGroupRequest) (*api.Group, error)
	DeleteGroupFn func(context.Context, string) error

	// Guests
	InviteGuestFn func(context.Context, string, *api.InviteGuestRequest) error
	GetGuestFn    func(context.Context, string, string) (*api.GuestResponse, error)
	EditGuestFn   func(context.Context, string, string, *api.EditGuestRequest) (*api.GuestResponse, error)
	RemoveGuestFn func(context.Context, string, string) error

	// Threaded Comments
	ListThreadedCommentsFn  func(context.Context, string) (*api.CommentsResponse, error)
	CreateThreadedCommentFn func(context.Context, string, *api.CreateCommentRequest) (*api.CreateCommentResponse, error)

	// View Comments
	ListViewCommentsFn  func(context.Context, string, string) (*api.CommentsResponse, error)
	CreateViewCommentFn func(context.Context, string, *api.CreateCommentRequest) (*api.CreateCommentResponse, error)

	// Guest Assignments
	AddGuestToTaskFn       func(context.Context, string, int, *api.GuestPermissionRequest, bool) (*api.GuestResponse, error)
	RemoveGuestFromTaskFn  func(context.Context, string, int, bool) error
	AddGuestToListFn       func(context.Context, string, int, *api.GuestPermissionRequest, bool) (*api.GuestResponse, error)
	RemoveGuestFromListFn  func(context.Context, string, int, bool) error
	AddGuestToFolderFn     func(context.Context, string, int, *api.GuestPermissionRequest, bool) (*api.GuestResponse, error)
	RemoveGuestFromFolderFn func(context.Context, string, int, bool) error

	// Users
	InviteUserFn  func(context.Context, string, *api.InviteUserRequest) (*api.TeamUserResponse, error)
	GetTeamUserFn func(context.Context, string, string, bool) (*api.TeamUserResponse, error)
	EditUserFn    func(context.Context, string, string, *api.EditUserRequest) (*api.TeamUserResponse, error)
	RemoveUserFn  func(context.Context, string, string) error

	// Roles
	GetCustomRolesFn func(context.Context, string, bool) (*api.CustomRolesResponse, error)

	// Custom Task Types
	GetCustomTaskTypesFn func(context.Context, string) (*api.CustomTaskTypesResponse, error)

	// Shared Hierarchy
	GetSharedHierarchyFn func(context.Context, string) (*api.SharedHierarchyResponse, error)

	// Workspace extras
	GetWorkspaceSeatsFn func(context.Context, string) (*api.SeatsResponse, error)
	GetWorkspacePlanFn  func(context.Context, string) (*api.PlanResponse, error)

	// Templates
	GetTaskTemplatesFn            func(context.Context, string, int) (*api.TaskTemplatesResponse, error)
	CreateTaskFromTemplateFn      func(context.Context, string, string, *api.CreateFromTemplateRequest) (*api.CreateFromTemplateResponse, error)
	CreateFolderFromTemplateFn    func(context.Context, string, string, *api.CreateFromTemplateRequest) (*api.CreateFromTemplateResponse, error)
	CreateListFromFolderTemplateFn func(context.Context, string, string, *api.CreateFromTemplateRequest) (*api.CreateFromTemplateResponse, error)
	CreateListFromSpaceTemplateFn  func(context.Context, string, string, *api.CreateFromTemplateRequest) (*api.CreateFromTemplateResponse, error)

	// Relationships
	AddDependencyFn       func(context.Context, string, *api.AddDependencyRequest) (*api.DependencyResponse, error)
	DeleteDependencyFn    func(context.Context, string, string, string) error
	AddTaskLinkFn         func(context.Context, string, string) (*api.TaskLinkResponse, error)
	DeleteTaskLinkFn      func(context.Context, string, string) error

	// Task Extras
	MergeTasksFn          func(context.Context, string, *api.MergeTasksRequest) error
	GetTimeInStatusFn     func(context.Context, string) (*api.TimeInStatusResponse, error)
	GetBulkTimeInStatusFn func(context.Context, []string) (*api.BulkTimeInStatusResponse, error)
	AddTaskToListFn       func(context.Context, string, string) error
	RemoveTaskFromListFn  func(context.Context, string, string) error

	// Attachments
	CreateTaskAttachmentFn func(context.Context, string, string) (*api.Attachment, error)
}

var _ api.ClientInterface = (*MockClient)(nil)

func (m *MockClient) GetUser(ctx context.Context) (*api.UserResponse, error) {
	return m.GetUserFn(ctx)
}
func (m *MockClient) ListWorkspaces(ctx context.Context) (*api.WorkspacesResponse, error) {
	return m.ListWorkspacesFn(ctx)
}
func (m *MockClient) ListSpaces(ctx context.Context, id string) (*api.SpacesResponse, error) {
	return m.ListSpacesFn(ctx, id)
}
func (m *MockClient) GetSpace(ctx context.Context, id string) (*api.Space, error) {
	return m.GetSpaceFn(ctx, id)
}
func (m *MockClient) CreateSpace(ctx context.Context, wid string, req *api.CreateSpaceRequest) (*api.Space, error) {
	return m.CreateSpaceFn(ctx, wid, req)
}
func (m *MockClient) UpdateSpace(ctx context.Context, id string, req *api.UpdateSpaceRequest) (*api.Space, error) {
	return m.UpdateSpaceFn(ctx, id, req)
}
func (m *MockClient) DeleteSpace(ctx context.Context, id string) error {
	return m.DeleteSpaceFn(ctx, id)
}
func (m *MockClient) ListFolders(ctx context.Context, id string) (*api.FoldersResponse, error) {
	return m.ListFoldersFn(ctx, id)
}
func (m *MockClient) GetFolder(ctx context.Context, id string) (*api.Folder, error) {
	return m.GetFolderFn(ctx, id)
}
func (m *MockClient) CreateFolder(ctx context.Context, sid string, req *api.CreateFolderRequest) (*api.Folder, error) {
	return m.CreateFolderFn(ctx, sid, req)
}
func (m *MockClient) UpdateFolder(ctx context.Context, id string, req *api.UpdateFolderRequest) (*api.Folder, error) {
	return m.UpdateFolderFn(ctx, id, req)
}
func (m *MockClient) DeleteFolder(ctx context.Context, id string) error {
	return m.DeleteFolderFn(ctx, id)
}
func (m *MockClient) ListLists(ctx context.Context, id string) (*api.ListsResponse, error) {
	return m.ListListsFn(ctx, id)
}
func (m *MockClient) ListFolderlessLists(ctx context.Context, id string) (*api.ListsResponse, error) {
	return m.ListFolderlessListsFn(ctx, id)
}
func (m *MockClient) GetList(ctx context.Context, id string) (*api.List, error) {
	return m.GetListFn(ctx, id)
}
func (m *MockClient) CreateList(ctx context.Context, fid string, req *api.CreateListRequest) (*api.List, error) {
	return m.CreateListFn(ctx, fid, req)
}
func (m *MockClient) CreateFolderlessList(ctx context.Context, sid string, req *api.CreateListRequest) (*api.List, error) {
	return m.CreateFolderlessListFn(ctx, sid, req)
}
func (m *MockClient) UpdateList(ctx context.Context, id string, req *api.UpdateListRequest) (*api.List, error) {
	return m.UpdateListFn(ctx, id, req)
}
func (m *MockClient) DeleteList(ctx context.Context, id string) error {
	return m.DeleteListFn(ctx, id)
}
func (m *MockClient) ListTasks(ctx context.Context, id string, opts *api.ListTasksOptions) (*api.TasksResponse, error) {
	return m.ListTasksFn(ctx, id, opts)
}
func (m *MockClient) GetTask(ctx context.Context, id string, opts ...api.GetTaskOptions) (*api.Task, error) {
	return m.GetTaskFn(ctx, id, opts...)
}
func (m *MockClient) CreateTask(ctx context.Context, lid string, req *api.CreateTaskRequest) (*api.Task, error) {
	return m.CreateTaskFn(ctx, lid, req)
}
func (m *MockClient) UpdateTask(ctx context.Context, id string, req *api.UpdateTaskRequest, opts ...api.UpdateTaskOptions) (*api.Task, error) {
	return m.UpdateTaskFn(ctx, id, req, opts...)
}
func (m *MockClient) DeleteTask(ctx context.Context, id string) error {
	return m.DeleteTaskFn(ctx, id)
}
func (m *MockClient) SearchTasks(ctx context.Context, teamID string, opts *api.SearchTasksOptions) (*api.TasksResponse, error) {
	return m.SearchTasksFn(ctx, teamID, opts)
}
func (m *MockClient) ListComments(ctx context.Context, id string, startID string, opts ...*api.TaskScopedOptions) (*api.CommentsResponse, error) {
	return m.ListCommentsFn(ctx, id, startID)
}
func (m *MockClient) ListListComments(ctx context.Context, id string, startID string) (*api.CommentsResponse, error) {
	return m.ListListCommentsFn(ctx, id, startID)
}
func (m *MockClient) CreateComment(ctx context.Context, tid string, req *api.CreateCommentRequest, opts ...*api.TaskScopedOptions) (*api.CreateCommentResponse, error) {
	return m.CreateCommentFn(ctx, tid, req)
}
func (m *MockClient) CreateListComment(ctx context.Context, lid string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error) {
	return m.CreateListCommentFn(ctx, lid, req)
}
func (m *MockClient) UpdateComment(ctx context.Context, id string, req *api.UpdateCommentRequest) error {
	return m.UpdateCommentFn(ctx, id, req)
}
func (m *MockClient) DeleteComment(ctx context.Context, id string) error {
	return m.DeleteCommentFn(ctx, id)
}

// Custom Fields
func (m *MockClient) GetListCustomFields(ctx context.Context, id string) (*api.CustomFieldsResponse, error) {
	return m.GetListCustomFieldsFn(ctx, id)
}
func (m *MockClient) GetFolderCustomFields(ctx context.Context, id string) (*api.CustomFieldsResponse, error) {
	return m.GetFolderCustomFieldsFn(ctx, id)
}
func (m *MockClient) GetSpaceCustomFields(ctx context.Context, id string) (*api.CustomFieldsResponse, error) {
	return m.GetSpaceCustomFieldsFn(ctx, id)
}
func (m *MockClient) GetWorkspaceCustomFields(ctx context.Context, id string) (*api.CustomFieldsResponse, error) {
	return m.GetWorkspaceCustomFieldsFn(ctx, id)
}
func (m *MockClient) SetCustomFieldValue(ctx context.Context, taskID, fieldID string, req *api.SetCustomFieldRequest, opts ...*api.TaskScopedOptions) error {
	return m.SetCustomFieldValueFn(ctx, taskID, fieldID, req)
}
func (m *MockClient) RemoveCustomFieldValue(ctx context.Context, taskID, fieldID string, opts ...*api.TaskScopedOptions) error {
	return m.RemoveCustomFieldValueFn(ctx, taskID, fieldID)
}

// Tags
func (m *MockClient) GetSpaceTags(ctx context.Context, id string) (*api.TagsResponse, error) {
	return m.GetSpaceTagsFn(ctx, id)
}
func (m *MockClient) CreateSpaceTag(ctx context.Context, id string, req *api.CreateTagRequest) error {
	return m.CreateSpaceTagFn(ctx, id, req)
}
func (m *MockClient) UpdateSpaceTag(ctx context.Context, id, name string, req *api.UpdateTagRequest) error {
	return m.UpdateSpaceTagFn(ctx, id, name, req)
}
func (m *MockClient) DeleteSpaceTag(ctx context.Context, id, name string) error {
	return m.DeleteSpaceTagFn(ctx, id, name)
}
func (m *MockClient) AddTagToTask(ctx context.Context, taskID, tagName string, opts ...*api.TaskScopedOptions) error {
	return m.AddTagToTaskFn(ctx, taskID, tagName)
}
func (m *MockClient) RemoveTagFromTask(ctx context.Context, taskID, tagName string, opts ...*api.TaskScopedOptions) error {
	return m.RemoveTagFromTaskFn(ctx, taskID, tagName)
}

// Checklists
func (m *MockClient) CreateChecklist(ctx context.Context, taskID string, req *api.CreateChecklistRequest, opts ...*api.TaskScopedOptions) (*api.ChecklistResponse, error) {
	return m.CreateChecklistFn(ctx, taskID, req)
}
func (m *MockClient) EditChecklist(ctx context.Context, id string, req *api.EditChecklistRequest) error {
	return m.EditChecklistFn(ctx, id, req)
}
func (m *MockClient) DeleteChecklist(ctx context.Context, id string) error {
	return m.DeleteChecklistFn(ctx, id)
}
func (m *MockClient) CreateChecklistItem(ctx context.Context, id string, req *api.CreateChecklistItemRequest) (*api.ChecklistResponse, error) {
	return m.CreateChecklistItemFn(ctx, id, req)
}
func (m *MockClient) EditChecklistItem(ctx context.Context, cid, iid string, req *api.EditChecklistItemRequest) (*api.ChecklistResponse, error) {
	return m.EditChecklistItemFn(ctx, cid, iid, req)
}
func (m *MockClient) DeleteChecklistItem(ctx context.Context, cid, iid string) error {
	return m.DeleteChecklistItemFn(ctx, cid, iid)
}

// Docs
func (m *MockClient) CreateDoc(ctx context.Context, wid string, req *api.CreateDocRequest) (*api.Doc, error) {
	return m.CreateDocFn(ctx, wid, req)
}
func (m *MockClient) SearchDocs(ctx context.Context, wid string) (*api.DocsResponse, error) {
	return m.SearchDocsFn(ctx, wid)
}
func (m *MockClient) GetDoc(ctx context.Context, wid, did string) (*api.Doc, error) {
	return m.GetDocFn(ctx, wid, did)
}
func (m *MockClient) CreatePage(ctx context.Context, wid, did string, req *api.CreatePageRequest) (*api.DocPage, error) {
	return m.CreatePageFn(ctx, wid, did, req)
}
func (m *MockClient) GetPage(ctx context.Context, wid, did, pid string) (*api.DocPage, error) {
	return m.GetPageFn(ctx, wid, did, pid)
}
func (m *MockClient) EditPage(ctx context.Context, wid, did, pid string, req *api.EditPageRequest) (*api.DocPage, error) {
	return m.EditPageFn(ctx, wid, did, pid, req)
}
func (m *MockClient) GetDocPageListing(ctx context.Context, wid, did string) (*api.DocPagesResponse, error) {
	return m.GetDocPageListingFn(ctx, wid, did)
}

// Time Tracking
func (m *MockClient) GetTimeEntries(ctx context.Context, teamID string, opts *api.ListTimeEntriesOptions) (*api.TimeEntriesResponse, error) {
	return m.GetTimeEntriesFn(ctx, teamID, opts)
}
func (m *MockClient) CreateTimeEntry(ctx context.Context, teamID string, req *api.CreateTimeEntryRequest) (*api.TimeEntry, error) {
	return m.CreateTimeEntryFn(ctx, teamID, req)
}
func (m *MockClient) GetTimeEntry(ctx context.Context, teamID, timerID string, opts *api.GetTimeEntryOptions) (*api.SingleTimeEntryResponse, error) {
	return m.GetTimeEntryFn(ctx, teamID, timerID)
}
func (m *MockClient) UpdateTimeEntry(ctx context.Context, teamID, timerID string, req *api.UpdateTimeEntryRequest) error {
	return m.UpdateTimeEntryFn(ctx, teamID, timerID, req)
}
func (m *MockClient) DeleteTimeEntry(ctx context.Context, teamID, timerID string) error {
	return m.DeleteTimeEntryFn(ctx, teamID, timerID)
}
func (m *MockClient) StartTimer(ctx context.Context, teamID string, req *api.StartTimerRequest) (*api.SingleTimeEntryResponse, error) {
	return m.StartTimerFn(ctx, teamID, req)
}
func (m *MockClient) StopTimer(ctx context.Context, teamID string) (*api.SingleTimeEntryResponse, error) {
	return m.StopTimerFn(ctx, teamID)
}
func (m *MockClient) GetRunningTimer(ctx context.Context, teamID, assignee string) (*api.SingleTimeEntryResponse, error) {
	return m.GetRunningTimerFn(ctx, teamID, assignee)
}
func (m *MockClient) GetTimeEntryTags(ctx context.Context, teamID string) (*api.TimeEntryTagsResponse, error) {
	return m.GetTimeEntryTagsFn(ctx, teamID)
}
func (m *MockClient) GetTimeEntryHistory(ctx context.Context, teamID, timerID string) (*api.TimeEntryHistoryResponse, error) {
	return m.GetTimeEntryHistoryFn(ctx, teamID, timerID)
}
func (m *MockClient) AddTagsToTimeEntries(ctx context.Context, teamID string, req *api.AddTagsToTimeEntriesRequest) error {
	return m.AddTagsToTimeEntriesFn(ctx, teamID, req)
}
func (m *MockClient) RemoveTagsFromTimeEntries(ctx context.Context, teamID string, req *api.RemoveTagsFromTimeEntriesRequest) error {
	return m.RemoveTagsFromTimeEntriesFn(ctx, teamID, req)
}
func (m *MockClient) ChangeTagNames(ctx context.Context, teamID string, req *api.ChangeTagNameRequest) error {
	return m.ChangeTagNamesFn(ctx, teamID, req)
}
func (m *MockClient) GetLegacyTrackedTime(ctx context.Context, taskID string, subcategoryID string, opts ...*api.TaskScopedOptions) (*api.LegacyTimeResponse, error) {
	return m.GetLegacyTrackedTimeFn(ctx, taskID, subcategoryID)
}
func (m *MockClient) TrackLegacyTime(ctx context.Context, taskID string, req *api.LegacyTrackTimeRequest, opts ...*api.TaskScopedOptions) (*api.LegacyTimeResponse, error) {
	return m.TrackLegacyTimeFn(ctx, taskID, req)
}
func (m *MockClient) EditLegacyTime(ctx context.Context, taskID, intervalID string, req *api.LegacyEditTimeRequest, opts ...*api.TaskScopedOptions) error {
	return m.EditLegacyTimeFn(ctx, taskID, intervalID, req)
}
func (m *MockClient) DeleteLegacyTime(ctx context.Context, taskID, intervalID string, opts ...*api.TaskScopedOptions) error {
	return m.DeleteLegacyTimeFn(ctx, taskID, intervalID)
}

// Webhooks
func (m *MockClient) GetWebhooks(ctx context.Context, teamID string) (*api.WebhooksResponse, error) {
	return m.GetWebhooksFn(ctx, teamID)
}
func (m *MockClient) CreateWebhook(ctx context.Context, teamID string, req *api.CreateWebhookRequest) (*api.CreateWebhookResponse, error) {
	return m.CreateWebhookFn(ctx, teamID, req)
}
func (m *MockClient) UpdateWebhook(ctx context.Context, id string, req *api.UpdateWebhookRequest) (*api.UpdateWebhookResponse, error) {
	return m.UpdateWebhookFn(ctx, id, req)
}
func (m *MockClient) DeleteWebhook(ctx context.Context, id string) error {
	return m.DeleteWebhookFn(ctx, id)
}

// Views
func (m *MockClient) GetTeamViews(ctx context.Context, id string) (*api.ViewsResponse, error) {
	return m.GetTeamViewsFn(ctx, id)
}
func (m *MockClient) GetSpaceViews(ctx context.Context, id string) (*api.ViewsResponse, error) {
	return m.GetSpaceViewsFn(ctx, id)
}
func (m *MockClient) GetFolderViews(ctx context.Context, id string) (*api.ViewsResponse, error) {
	return m.GetFolderViewsFn(ctx, id)
}
func (m *MockClient) GetListViews(ctx context.Context, id string) (*api.ViewsResponse, error) {
	return m.GetListViewsFn(ctx, id)
}
func (m *MockClient) GetView(ctx context.Context, id string) (*api.ViewResponse, error) {
	return m.GetViewFn(ctx, id)
}
func (m *MockClient) CreateTeamView(ctx context.Context, id string, req *api.CreateViewRequest) (*api.ViewResponse, error) {
	return m.CreateTeamViewFn(ctx, id, req)
}
func (m *MockClient) CreateSpaceView(ctx context.Context, id string, req *api.CreateViewRequest) (*api.ViewResponse, error) {
	return m.CreateSpaceViewFn(ctx, id, req)
}
func (m *MockClient) CreateFolderView(ctx context.Context, id string, req *api.CreateViewRequest) (*api.ViewResponse, error) {
	return m.CreateFolderViewFn(ctx, id, req)
}
func (m *MockClient) CreateListView(ctx context.Context, id string, req *api.CreateViewRequest) (*api.ViewResponse, error) {
	return m.CreateListViewFn(ctx, id, req)
}
func (m *MockClient) UpdateView(ctx context.Context, id string, req *api.UpdateViewRequest) (*api.ViewResponse, error) {
	return m.UpdateViewFn(ctx, id, req)
}
func (m *MockClient) DeleteView(ctx context.Context, id string) error {
	return m.DeleteViewFn(ctx, id)
}
func (m *MockClient) GetViewTasks(ctx context.Context, id string, page int) (*api.ViewTasksResponse, error) {
	return m.GetViewTasksFn(ctx, id, page)
}

// Goals
func (m *MockClient) GetGoals(ctx context.Context, teamID string, includeCompleted bool) (*api.GoalsResponse, error) {
	return m.GetGoalsFn(ctx, teamID, includeCompleted)
}
func (m *MockClient) GetGoal(ctx context.Context, id string) (*api.GoalResponse, error) {
	return m.GetGoalFn(ctx, id)
}
func (m *MockClient) CreateGoal(ctx context.Context, teamID string, req *api.CreateGoalRequest) (*api.GoalResponse, error) {
	return m.CreateGoalFn(ctx, teamID, req)
}
func (m *MockClient) UpdateGoal(ctx context.Context, id string, req *api.UpdateGoalRequest) (*api.GoalResponse, error) {
	return m.UpdateGoalFn(ctx, id, req)
}
func (m *MockClient) DeleteGoal(ctx context.Context, id string) error {
	return m.DeleteGoalFn(ctx, id)
}
func (m *MockClient) CreateKeyResult(ctx context.Context, goalID string, req *api.CreateKeyResultRequest) (*api.KeyResultResponse, error) {
	return m.CreateKeyResultFn(ctx, goalID, req)
}
func (m *MockClient) UpdateKeyResult(ctx context.Context, keyResultID string, req *api.UpdateKeyResultRequest) (*api.KeyResultResponse, error) {
	return m.UpdateKeyResultFn(ctx, keyResultID, req)
}
func (m *MockClient) DeleteKeyResult(ctx context.Context, keyResultID string) error {
	return m.DeleteKeyResultFn(ctx, keyResultID)
}

// Members
func (m *MockClient) GetListMembers(ctx context.Context, id string) (*api.MembersResponse, error) {
	return m.GetListMembersFn(ctx, id)
}
func (m *MockClient) GetTaskMembers(ctx context.Context, id string) (*api.MembersResponse, error) {
	return m.GetTaskMembersFn(ctx, id)
}

// Groups
func (m *MockClient) GetGroups(ctx context.Context, teamID string, groupIDs []string) (*api.GroupsResponse, error) {
	return m.GetGroupsFn(ctx, teamID, groupIDs)
}
func (m *MockClient) CreateGroup(ctx context.Context, teamID string, req *api.CreateGroupRequest) (*api.Group, error) {
	return m.CreateGroupFn(ctx, teamID, req)
}
func (m *MockClient) UpdateGroup(ctx context.Context, id string, req *api.UpdateGroupRequest) (*api.Group, error) {
	return m.UpdateGroupFn(ctx, id, req)
}
func (m *MockClient) DeleteGroup(ctx context.Context, id string) error {
	return m.DeleteGroupFn(ctx, id)
}

// Guests
func (m *MockClient) InviteGuest(ctx context.Context, teamID string, req *api.InviteGuestRequest) error {
	return m.InviteGuestFn(ctx, teamID, req)
}
func (m *MockClient) GetGuest(ctx context.Context, teamID, guestID string) (*api.GuestResponse, error) {
	return m.GetGuestFn(ctx, teamID, guestID)
}
func (m *MockClient) EditGuest(ctx context.Context, teamID, guestID string, req *api.EditGuestRequest) (*api.GuestResponse, error) {
	return m.EditGuestFn(ctx, teamID, guestID, req)
}
func (m *MockClient) RemoveGuest(ctx context.Context, teamID, guestID string) error {
	return m.RemoveGuestFn(ctx, teamID, guestID)
}

// Threaded Comments
func (m *MockClient) ListThreadedComments(ctx context.Context, commentID string) (*api.CommentsResponse, error) {
	return m.ListThreadedCommentsFn(ctx, commentID)
}
func (m *MockClient) CreateThreadedComment(ctx context.Context, commentID string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error) {
	return m.CreateThreadedCommentFn(ctx, commentID, req)
}

// View Comments
func (m *MockClient) ListViewComments(ctx context.Context, viewID string, startID string) (*api.CommentsResponse, error) {
	return m.ListViewCommentsFn(ctx, viewID, startID)
}
func (m *MockClient) CreateViewComment(ctx context.Context, viewID string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error) {
	return m.CreateViewCommentFn(ctx, viewID, req)
}

// Guest Assignments
func (m *MockClient) AddGuestToTask(ctx context.Context, taskID string, guestID int, req *api.GuestPermissionRequest, includeShared bool, opts ...*api.TaskScopedOptions) (*api.GuestResponse, error) {
	return m.AddGuestToTaskFn(ctx, taskID, guestID, req, includeShared)
}
func (m *MockClient) RemoveGuestFromTask(ctx context.Context, taskID string, guestID int, includeShared bool, opts ...*api.TaskScopedOptions) error {
	return m.RemoveGuestFromTaskFn(ctx, taskID, guestID, includeShared)
}
func (m *MockClient) AddGuestToList(ctx context.Context, listID string, guestID int, req *api.GuestPermissionRequest, includeShared bool) (*api.GuestResponse, error) {
	return m.AddGuestToListFn(ctx, listID, guestID, req, includeShared)
}
func (m *MockClient) RemoveGuestFromList(ctx context.Context, listID string, guestID int, includeShared bool) error {
	return m.RemoveGuestFromListFn(ctx, listID, guestID, includeShared)
}
func (m *MockClient) AddGuestToFolder(ctx context.Context, folderID string, guestID int, req *api.GuestPermissionRequest, includeShared bool) (*api.GuestResponse, error) {
	return m.AddGuestToFolderFn(ctx, folderID, guestID, req, includeShared)
}
func (m *MockClient) RemoveGuestFromFolder(ctx context.Context, folderID string, guestID int, includeShared bool) error {
	return m.RemoveGuestFromFolderFn(ctx, folderID, guestID, includeShared)
}

// Templates
func (m *MockClient) GetTaskTemplates(ctx context.Context, teamID string, page int) (*api.TaskTemplatesResponse, error) {
	return m.GetTaskTemplatesFn(ctx, teamID, page)
}
func (m *MockClient) CreateTaskFromTemplate(ctx context.Context, listID, templateID string, req *api.CreateFromTemplateRequest) (*api.CreateFromTemplateResponse, error) {
	return m.CreateTaskFromTemplateFn(ctx, listID, templateID, req)
}
func (m *MockClient) CreateFolderFromTemplate(ctx context.Context, spaceID, templateID string, req *api.CreateFromTemplateRequest) (*api.CreateFromTemplateResponse, error) {
	return m.CreateFolderFromTemplateFn(ctx, spaceID, templateID, req)
}
func (m *MockClient) CreateListFromFolderTemplate(ctx context.Context, folderID, templateID string, req *api.CreateFromTemplateRequest) (*api.CreateFromTemplateResponse, error) {
	return m.CreateListFromFolderTemplateFn(ctx, folderID, templateID, req)
}
func (m *MockClient) CreateListFromSpaceTemplate(ctx context.Context, spaceID, templateID string, req *api.CreateFromTemplateRequest) (*api.CreateFromTemplateResponse, error) {
	return m.CreateListFromSpaceTemplateFn(ctx, spaceID, templateID, req)
}
func (m *MockClient) InviteUser(ctx context.Context, teamID string, req *api.InviteUserRequest) (*api.TeamUserResponse, error) {
	return m.InviteUserFn(ctx, teamID, req)
}
func (m *MockClient) GetTeamUser(ctx context.Context, teamID, userID string, includeShared bool) (*api.TeamUserResponse, error) {
	return m.GetTeamUserFn(ctx, teamID, userID, includeShared)
}
func (m *MockClient) EditUser(ctx context.Context, teamID, userID string, req *api.EditUserRequest) (*api.TeamUserResponse, error) {
	return m.EditUserFn(ctx, teamID, userID, req)
}
func (m *MockClient) RemoveUser(ctx context.Context, teamID, userID string) error {
	return m.RemoveUserFn(ctx, teamID, userID)
}
func (m *MockClient) GetCustomRoles(ctx context.Context, teamID string, includeMembers bool) (*api.CustomRolesResponse, error) {
	return m.GetCustomRolesFn(ctx, teamID, includeMembers)
}
func (m *MockClient) GetCustomTaskTypes(ctx context.Context, teamID string) (*api.CustomTaskTypesResponse, error) {
	return m.GetCustomTaskTypesFn(ctx, teamID)
}
func (m *MockClient) GetSharedHierarchy(ctx context.Context, teamID string) (*api.SharedHierarchyResponse, error) {
	return m.GetSharedHierarchyFn(ctx, teamID)
}
func (m *MockClient) GetWorkspaceSeats(ctx context.Context, teamID string) (*api.SeatsResponse, error) {
	return m.GetWorkspaceSeatsFn(ctx, teamID)
}
func (m *MockClient) GetWorkspacePlan(ctx context.Context, teamID string) (*api.PlanResponse, error) {
	return m.GetWorkspacePlanFn(ctx, teamID)
}

// Relationships
func (m *MockClient) AddDependency(ctx context.Context, taskID string, req *api.AddDependencyRequest, opts ...*api.TaskScopedOptions) (*api.DependencyResponse, error) {
	return m.AddDependencyFn(ctx, taskID, req)
}
func (m *MockClient) DeleteDependency(ctx context.Context, taskID, dependsOn, dependencyOf string, opts ...*api.TaskScopedOptions) error {
	return m.DeleteDependencyFn(ctx, taskID, dependsOn, dependencyOf)
}
func (m *MockClient) AddTaskLink(ctx context.Context, taskID, linksTo string, opts ...*api.TaskScopedOptions) (*api.TaskLinkResponse, error) {
	return m.AddTaskLinkFn(ctx, taskID, linksTo)
}
func (m *MockClient) DeleteTaskLink(ctx context.Context, taskID, linksTo string, opts ...*api.TaskScopedOptions) error {
	return m.DeleteTaskLinkFn(ctx, taskID, linksTo)
}

// Task Extras
func (m *MockClient) MergeTasks(ctx context.Context, taskID string, req *api.MergeTasksRequest, opts ...*api.TaskScopedOptions) error {
	return m.MergeTasksFn(ctx, taskID, req)
}
func (m *MockClient) GetTimeInStatus(ctx context.Context, taskID string, opts ...*api.TaskScopedOptions) (*api.TimeInStatusResponse, error) {
	return m.GetTimeInStatusFn(ctx, taskID)
}
func (m *MockClient) GetBulkTimeInStatus(ctx context.Context, taskIDs []string) (*api.BulkTimeInStatusResponse, error) {
	return m.GetBulkTimeInStatusFn(ctx, taskIDs)
}
func (m *MockClient) AddTaskToList(ctx context.Context, listID, taskID string) error {
	return m.AddTaskToListFn(ctx, listID, taskID)
}
func (m *MockClient) RemoveTaskFromList(ctx context.Context, listID, taskID string) error {
	return m.RemoveTaskFromListFn(ctx, listID, taskID)
}

// Attachments
func (m *MockClient) CreateTaskAttachment(ctx context.Context, taskID, filePath string, opts ...*api.TaskScopedOptions) (*api.Attachment, error) {
	return m.CreateTaskAttachmentFn(ctx, taskID, filePath)
}
