<p align='center'>
    <img src='https://github.com/ConnectAI-E/Go-Claude/assets/50035229/88d59926-8111-46b5-9a26-e86321e67bca' alt='' width='1300'/>
</p>


## CLAUDE-SDK
> Anthropic Claude API

[😎点击查看官方文档](https://docs.anthropic.com/claude/docs)

### 功能及特点

1. 全接口字段注释
2. Chatcompletion 文本对话接口
3. 无缝对接官方文档：单轮问答、历史记忆问答、流返回
4. 修改Prompt传参为数组模式，更方便的兼容openai的接口
5. 支持swagger文档 和 ApiFox 文档


### Swagger 文档
- [打开在线Swagger编辑器](https://editor.swagger.io/)
- 导入[Swagger Api 文档](./output/apis.swagger.yaml)
  <img width="120" alt="image" src="https://github.com/imroc/req/assets/50035229/efdbf241-0bb7-43a8-8389-749724f0e232">

### 使用方法

1. 访问 [Anthropic 官网](https://www.anthropic.com/index/introducing-claude)，申请[Claude API](https://www.anthropic.com/product) 权限。
2. 登录 [Claude Console](https://console.anthropic.com/login)，获取 [Api-Key](https://console.anthropic.com/account/keys)。
3. 参考下面示例使用

### 示例

```go
package main

import (
	"context"
	"fmt"
	"github.com/ConnectAI-E/go-claude/claude"
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"os"
)

//init client

func main() {

	ctx := context.Background()
	client, _ := claude.New(
		claude.WithApiToken(os.Getenv("TEST_API_TOKEN")),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				Role:    "Human",
				Content: "hi~",
			},
		},
		Model:             "claude-1-100k",
		Temperature:       0.7,
		MaxTokensToSample: 500,
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Completion) // output: Hello

}

```


### 快速上手:

<details>
<summary>Claude Completion</summary>

```go
package main

import (
	"context"
	"fmt"
	"github.com/ConnectAI-E/go-claude/claude"
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"os"
)



func main() {
	ctx := context.Background()

	client, _ := claude.New(
		claude.WithApiToken(os.Getenv("TEST_API_TOKEN")),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				Role:    "Human",
				Content: "hi~",
			},
		},
		Model:             "claude-1-100k",
		Temperature:       0.7,
		MaxTokensToSample: 500,
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Completion) // output: Hello

}
```
</details>


<details>
<summary>Claude stream completion</summary>

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ConnectAI-E/go-claude/claude"
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"io"
	"os"
)



func main() {
	ctx := context.Background()
	
	//init client
	client, _ := claude.New(
		claude.WithApiToken(os.Getenv("TEST_API_TOKEN")),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				Role:    "Human",
				Content: "hi~",
			},
		},
		Model:             "claude-1-100k",
		Temperature:       0.7,
		MaxTokensToSample: 500,
	}

	stream, _ := client.ChatCompletionStream(ctx, req)
	defer stream.CloseSend()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf(response.Completion + "\n") 
	}
}


```
</details>


<details>
<summary>Claude history stream completion</summary>

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ConnectAI-E/go-claude/claude"
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"io"
	"os"
)



func main() {
	ctx := context.Background()

	//init client
	client, _ := claude.New(
		claude.WithApiToken(os.Getenv("TEST_API_TOKEN")),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				Role:    "Human",
				Content: "hi~",
			},
			{
				Role:    "Assistant",
				Content: "Hello",
            },
			{
				Role:    "Human",
				Content: "How are you?",
            },
		},
		Model:             "claude-1-100k",
		Temperature:       0.7,
		MaxTokensToSample: 500,
	}

	stream, _ := client.ChatCompletionStream(ctx, req)
	defer stream.CloseSend()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf(response.Completion + "\n")
	}
}

```
</details>


