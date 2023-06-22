package claude

import (
	"context"
	"fmt"
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Completion(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
	res, err := cli.ChatCompletions(context.Background(),
		&textv1.ChatCompletionsRequest{
			Prompt:            "Human: hello; Assistant:",
			Model:             "claude-1-100k",
			Temperature:       0.7,
			MaxTokensToSample: 500,
		})
	fmt.Printf("res: %+v\n", res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestClient_Completion_Message(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
	res, err := cli.ChatCompletions(context.Background(),
		&textv1.ChatCompletionsRequest{
			Messages: []*textv1.Message{
				{
					Role: "Human", Content: "hello",
				},
			},
			Model:             "claude-1-100k",
			Temperature:       0.7,
			MaxTokensToSample: 500,
		})
	fmt.Printf("res: %+v\n", res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestClient_Completions_Fail(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
	res, err := cli.ChatCompletions(context.Background(),
		&textv1.ChatCompletionsRequest{
			Prompt:            "hello",
			Model:             "claude-1-100k",
			Temperature:       0.7,
			MaxTokensToSample: 500,
		})
	fmt.Printf("res: %+v\n", res)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}
}

func TestClient_CompletionStream(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
	stream, err := cli.ChatCompletionStream(context.Background(),
		&textv1.ChatCompletionsRequest{
			Prompt:            "Human: hello; Assistant:",
			Model:             "claude-1-100k",
			Temperature:       0.7,
			MaxTokensToSample: 500,
		})
	assert.Nil(t, err)
	assert.NotNil(t, stream)
	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("res: %+v\n", res)
	}
}
