package slack

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Client -
type Client struct {
	URL                *url.URL
	HTTPClient         *http.Client
	botUserAccessToken string
}

// SlackRequestBody -
type SlackRequestBody struct {
	Token       string        `json:"token"`
	Channel     string        `json:"channel"`
	Attachments []*Attachment `json:"attachments"`
	UserName    string        `json:"username"`
}

// Attachment -
// https://api.slack.com/docs/message-attachments
type Attachment struct {
	Color     string `json:"color"` // good or warning or danger or colorcode
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
}

// NewClient Constructor -
func NewClient(host string, botUserAccessToken string) (*Client, error) {
	if len(botUserAccessToken) == 0 {
		return nil, errors.New("missing  botUserAccessToken")
	}

	client := Client{botUserAccessToken: botUserAccessToken,
		HTTPClient: &http.Client{Timeout: time.Duration(10) * time.Second}}
	client.URL, _ = url.Parse(host + "/api/chat.postMessage")

	return &client, nil
}

// Notification -
func (c *Client) Notification(body SlackRequestBody) (*http.Response, error) {
	bodyByte, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(bodyByte)

	req, err := c.newRequest("POST", bodyReader)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	return c.HTTPClient.Do(req)
}

// newRequest -
func (c *Client) newRequest(method string, body io.Reader) (*http.Request, error) {
	url := *c.URL

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.botUserAccessToken)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return req, nil
}