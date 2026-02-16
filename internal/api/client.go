package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const DefaultBaseURL = "https://api.clickup.com/api"

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

type APIError struct {
	Err     string `json:"err"`
	ECODE   string `json:"ECODE"`
	Message string `json:"message,omitempty"`
}

type ClientError struct {
	StatusCode int
	Code       string
	Message    string
}

func (e *ClientError) Error() string {
	return e.Message
}

func NewClient(token string) *Client {
	return &Client{
		BaseURL: DefaultBaseURL,
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Do(method, path string, body interface{}, result interface{}) error {
	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return &ClientError{Code: "MARSHAL_ERROR", Message: fmt.Sprintf("failed to marshal request body: %v", err)}
		}
		reqBody = bytes.NewReader(data)
	}

	url := c.BaseURL + path
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return &ClientError{Code: "REQUEST_ERROR", Message: fmt.Sprintf("failed to create request: %v", err)}
	}

	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return &ClientError{Code: "NETWORK_ERROR", Message: fmt.Sprintf("request failed: %v", err)}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ClientError{Code: "READ_ERROR", Message: fmt.Sprintf("failed to read response: %v", err)}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		code := errorCodeFromStatus(resp.StatusCode)
		// Try to parse ClickUp error
		var apiErr APIError
		if json.Unmarshal(respBody, &apiErr) == nil && (apiErr.Err != "" || apiErr.Message != "") {
			msg := apiErr.Err
			if msg == "" {
				msg = apiErr.Message
			}
			return &ClientError{StatusCode: resp.StatusCode, Code: code, Message: msg}
		}
		return &ClientError{StatusCode: resp.StatusCode, Code: code, Message: fmt.Sprintf("API returned status %d", resp.StatusCode)}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return &ClientError{Code: "UNMARSHAL_ERROR", Message: fmt.Sprintf("failed to parse response: %v", err)}
		}
	}

	return nil
}

func errorCodeFromStatus(status int) string {
	switch status {
	case 401:
		return "UNAUTHORIZED"
	case 403:
		return "FORBIDDEN"
	case 404:
		return "NOT_FOUND"
	case 429:
		return "RATE_LIMITED"
	default:
		return "API_ERROR"
	}
}
