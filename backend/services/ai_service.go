package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type AIResponse struct {
	Drugs []struct {
		Name   string `json:"name"`
		Dosage string `json:"dosage"`
	} `json:"drugs"`

	Labs  []string `json:"labs"`
	Notes []string `json:"notes"`
}

func ParseWithAI(input string) (AIResponse, error) {

	apiKey := os.Getenv("OPENAI_API_KEY")

	payload := map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "Extract drugs, dosages, lab tests, and notes from medical text. Return JSON like {drugs:[], labs:[], notes:[]}",
			},
			{
				"role":    "user",
				"content": input,
			},
		},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return AIResponse{}, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	content := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	var parsed AIResponse
	json.Unmarshal([]byte(content), &parsed)

	return parsed, nil
}
