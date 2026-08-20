package api

import (
	"context"
	"encoding/json"
	"fmt"
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
		_ = json.NewEncoder(w).Encode(WebhooksResponse{Webhooks: []Webhook{{ID: "wh1"}}})
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
		_ = json.NewEncoder(w).Encode(CreateWebhookResponse{ID: "wh1"})
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

// The two shapes that broke a live repoint on 20/ago/2026: a scalar `events`
// (400 "Invalid events" for a value GET had just returned) and an empty
// `status` riding along on an endpoint-only update (500). Both are payload
// bugs, so assert on the bytes on the wire, not on the response.
func TestUpdateWebhookBody(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		req  *UpdateWebhookRequest
		want map[string]interface{}
	}{
		{
			name: "endpoint only omits events and status",
			req:  &UpdateWebhookRequest{Endpoint: "https://example.com/hook"},
			want: map[string]interface{}{"endpoint": "https://example.com/hook"},
		},
		{
			name: "events marshal as an array",
			req:  &UpdateWebhookRequest{Events: []string{"taskStatusUpdated"}},
			want: map[string]interface{}{"events": []interface{}{"taskStatusUpdated"}},
		},
		{
			name: "status alone",
			req:  &UpdateWebhookRequest{Status: "active"},
			want: map[string]interface{}{"status": "active"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got map[string]interface{}
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "PUT" || r.URL.Path != "/v2/webhook/wh1" {
					t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
				}
				if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
					t.Fatalf("decode body: %v", err)
				}
				_ = json.NewEncoder(w).Encode(UpdateWebhookResponse{ID: "wh1"})
			}))
			defer srv.Close()

			c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
			if _, err := c.UpdateWebhook(ctx, "wh1", tc.req); err != nil {
				t.Fatal(err)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("body = %v, want exactly %v", got, tc.want)
			}
			for k, v := range tc.want {
				if fmt.Sprintf("%v", got[k]) != fmt.Sprintf("%v", v) {
					t.Errorf("body[%q] = %v, want %v", k, got[k], v)
				}
			}
		})
	}
}
