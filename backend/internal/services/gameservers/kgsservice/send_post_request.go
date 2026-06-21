package kgsservice

import (
	"bytes"
	"encoding/json"
	"log"
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
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("close response body error: %v", err)
		}
	}()

	return nil
}
