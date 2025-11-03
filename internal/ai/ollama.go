package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestBody struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ResponseChunk struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"respones"`
	Done      bool   `json:"done"`
}

func GenerateResponse(prompt string) (string, error) {
	url := "http://localhost:11434/api/generate"

	reqBody := RequestBody{
		Model:  "llama3:latest",
		Prompt: prompt,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("failed to send request to ollama: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama API error: %s", string(bodyBytes))
	}

	var output string
	decoder := json.NewDecoder(resp.Body)

	for decoder.More() {
		var chunk ResponseChunk
		if err := decoder.Decode(&chunk); err != nil {
			break
		}
		output += chunk.Response
		if chunk.Done {
			break
		}
	}

	return output, nil
}
