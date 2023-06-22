package claude

import (
	textv1 "github.com/ConnectAI-E/go-claude/gen/go/claude/text/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatMsg(t *testing.T) {
	type args struct {
		messages []*textv1.Message
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				messages: []*textv1.Message{
					{
						Role:    "User",
						Content: "hello",
					},
				},
			},
			want: "User: hello \\nAssistant:",
		},
		{
			name: "test2",
			args: args{
				messages: []*textv1.Message{
					{
						Role:    "User",
						Content: "hello",
					},
					{
						Role:    "Assistant",
						Content: "hi",
					}, {
						Role:    "User",
						Content: "how are you?",
					},
				},
			},
			want: "User: hello Assistant: hi User: how are you" +
				"? \\nAssistant:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FormatMsg(tt.args.messages), "FormatMsg(%v)", tt.args.messages)
		})
	}
}
