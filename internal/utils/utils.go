package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type CloudflareResult struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
	Error  string `json:"error,omitempty"`
}

func CloudflareRequest(method, url, headersJSON, payloadJSON string) (*CloudflareResult, error) {
	args := []string{"./internal/utils/cfbypass.py", method, url}
	if headersJSON != "" {
		args = append(args, parseAndEscapeJSON(headersJSON))
	}
	if payloadJSON != "" {
		args = append(args, payloadJSON)
	}

	cmd := exec.Command("python", args...)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("python error: %v", err)
	}

	var result CloudflareResult
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("python reported error: %s", result.Error)
	}
	return &result, nil
}

func parseAndEscapeJSON(input string) string {
	var headers map[string]string
	if err := json.Unmarshal([]byte(input), &headers); err != nil {
		log.Fatalf("Invalid headers JSON: %v", err)
	}
	out, err := json.Marshal(headers)
	if err != nil {
		log.Fatalf("Failed to re-encode headers: %v", err)
	}
	return string(out)
}
