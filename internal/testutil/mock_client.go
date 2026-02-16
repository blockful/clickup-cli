package testutil

import "github.com/blockful/clickup-cli/internal/api"

// MockClient implements api.ClientInterface for testing.
// Set the return values on the fields before calling the method.
type MockClient struct {
	GetUserFn        func() (*api.UserResponse, error)
	ListWorkspacesFn func() (*api.WorkspacesResponse, error)
	ListSpacesFn     func(workspaceID string) (*api.SpacesResponse, error)
	GetSpaceFn       func(spaceID string) (*api.Space, error)
	CreateSpaceFn    func(workspaceID string, req *api.CreateSpaceRequest) (*api.Space, error)
	ListFoldersFn    func(spaceID string) (*api.FoldersResponse, error)
	GetFolderFn      func(folderID string) (*api.Folder, error)
	CreateFolderFn   func(spaceID string, req *api.CreateFolderRequest) (*api.Folder, error)
	ListListsFn      func(folderID string) (*api.ListsResponse, error)
	GetListFn        func(listID string) (*api.List, error)
	CreateListFn     func(folderID string, req *api.CreateListRequest) (*api.List, error)
	ListTasksFn      func(listID string, opts *api.ListTasksOptions) (*api.TasksResponse, error)
	GetTaskFn        func(taskID string) (*api.Task, error)
	CreateTaskFn     func(listID string, req *api.CreateTaskRequest) (*api.Task, error)
	UpdateTaskFn     func(taskID string, req *api.UpdateTaskRequest) (*api.Task, error)
	DeleteTaskFn     func(taskID string) error
	ListCommentsFn   func(taskID string) (*api.CommentsResponse, error)
	CreateCommentFn  func(taskID string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error)
}

var _ api.ClientInterface = (*MockClient)(nil)

func (m *MockClient) GetUser() (*api.UserResponse, error)   { return m.GetUserFn() }
func (m *MockClient) ListWorkspaces() (*api.WorkspacesResponse, error) {
	return m.ListWorkspacesFn()
}
func (m *MockClient) ListSpaces(id string) (*api.SpacesResponse, error)   { return m.ListSpacesFn(id) }
func (m *MockClient) GetSpace(id string) (*api.Space, error)              { return m.GetSpaceFn(id) }
func (m *MockClient) CreateSpace(wid string, req *api.CreateSpaceRequest) (*api.Space, error) {
	return m.CreateSpaceFn(wid, req)
}
func (m *MockClient) ListFolders(id string) (*api.FoldersResponse, error) {
	return m.ListFoldersFn(id)
}
func (m *MockClient) GetFolder(id string) (*api.Folder, error) { return m.GetFolderFn(id) }
func (m *MockClient) CreateFolder(sid string, req *api.CreateFolderRequest) (*api.Folder, error) {
	return m.CreateFolderFn(sid, req)
}
func (m *MockClient) ListLists(id string) (*api.ListsResponse, error) { return m.ListListsFn(id) }
func (m *MockClient) GetList(id string) (*api.List, error)            { return m.GetListFn(id) }
func (m *MockClient) CreateList(fid string, req *api.CreateListRequest) (*api.List, error) {
	return m.CreateListFn(fid, req)
}
func (m *MockClient) ListTasks(id string, opts *api.ListTasksOptions) (*api.TasksResponse, error) {
	return m.ListTasksFn(id, opts)
}
func (m *MockClient) GetTask(id string) (*api.Task, error) { return m.GetTaskFn(id) }
func (m *MockClient) CreateTask(lid string, req *api.CreateTaskRequest) (*api.Task, error) {
	return m.CreateTaskFn(lid, req)
}
func (m *MockClient) UpdateTask(id string, req *api.UpdateTaskRequest) (*api.Task, error) {
	return m.UpdateTaskFn(id, req)
}
func (m *MockClient) DeleteTask(id string) error { return m.DeleteTaskFn(id) }
func (m *MockClient) ListComments(id string) (*api.CommentsResponse, error) {
	return m.ListCommentsFn(id)
}
func (m *MockClient) CreateComment(tid string, req *api.CreateCommentRequest) (*api.CreateCommentResponse, error) {
	return m.CreateCommentFn(tid, req)
}
