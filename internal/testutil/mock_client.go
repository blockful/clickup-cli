package testutil

import "github.com/blockful/clickup-cli/internal/api"

// MockClient implements api.ClientInterface for testing.
type MockClient struct {
	GetUserFn              func() (*api.UserResponse, error)
	ListWorkspacesFn       func() (*api.WorkspacesResponse, error)
	ListSpacesFn           func(workspaceID string) (*api.SpacesResponse, error)
	GetSpaceFn             func(spaceID string) (*api.Space, error)
	CreateSpaceFn          func(workspaceID string, req *api.CreateSpaceRequest) (*api.Space, error)
	UpdateSpaceFn          func(spaceID string, req *api.UpdateSpaceRequest) (*api.Space, error)
	DeleteSpaceFn          func(spaceID string) error
	ListFoldersFn          func(spaceID string) (*api.FoldersResponse, error)
	GetFolderFn            func(folderID string) (*api.Folder, error)
	CreateFolderFn         func(spaceID string, req *api.CreateFolderRequest) (*api.Folder, error)
	UpdateFolderFn         func(folderID string, req *api.UpdateFolderRequest) (*api.Folder, error)
	DeleteFolderFn         func(folderID string) error
	ListListsFn            func(folderID string) (*api.ListsResponse, error)
	ListFolderlessListsFn  func(spaceID string) (*api.ListsResponse, error)
	GetListFn              func(listID string) (*api.List, error)
	CreateListFn           func(folderID string, req *api.CreateListRequest) (*api.List, error)
	CreateFolderlessListFn func(spaceID string, req *api.CreateListRequest) (*api.List, error)
	UpdateListFn           func(listID string, req *api.UpdateListRequest) (*api.List, error)
	DeleteListFn           func(listID string) error
	ListTasksFn            func(listID string, opts *api.ListTasksOptions) (*api.TasksResponse, error)
	GetTaskFn              func(taskID string, opts ...api.GetTaskOptions) (*api.Task, error)
	CreateTaskFn           func(listID string, req *api.CreateTaskRequest) (*api.Task, error)
	UpdateTaskFn           func(taskID string, req *api.UpdateTaskRequest, opts ...api.UpdateTaskOptions) (*api.Task, error)
	DeleteTaskFn           func(taskID string) error
	SearchTasksFn          func(teamID string, opts *api.SearchTasksOptions) (*api.TasksResponse, error)
	ListCommentsFn         func(taskID string) (*api.CommentsResponse, error)
	ListListCommentsFn     func(listID string) (*api.CommentsResponse, error)
	CreateCommentFn        func(taskID string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error)
	CreateListCommentFn    func(listID string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error)
	UpdateCommentFn        func(commentID string, req *api.UpdateCommentRequest) error
	DeleteCommentFn        func(commentID string) error

	// Custom Fields
	GetListCustomFieldsFn      func(listID string) (*api.CustomFieldsResponse, error)
	GetFolderCustomFieldsFn    func(folderID string) (*api.CustomFieldsResponse, error)
	GetSpaceCustomFieldsFn     func(spaceID string) (*api.CustomFieldsResponse, error)
	GetWorkspaceCustomFieldsFn func(teamID string) (*api.CustomFieldsResponse, error)
	SetCustomFieldValueFn      func(taskID, fieldID string, req *api.SetCustomFieldRequest) error
	RemoveCustomFieldValueFn   func(taskID, fieldID string) error

	// Tags
	GetSpaceTagsFn      func(spaceID string) (*api.TagsResponse, error)
	CreateSpaceTagFn    func(spaceID string, req *api.CreateTagRequest) error
	UpdateSpaceTagFn    func(spaceID, tagName string, req *api.UpdateTagRequest) error
	DeleteSpaceTagFn    func(spaceID, tagName string) error
	AddTagToTaskFn      func(taskID, tagName string) error
	RemoveTagFromTaskFn func(taskID, tagName string) error

	// Checklists
	CreateChecklistFn     func(taskID string, req *api.CreateChecklistRequest) (*api.ChecklistResponse, error)
	EditChecklistFn       func(checklistID string, req *api.EditChecklistRequest) error
	DeleteChecklistFn     func(checklistID string) error
	CreateChecklistItemFn func(checklistID string, req *api.CreateChecklistItemRequest) (*api.ChecklistResponse, error)
	EditChecklistItemFn   func(checklistID, checklistItemID string, req *api.EditChecklistItemRequest) (*api.ChecklistResponse, error)
	DeleteChecklistItemFn func(checklistID, checklistItemID string) error

	// Docs
	CreateDocFn        func(workspaceID string, req *api.CreateDocRequest) (*api.Doc, error)
	SearchDocsFn       func(workspaceID string) (*api.DocsResponse, error)
	GetDocFn           func(workspaceID, docID string) (*api.Doc, error)
	CreatePageFn       func(workspaceID, docID string, req *api.CreatePageRequest) (*api.DocPage, error)
	GetPageFn          func(workspaceID, docID, pageID string) (*api.DocPage, error)
	EditPageFn         func(workspaceID, docID, pageID string, req *api.EditPageRequest) (*api.DocPage, error)
	GetDocPageListingFn func(workspaceID, docID string) (*api.DocPagesResponse, error)

	// Time Tracking
	GetTimeEntriesFn    func(teamID string, opts *api.ListTimeEntriesOptions) (*api.TimeEntriesResponse, error)
	CreateTimeEntryFn   func(teamID string, req *api.CreateTimeEntryRequest) (*api.TimeEntry, error)
	GetTimeEntryFn      func(teamID, timerID string) (*api.SingleTimeEntryResponse, error)
	UpdateTimeEntryFn   func(teamID, timerID string, req *api.UpdateTimeEntryRequest) error
	DeleteTimeEntryFn   func(teamID, timerID string) error
	StartTimerFn        func(teamID string, req *api.StartTimerRequest) (*api.SingleTimeEntryResponse, error)
	StopTimerFn         func(teamID string) (*api.SingleTimeEntryResponse, error)
	GetRunningTimerFn   func(teamID string, assignee string) (*api.SingleTimeEntryResponse, error)
	GetTimeEntryTagsFn  func(teamID string) (*api.TimeEntryTagsResponse, error)

	// Webhooks
	GetWebhooksFn    func(teamID string) (*api.WebhooksResponse, error)
	CreateWebhookFn  func(teamID string, req *api.CreateWebhookRequest) (*api.CreateWebhookResponse, error)
	UpdateWebhookFn  func(webhookID string, req *api.UpdateWebhookRequest) (*api.UpdateWebhookResponse, error)
	DeleteWebhookFn  func(webhookID string) error

	// Views
	GetTeamViewsFn    func(teamID string) (*api.ViewsResponse, error)
	GetSpaceViewsFn   func(spaceID string) (*api.ViewsResponse, error)
	GetFolderViewsFn  func(folderID string) (*api.ViewsResponse, error)
	GetListViewsFn    func(listID string) (*api.ViewsResponse, error)
	GetViewFn         func(viewID string) (*api.ViewResponse, error)
	CreateTeamViewFn  func(teamID string, req *api.CreateViewRequest) (*api.ViewResponse, error)
	CreateSpaceViewFn func(spaceID string, req *api.CreateViewRequest) (*api.ViewResponse, error)
	CreateFolderViewFn func(folderID string, req *api.CreateViewRequest) (*api.ViewResponse, error)
	CreateListViewFn  func(listID string, req *api.CreateViewRequest) (*api.ViewResponse, error)
	UpdateViewFn      func(viewID string, req *api.UpdateViewRequest) (*api.ViewResponse, error)
	DeleteViewFn      func(viewID string) error
	GetViewTasksFn    func(viewID string, page int) (*api.ViewTasksResponse, error)

	// Goals
	GetGoalsFn       func(teamID string, includeCompleted bool) (*api.GoalsResponse, error)
	GetGoalFn        func(goalID string) (*api.GoalResponse, error)
	CreateGoalFn     func(teamID string, req *api.CreateGoalRequest) (*api.GoalResponse, error)
	UpdateGoalFn     func(goalID string, req *api.UpdateGoalRequest) (*api.GoalResponse, error)
	DeleteGoalFn     func(goalID string) error
	CreateKeyResultFn func(goalID string, req *api.CreateKeyResultRequest) (*api.KeyResultResponse, error)

	// Members
	GetListMembersFn func(listID string) (*api.MembersResponse, error)
	GetTaskMembersFn func(taskID string) (*api.MembersResponse, error)

	// Groups
	GetGroupsFn    func(teamID string) (*api.GroupsResponse, error)
	CreateGroupFn  func(teamID string, req *api.CreateGroupRequest) (*api.Group, error)
	UpdateGroupFn  func(groupID string, req *api.UpdateGroupRequest) (*api.Group, error)
	DeleteGroupFn  func(groupID string) error

	// Guests
	InviteGuestFn func(teamID string, req *api.InviteGuestRequest) error
	GetGuestFn    func(teamID, guestID string) (*api.GuestResponse, error)
	EditGuestFn   func(teamID, guestID string, req *api.EditGuestRequest) (*api.GuestResponse, error)
	RemoveGuestFn func(teamID, guestID string) error
}

