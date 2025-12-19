package kgsservice

import (
	"bytes"
	"encoding/json"
)

func (kgs *KGSService) sendPostRequest(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := kgs.client.Post(kgs.url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}