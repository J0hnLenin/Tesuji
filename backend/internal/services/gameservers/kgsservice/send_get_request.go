package kgsservice

import "io"

func (kgs *KGSService) sendGetRequest() (string, error) {

	resp, err := kgs.client.Get(kgs.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}