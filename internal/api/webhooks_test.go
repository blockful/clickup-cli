package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWebhooks(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/123/webhook" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(WebhooksResponse{Webhooks: []Webhook{{ID: "wh1"}}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetWebhooks(ctx, "123")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Webhooks) != 1 {
		t.Errorf("count = %d", len(resp.Webhooks))
	}
}

func TestCreateWebhook(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/team/123/webhook" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		json.NewEncoder(w).Encode(CreateWebhookResponse{ID: "wh1"})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.CreateWebhook(ctx, "123", &CreateWebhookRequest{Endpoint: "https://example.com", Events: []string{"*"}})
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != "wh1" {
		t.Errorf("id = %s", resp.ID)
	}
}

func TestDeleteWebhook(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/webhook/wh1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteWebhook(ctx, "wh1"); err != nil {
		t.Fatal(err)
	}
}
