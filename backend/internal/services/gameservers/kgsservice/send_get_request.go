package kgsservice

import (
	"io"
	"log"
)

func (kgs *KGSService) sendGetRequest() (string, error) {

	resp, err := kgs.client.Get(kgs.url)
	if err != nil {
		return "", err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("close response body error: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
