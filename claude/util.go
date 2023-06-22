package claude

import (
	"fmt"
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"strings"
)

func convertArrayToString(messages []*textv1.Message) string {
	var formattedMessages []string
	for _, msg := range messages {
		formattedMessages = append(formattedMessages, fmt.Sprintf("%s: %s", msg.Role, msg.Content))
	}
	return strings.Join(formattedMessages, "\\n")
}

func FormatMsg(messages []*textv1.Message) string {
	var formattedMessages []string
	for _, msg := range messages {
		formattedMessages = append(formattedMessages, fmt.Sprintf("%s: %s", msg.Role, msg.Content))
	}
	joinStr := strings.Join(formattedMessages, " ")
	lastPrompt := "\\nAssistant:"
	return fmt.Sprintf("%s %s", joinStr, lastPrompt)
}