var _ api.ClientInterface = (*MockClient)(nil)

func (m *MockClient) GetUser() (*api.UserResponse, error) { return m.GetUserFn() }
func (m *MockClient) ListWorkspaces() (*api.WorkspacesResponse, error) {
	return m.ListWorkspacesFn()
}
func (m *MockClient) ListSpaces(id string) (*api.SpacesResponse, error) { return m.ListSpacesFn(id) }
func (m *MockClient) GetSpace(id string) (*api.Space, error)            { return m.GetSpaceFn(id) }
func (m *MockClient) CreateSpace(wid string, req *api.CreateSpaceRequest) (*api.Space, error) {
	return m.CreateSpaceFn(wid, req)
}
func (m *MockClient) UpdateSpace(id string, req *api.UpdateSpaceRequest) (*api.Space, error) {
	return m.UpdateSpaceFn(id, req)
}
func (m *MockClient) DeleteSpace(id string) error { return m.DeleteSpaceFn(id) }
func (m *MockClient) ListFolders(id string) (*api.FoldersResponse, error) {
	return m.ListFoldersFn(id)
}
func (m *MockClient) GetFolder(id string) (*api.Folder, error) { return m.GetFolderFn(id) }
func (m *MockClient) CreateFolder(sid string, req *api.CreateFolderRequest) (*api.Folder, error) {
	return m.CreateFolderFn(sid, req)
}
func (m *MockClient) UpdateFolder(id string, req *api.UpdateFolderRequest) (*api.Folder, error) {
	return m.UpdateFolderFn(id, req)
}
func (m *MockClient) DeleteFolder(id string) error { return m.DeleteFolderFn(id) }
func (m *MockClient) ListLists(id string) (*api.ListsResponse, error) { return m.ListListsFn(id) }
func (m *MockClient) ListFolderlessLists(id string) (*api.ListsResponse, error) {
	return m.ListFolderlessListsFn(id)
}
func (m *MockClient) GetList(id string) (*api.List, error) { return m.GetListFn(id) }
func (m *MockClient) CreateList(fid string, req *api.CreateListRequest) (*api.List, error) {
	return m.CreateListFn(fid, req)
}
func (m *MockClient) CreateFolderlessList(sid string, req *api.CreateListRequest) (*api.List, error) {
	return m.CreateFolderlessListFn(sid, req)
}
func (m *MockClient) UpdateList(id string, req *api.UpdateListRequest) (*api.List, error) {
	return m.UpdateListFn(id, req)
}
func (m *MockClient) DeleteList(id string) error { return m.DeleteListFn(id) }
func (m *MockClient) ListTasks(id string, opts *api.ListTasksOptions) (*api.TasksResponse, error) {
	return m.ListTasksFn(id, opts)
}
func (m *MockClient) GetTask(id string, opts ...api.GetTaskOptions) (*api.Task, error) {
	return m.GetTaskFn(id, opts...)
}
func (m *MockClient) CreateTask(lid string, req *api.CreateTaskRequest) (*api.Task, error) {
	return m.CreateTaskFn(lid, req)
}
func (m *MockClient) UpdateTask(id string, req *api.UpdateTaskRequest, opts ...api.UpdateTaskOptions) (*api.Task, error) {
	return m.UpdateTaskFn(id, req, opts...)
}
func (m *MockClient) DeleteTask(id string) error { return m.DeleteTaskFn(id) }
func (m *MockClient) SearchTasks(teamID string, opts *api.SearchTasksOptions) (*api.TasksResponse, error) {
	return m.SearchTasksFn(teamID, opts)
}
func (m *MockClient) ListComments(id string) (*api.CommentsResponse, error) {
	return m.ListCommentsFn(id)
}
func (m *MockClient) ListListComments(id string) (*api.CommentsResponse, error) {
	return m.ListListCommentsFn(id)
}
func (m *MockClient) CreateComment(tid string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error) {
	return m.CreateCommentFn(tid, req)
}
func (m *MockClient) CreateListComment(lid string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error) {
	return m.CreateListCommentFn(lid, req)
}
func (m *MockClient) UpdateComment(id string, req *api.UpdateCommentRequest) error {
	return m.UpdateCommentFn(id, req)
}
func (m *MockClient) DeleteComment(id string) error { return m.DeleteCommentFn(id) }

