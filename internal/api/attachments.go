package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Attachment struct {
	ID             string `json:"id"`
	Version        string `json:"version"`
	Date           int64  `json:"date"`
	Title          string `json:"title"`
	Extension      string `json:"extension"`
	ThumbnailSmall string `json:"thumbnail_small"`
	ThumbnailLarge string `json:"thumbnail_large"`
	URL            string `json:"url"`
}

func (c *Client) CreateTaskAttachment(ctx context.Context, taskID, filePath string, opts ...*TaskScopedOptions) (*Attachment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, &ClientError{Code: "FILE_ERROR", Message: fmt.Sprintf("failed to open file: %v", err)}
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("attachment", filepath.Base(filePath))
	if err != nil {
		return nil, &ClientError{Code: "MULTIPART_ERROR", Message: fmt.Sprintf("failed to create form file: %v", err)}
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, &ClientError{Code: "COPY_ERROR", Message: fmt.Sprintf("failed to copy file: %v", err)}
	}
	writer.Close()

	var o *TaskScopedOptions
	if len(opts) > 0 {
		o = opts[0]
	}
	url := c.BaseURL + fmt.Sprintf("/v2/task/%s/attachment", taskID) + taskScopedQuery(o)
	req, err := http.NewRequestWithContext(ctx, "POST", url, &buf)
	if err != nil {
		return nil, &ClientError{Code: "REQUEST_ERROR", Message: fmt.Sprintf("failed to create request: %v", err)}
	}
	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return nil, &ClientError{Code: "CANCELLED", Message: "request cancelled"}
		}
		return nil, &ClientError{Code: "NETWORK_ERROR", Message: fmt.Sprintf("request failed: %v", err), Retryable: true}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &ClientError{Code: "READ_ERROR", Message: fmt.Sprintf("failed to read response: %v", err)}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		code := errorCodeFromStatus(resp.StatusCode)
		return nil, &ClientError{StatusCode: resp.StatusCode, Code: code, Message: fmt.Sprintf("API returned status %d: %s", resp.StatusCode, string(respBody))}
	}

	var attachment Attachment
	if err := json.Unmarshal(respBody, &attachment); err != nil {
		return nil, &ClientError{Code: "UNMARSHAL_ERROR", Message: fmt.Sprintf("failed to parse response: %v", err)}
	}
	return &attachment, nil
}
