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
4. 支持swagger文档 和 ApiFox 文档

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
	"os"
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"github.com/ConnectAI-E/go-minimax/minimax"
)

//init client

func main() {

	ctx := context.Background()
	client, _ := minimax.New(
		minimax.WithApiToken(os.Getenv("TEST_MINIMAX_API_TOKEN")),
		minimax.WithGroupId(os.Getenv("TEST_MINIMAX_GROUP_ID")),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				SenderType: "USER",
				Text:       "hi~",
			},
		},
		Model:       "abab5-chat",
		Temperature: 0.7,
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Choices) // output: 你好！有什么我可以帮助你的吗？
}

```