// Custom Fields
func (m *MockClient) GetListCustomFields(id string) (*api.CustomFieldsResponse, error) {
	return m.GetListCustomFieldsFn(id)
}
func (m *MockClient) GetFolderCustomFields(id string) (*api.CustomFieldsResponse, error) {
	return m.GetFolderCustomFieldsFn(id)
}
func (m *MockClient) GetSpaceCustomFields(id string) (*api.CustomFieldsResponse, error) {
	return m.GetSpaceCustomFieldsFn(id)
}
func (m *MockClient) GetWorkspaceCustomFields(id string) (*api.CustomFieldsResponse, error) {
	return m.GetWorkspaceCustomFieldsFn(id)
}
func (m *MockClient) SetCustomFieldValue(taskID, fieldID string, req *api.SetCustomFieldRequest) error {
	return m.SetCustomFieldValueFn(taskID, fieldID, req)
}
func (m *MockClient) RemoveCustomFieldValue(taskID, fieldID string) error {
	return m.RemoveCustomFieldValueFn(taskID, fieldID)
}

// Tags
func (m *MockClient) GetSpaceTags(id string) (*api.TagsResponse, error) {
	return m.GetSpaceTagsFn(id)
}
func (m *MockClient) CreateSpaceTag(id string, req *api.CreateTagRequest) error {
	return m.CreateSpaceTagFn(id, req)
}
func (m *MockClient) UpdateSpaceTag(id, name string, req *api.UpdateTagRequest) error {
	return m.UpdateSpaceTagFn(id, name, req)
}
func (m *MockClient) DeleteSpaceTag(id, name string) error { return m.DeleteSpaceTagFn(id, name) }
func (m *MockClient) AddTagToTask(taskID, tagName string) error {
	return m.AddTagToTaskFn(taskID, tagName)
}
func (m *MockClient) RemoveTagFromTask(taskID, tagName string) error {
	return m.RemoveTagFromTaskFn(taskID, tagName)
}

