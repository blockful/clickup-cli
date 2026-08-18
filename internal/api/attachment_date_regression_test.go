package api

import (
	"encoding/json"
	"testing"
)

// ClickUp sends an attachment's scalar fields as either a JSON string or a JSON
// number, and not consistently per field or per endpoint: `date` arrived as
// "1749381600000" (#6) and `version` as 0 (2026-08-18). Either one typed as a
// plain Go scalar breaks unmarshalling of EVERY task response that carries an
// attachment, which takes out `task get` and `task update` for that task. Both
// fields are FlexString; this pins both shapes.
func TestAttachmentScalarsAcceptStringOrNumber(t *testing.T) {
	for _, tc := range []struct {
		name        string
		payload     string
		wantVersion FlexString
		wantDate    FlexString
	}{
		{
			name:        "strings",
			payload:     `{"id":"abc123","version":"0","date":"1749381600000","title":"file.md"}`,
			wantVersion: "0",
			wantDate:    "1749381600000",
		},
		{
			name:        "numbers",
			payload:     `{"id":"abc123","version":0,"date":1749381600000,"title":"file.md"}`,
			wantVersion: "0",
			wantDate:    "1749381600000",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var a Attachment
			if err := json.Unmarshal([]byte(tc.payload), &a); err != nil {
				t.Fatalf("failed to unmarshal attachment: %v", err)
			}
			if a.Version != tc.wantVersion {
				t.Fatalf("version: got %q, want %q", a.Version, tc.wantVersion)
			}
			if a.Date != tc.wantDate {
				t.Fatalf("date: got %q, want %q", a.Date, tc.wantDate)
			}
		})
	}
}

// A field that is neither shape must fail loudly rather than land as an empty
// string that reads like a missing value downstream.
func TestAttachmentScalarRejectsNonScalar(t *testing.T) {
	var a Attachment
	if err := json.Unmarshal([]byte(`{"id":"abc","version":{"n":1}}`), &a); err == nil {
		t.Fatal("expected an error for an object in a FlexString field")
	}
}
