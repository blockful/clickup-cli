package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListListComments(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/list/l1/comment" {
			t.Errorf("path: %s", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(CommentsResponse{})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.ListListComments(ctx, "l1", "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateComment_WithAssignee(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreateCommentRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		if req.Assignee == nil || *req.Assignee != 5 {
			t.Error("assignee not set")
		}
		if !req.NotifyAll {
			t.Error("notify_all should be true")
		}
		_ = json.NewEncoder(w).Encode(CreateCommentResponse{ID: "c1"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	a := 5
	_, err := c.CreateComment(ctx, "t1", &CreateCommentRequest{CommentText: "hi", Assignee: &a, NotifyAll: true})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateComment(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method: %s", r.Method)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.UpdateComment(ctx, "c1", &UpdateCommentRequest{CommentText: "updated"}); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteComment(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method: %s", r.Method)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteComment(ctx, "c1"); err != nil {
		t.Fatal(err)
	}
}
