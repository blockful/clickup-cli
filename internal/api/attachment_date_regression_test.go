package api

import (
	"encoding/json"
	"testing"
)

// TestAttachmentDateUnmarshalsStringFromAPI is a regression test: ClickUp's API
// returns an attachment "date" as a JSON string (e.g. "1749381600000"), not a
// number. When Date was typed int64, unmarshalling any task/attachment response
// failed with "cannot unmarshal string into Go struct field Attachment.date".
func TestAttachmentDateUnmarshalsStringFromAPI(t *testing.T) {
	payload := []byte(`{"id":"abc123","version":"0","date":"1749381600000","title":"file.md","extension":"md"}`)
	var a Attachment
	if err := json.Unmarshal(payload, &a); err != nil {
		t.Fatalf("failed to unmarshal attachment with string date: %v", err)
	}
	if a.Date != "1749381600000" {
		t.Fatalf("unexpected date: %q", a.Date)
	}
}