// Checklists
func (m *MockClient) CreateChecklist(taskID string, req *api.CreateChecklistRequest) (*api.ChecklistResponse, error) {
	return m.CreateChecklistFn(taskID, req)
}
func (m *MockClient) EditChecklist(id string, req *api.EditChecklistRequest) error {
	return m.EditChecklistFn(id, req)
}
func (m *MockClient) DeleteChecklist(id string) error { return m.DeleteChecklistFn(id) }
func (m *MockClient) CreateChecklistItem(id string, req *api.CreateChecklistItemRequest) (*api.ChecklistResponse, error) {
	return m.CreateChecklistItemFn(id, req)
}
func (m *MockClient) EditChecklistItem(cid, iid string, req *api.EditChecklistItemRequest) (*api.ChecklistResponse, error) {
	return m.EditChecklistItemFn(cid, iid, req)
}
func (m *MockClient) DeleteChecklistItem(cid, iid string) error {
	return m.DeleteChecklistItemFn(cid, iid)
}

// Docs
func (m *MockClient) CreateDoc(wid string, req *api.CreateDocRequest) (*api.Doc, error) {
	return m.CreateDocFn(wid, req)
}
func (m *MockClient) SearchDocs(wid string) (*api.DocsResponse, error) { return m.SearchDocsFn(wid) }
func (m *MockClient) GetDoc(wid, did string) (*api.Doc, error)         { return m.GetDocFn(wid, did) }
func (m *MockClient) CreatePage(wid, did string, req *api.CreatePageRequest) (*api.DocPage, error) {
	return m.CreatePageFn(wid, did, req)
}
func (m *MockClient) GetPage(wid, did, pid string) (*api.DocPage, error) {
	return m.GetPageFn(wid, did, pid)
}
func (m *MockClient) EditPage(wid, did, pid string, req *api.EditPageRequest) (*api.DocPage, error) {
	return m.EditPageFn(wid, did, pid, req)
}
func (m *MockClient) GetDocPageListing(wid, did string) (*api.DocPagesResponse, error) {
	return m.GetDocPageListingFn(wid, did)
}

