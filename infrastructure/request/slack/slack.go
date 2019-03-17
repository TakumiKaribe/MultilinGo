package slack

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/TakumiKaribe/multilingo/entity/slack"
	infraRequest "github.com/TakumiKaribe/multilingo/infrastructure/request"
	interfacesRequest "github.com/TakumiKaribe/multilingo/usecase/interfaces/request"
	requestSlack "github.com/TakumiKaribe/multilingo/usecase/interfaces/request/slack"

	"github.com/pkg/errors"
)

const (
	postMessage = "/api/chat.postMessage"
)

// Client -
type client struct {
	url                *url.URL
	requester          *infraRequest.Reqeuster
	botUserAccessToken string
}

// NewClient Constructor -
func NewClient(host string, botUserAccessToken string) (requestSlack.Client, error) {
	if len(botUserAccessToken) == 0 {
		// TODO: use multilingo error
		return nil, errors.New("missing  botUserAccessToken")
	}

	client := client{botUserAccessToken: botUserAccessToken, requester: infraRequest.NewRequester()}
	client.url, _ = url.Parse(host + postMessage)

	return &client, nil
}

// Notification -
func (c *client) Notify(requestBody *slack.RequestBody) error {
	bodyByte, _ := json.Marshal(requestBody)
	bodyReader := bytes.NewReader(bodyByte)

	header := map[string]string{}
	header["Authorization"] = "Bearer " + c.botUserAccessToken
	header["Content-Type"] = "application/json; charset=UTF-8"
	body, err := c.requester.Request(interfacesRequest.Post, c.url.String(), bodyReader, header)
	if err != nil {
		return err
	}

	defer body.Close()
	return nil
}
