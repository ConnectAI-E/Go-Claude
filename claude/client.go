package claude

import (
	"github.com/imroc/req/v3"
	"os"
)

const BaseUrl = "https://api.anthropic.com"

type Client struct {
	client   *req.Client
	apiToken string
	groupId  string
}

func New(opts ...Option) (*Client, error) {
	baseClient := req.C().SetBaseURL(BaseUrl)
	if os.Getenv("APP_ENV") == "debug" {
		baseClient = baseClient.DevMode()
	}
	cli := &Client{
		client: baseClient,
	}
	//curl --location "https://api.minimax.chat/v1/text/chatcompletion?GroupId=${group_id}" \

	cli.client.OnBeforeRequest(func(client *req.Client,
		req *req.Request) error {
		if len(cli.apiToken) > 0 {
			req.SetHeader("x-api-key", cli.apiToken)
		}
		req.SetHeader("Content-Type", "application/json")
		return nil
	})
	for _, opt := range opts {
		err := opt(cli)
		if err != nil {
			return nil, err
		}
	}
	return cli, nil
}