// Time Tracking
func (m *MockClient) GetTimeEntries(teamID string, opts *api.ListTimeEntriesOptions) (*api.TimeEntriesResponse, error) {
	return m.GetTimeEntriesFn(teamID, opts)
}
func (m *MockClient) CreateTimeEntry(teamID string, req *api.CreateTimeEntryRequest) (*api.TimeEntry, error) {
	return m.CreateTimeEntryFn(teamID, req)
}
func (m *MockClient) GetTimeEntry(teamID, timerID string) (*api.SingleTimeEntryResponse, error) {
	return m.GetTimeEntryFn(teamID, timerID)
}
func (m *MockClient) UpdateTimeEntry(teamID, timerID string, req *api.UpdateTimeEntryRequest) error {
	return m.UpdateTimeEntryFn(teamID, timerID, req)
}
func (m *MockClient) DeleteTimeEntry(teamID, timerID string) error {
	return m.DeleteTimeEntryFn(teamID, timerID)
}
func (m *MockClient) StartTimer(teamID string, req *api.StartTimerRequest) (*api.SingleTimeEntryResponse, error) {
	return m.StartTimerFn(teamID, req)
}
func (m *MockClient) StopTimer(teamID string) (*api.SingleTimeEntryResponse, error) {
	return m.StopTimerFn(teamID)
}
func (m *MockClient) GetRunningTimer(teamID string, assignee string) (*api.SingleTimeEntryResponse, error) {
	return m.GetRunningTimerFn(teamID, assignee)
}
func (m *MockClient) GetTimeEntryTags(teamID string) (*api.TimeEntryTagsResponse, error) {
	return m.GetTimeEntryTagsFn(teamID)
}

// Webhooks
func (m *MockClient) GetWebhooks(teamID string) (*api.WebhooksResponse, error) {
	return m.GetWebhooksFn(teamID)
}
func (m *MockClient) CreateWebhook(teamID string, req *api.CreateWebhookRequest) (*api.CreateWebhookResponse, error) {
	return m.CreateWebhookFn(teamID, req)
}
func (m *MockClient) UpdateWebhook(id string, req *api.UpdateWebhookRequest) (*api.UpdateWebhookResponse, error) {
	return m.UpdateWebhookFn(id, req)
}
func (m *MockClient) DeleteWebhook(id string) error { return m.DeleteWebhookFn(id) }

