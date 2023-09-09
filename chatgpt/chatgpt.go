package chatgpt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const openAIURL = "https://api.openai.com/v1/chat/completions"

type Payload struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

func GetAPIKey(apiKey string) (string, error) {
	if apiKey != "" {
		return apiKey, nil
	}
	filePath := path.Join(os.Getenv("HOME"), ".openai", "auth.json")
	expandedPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", errors.Errorf("Error expanding path %s: %v", filePath, err)
	}
	data, err := os.ReadFile(expandedPath)
	if err != nil {
		return "", errors.Errorf("Error reading file %s: %v", expandedPath, err)
	}
	var config struct {
		PrivateAPIKey string `json:"private_api_key"`
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return "", errors.Errorf("Error unmarshalling file %s: %v", expandedPath, err)
	}
	return config.PrivateAPIKey, nil
}

//go:generate genopts --function AskQuestion verbose inferResultType model:string:gpt-3.5-turbo
func AskQuestion(apiKey string, question string, res interface{}, optss ...AskQuestionOption) (string, error) {
	opts := MakeAskQuestionOptions(optss...)

	client := resty.New()

	type Response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	content := question
	if opts.InferResultType() {
		if res == nil {
			return "", errors.Errorf("Cannot infer result type when res is nil")
		}
		SetDefaults(res)
		b, err := json.Marshal(res)
		if err != nil {
			return "", errors.Errorf("Error marshalling res: %v", err)
		}
		content += fmt.Sprintf("\n Return the result in JSON as: %s", string(b))
	}

	if opts.Verbose() {
		log.Printf("[AskQuestion] content: %s", content)
	}

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+apiKey).
		SetResult(&Response{}).
		SetBody(map[string]interface{}{
			"model":    opts.Model(),
			"messages": []interface{}{map[string]interface{}{"role": "system", "content": content}},
		}).
		Post(openAIURL)

	if err != nil {
		return "", errors.Errorf("Error while sending send the request: %v", err)
	}

	body := response.Body()

	if opts.Verbose() {
		log.Printf("[AskQuestion] body: %s", body)
	}

	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", errors.Errorf("Error unmarshalling response: %v", err)
	}

	ret := resp.Choices[0].Message.Content
	if err := json.Unmarshal([]byte(ret), &res); err != nil {
		return "", errors.Errorf("Error unmarshalling response: %v", err)
	}

	if opts.Verbose() {
		log.Printf("[AskQuestion] ret: %s", ret)
	}

	return ret, nil
}
