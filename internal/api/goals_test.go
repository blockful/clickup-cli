package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGoals(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/123/goal" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(GoalsResponse{Goals: []Goal{{ID: "g1", Name: "Q1"}}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetGoals(ctx, "123", false)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Goals) != 1 {
		t.Errorf("count = %d", len(resp.Goals))
	}
}

func TestGetGoal(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/goal/g1" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(GoalResponse{Goal: Goal{ID: "g1"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetGoal(ctx, "g1")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Goal.ID != "g1" {
		t.Errorf("id = %s", resp.Goal.ID)
	}
}

func TestCreateGoal(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/team/123/goal" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		json.NewEncoder(w).Encode(GoalResponse{Goal: Goal{ID: "g2"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.CreateGoal(ctx, "123", &CreateGoalRequest{Name: "Q2"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Goal.ID != "g2" {
		t.Errorf("id = %s", resp.Goal.ID)
	}
}

func TestDeleteGoal(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/goal/g1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteGoal(ctx, "g1"); err != nil {
		t.Fatal(err)
	}
}
