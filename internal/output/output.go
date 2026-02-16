package output

import (
	"encoding/json"
	"fmt"
	"os"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func JSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		PrintError("MARSHAL_ERROR", fmt.Sprintf("failed to marshal output: %v", err))
		return
	}
	fmt.Println(string(data))
}

func PrintError(code, message string) {
	resp := ErrorResponse{
		Error: message,
		Code:  code,
	}
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintln(os.Stderr, string(data))
}

func PrintErrorAndExit(code, message string, exitCode int) {
	PrintError(code, message)
	os.Exit(exitCode)
}
