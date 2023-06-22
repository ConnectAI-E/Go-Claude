package claude

import (
	"context"
	"errors"
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"github.com/ConnectAI-E/go-claude/internal"
	"google.golang.org/grpc"
	"io"
)

var _ textv1.MinimaxServiceClient = new(Client)

func (cli *Client) ChatCompletions(ctx context.Context, in *textv1.ChatCompletionsRequest, opts ...grpc.CallOption) (*textv1.ChatCompletionsResponse, error) {
	res := new(struct {
		textv1.ChatCompletionsResponse
	})
	in.Stream = false
	if in.Prompt == "" {
		in.Prompt = FormatMsg(in.Messages)
	}
	resp, err := cli.client.R().
		SetBody(in).
		SetSuccessResult(res).
		SetErrorResult(res).
		Post("/v1/complete")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(res.Error.Message)
	}
	return &res.ChatCompletionsResponse, err
}

func (cli *Client) ChatCompletionStream(ctx context.Context, in *textv1.ChatCompletionsRequest, opts ...grpc.CallOption) (textv1.MinimaxService_ChatCompletionStreamClient, error) {

	in.Stream = true
	if in.Prompt == "" {
		in.Prompt = FormatMsg(in.Messages)
	}
	resp, err := cli.client.R().
		DisableAutoReadResponse().
		SetBody(in).
		Post("/v1/complete")

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	return internal.NewStreamReader[*textv1.ChatCompletionsResponse](resp.Body), err
}

// [
//{role:Human,content:"Hello, how are you?"},
//{role:Assistant,content:"I'm fine, thanks. How are you?"},
//] -> "Human: Hello, how are you? Assistant: I'm fine, thanks. How are you?"