// Views
func (m *MockClient) GetTeamViews(id string) (*api.ViewsResponse, error) {
	return m.GetTeamViewsFn(id)
}
func (m *MockClient) GetSpaceViews(id string) (*api.ViewsResponse, error) {
	return m.GetSpaceViewsFn(id)
}
func (m *MockClient) GetFolderViews(id string) (*api.ViewsResponse, error) {
	return m.GetFolderViewsFn(id)
}
func (m *MockClient) GetListViews(id string) (*api.ViewsResponse, error) {
	return m.GetListViewsFn(id)
}
func (m *MockClient) GetView(id string) (*api.ViewResponse, error) { return m.GetViewFn(id) }
func (m *MockClient) CreateTeamView(id string, req *api.CreateViewRequest) (*api.ViewResponse, error) {
	return m.CreateTeamViewFn(id, req)
}
func (m *MockClient) CreateSpaceView(id string, req *api.CreateViewRequest) (*api.ViewResponse, error) {
	return m.CreateSpaceViewFn(id, req)
}
func (m *MockClient) CreateFolderView(id string, req *api.CreateViewRequest) (*api.ViewResponse, error) {
	return m.CreateFolderViewFn(id, req)
}
func (m *MockClient) CreateListView(id string, req *api.CreateViewRequest) (*api.ViewResponse, error) {
	return m.CreateListViewFn(id, req)
}
func (m *MockClient) UpdateView(id string, req *api.UpdateViewRequest) (*api.ViewResponse, error) {
	return m.UpdateViewFn(id, req)
}
func (m *MockClient) DeleteView(id string) error { return m.DeleteViewFn(id) }
func (m *MockClient) GetViewTasks(id string, page int) (*api.ViewTasksResponse, error) {
	return m.GetViewTasksFn(id, page)
}

// Goals
func (m *MockClient) GetGoals(teamID string, includeCompleted bool) (*api.GoalsResponse, error) {
	return m.GetGoalsFn(teamID, includeCompleted)
}
func (m *MockClient) GetGoal(id string) (*api.GoalResponse, error) { return m.GetGoalFn(id) }
func (m *MockClient) CreateGoal(teamID string, req *api.CreateGoalRequest) (*api.GoalResponse, error) {
	return m.CreateGoalFn(teamID, req)
}
func (m *MockClient) UpdateGoal(id string, req *api.UpdateGoalRequest) (*api.GoalResponse, error) {
	return m.UpdateGoalFn(id, req)
}
func (m *MockClient) DeleteGoal(id string) error { return m.DeleteGoalFn(id) }
func (m *MockClient) CreateKeyResult(goalID string, req *api.CreateKeyResultRequest) (*api.KeyResultResponse, error) {
	return m.CreateKeyResultFn(goalID, req)
}

// Members
func (m *MockClient) GetListMembers(id string) (*api.MembersResponse, error) {
	return m.GetListMembersFn(id)
}
func (m *MockClient) GetTaskMembers(id string) (*api.MembersResponse, error) {
	return m.GetTaskMembersFn(id)
}

// Groups
func (m *MockClient) GetGroups(teamID string) (*api.GroupsResponse, error) {
	return m.GetGroupsFn(teamID)
}
func (m *MockClient) CreateGroup(teamID string, req *api.CreateGroupRequest) (*api.Group, error) {
	return m.CreateGroupFn(teamID, req)
}
func (m *MockClient) UpdateGroup(id string, req *api.UpdateGroupRequest) (*api.Group, error) {
	return m.UpdateGroupFn(id, req)
}
func (m *MockClient) DeleteGroup(id string) error { return m.DeleteGroupFn(id) }

// Guests
func (m *MockClient) InviteGuest(teamID string, req *api.InviteGuestRequest) error {
	return m.InviteGuestFn(teamID, req)
}
func (m *MockClient) GetGuest(teamID, guestID string) (*api.GuestResponse, error) {
	return m.GetGuestFn(teamID, guestID)
}
func (m *MockClient) EditGuest(teamID, guestID string, req *api.EditGuestRequest) (*api.GuestResponse, error) {
	return m.EditGuestFn(teamID, guestID, req)
}
func (m *MockClient) RemoveGuest(teamID, guestID string) error {
	return m.RemoveGuestFn(teamID, guestID)
}
