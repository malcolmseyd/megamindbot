package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func generateMemeURL(text string) (string, error) {
	// make POST request to API with our meme text
	resp, err := http.Post(makePostURL(text), "", nil)
	if err != nil {
		return "", err
	}

	// read the response
	body := bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	// parse the response
	bodyJSON := responseJSON{}
	err = json.Unmarshal(body.Bytes(), &bodyJSON)
	if err != nil {
		return "", fmt.Errorf("failed to parse json: %v\n%v", err, body.String())
	}
	if !bodyJSON.Success {
		return "", fmt.Errorf("got success = false in response:\n%v", body.String())
	}

	// return the URL of the generated meme
	return bodyJSON.Data.URL, nil
}

// fill in the URL with the meme text + auth
func makePostURL(text string) string {
	return ("https://api.imgflip.com/caption_image?template_id=370867422&username=" +
		url.QueryEscape(username) +
		"&password=" +
		url.QueryEscape(password) +
		"&text0=" +
		url.QueryEscape(text) +
		"&text1")
}

type responseJSON struct {
	Success bool `json:"success"`
	Data    struct {
		URL     string `json:"url"`
		PageURL string `json:"page_url"`
	} `json:"data"`
}
