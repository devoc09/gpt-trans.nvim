package translator

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/neovim/go-client/nvim"
	openai "github.com/sashabaranov/go-openai"
)

type (
	Translator interface {
		TranslatePos(ctx context.Context, opt Option) ([]string, error)
		TranslateWord(ctx context.Context, opt Option) ([]string, error)
	}

	translator struct {
		vim *nvim.Nvim
	}

	Option struct {
		Source          string
		Target          string
		Cutset          []string
		CredentialsFile string
	}
)

func New(vim *nvim.Nvim) Translator {
	return &translator{
		vim: vim,
	}
}

func (t *translator) TranslatePos(ctx context.Context, opt Option) ([]string, error) {
	var startPos []int
	if err := t.vim.Eval(`getpos("'<")`, &startPos); err != nil {
		return nil, err
	}
	var endPos []int
	if err := t.vim.Eval(`getpos("'>")`, &endPos); err != nil {
		return nil, err
	}

	if startPos[1] == 0 && startPos[2] == 0 && endPos[1] == 0 && endPos[2] == 0 {
		return nil, nil
	}

	var text string
	if startPos[1] == endPos[1] {
		b, err := t.vim.CurrentLine()
		if err != nil {
			return nil, err
		}
		text = string(b)
		if endPos[2] > len(text) {
			endPos[2] = len(text)
		}
		text = text[startPos[2]-1 : endPos[2]]
	} else {
		b, err := t.vim.CurrentBuffer()
		if err != nil {
			return nil, err
		}
		bytes, err := t.vim.BufferLines(b, startPos[1]-1, endPos[1], true)
		if err != nil {
			return nil, err
		}

		lines := make([]string, len(bytes))
		for i, b := range bytes {
			line := string(b)
			if i == 0 {
				line = line[startPos[2]-1:]
			} else if i == len(lines)-1 {
				if endPos[2] > len(line) {
					endPos[2] = len(line)
				}
				line = line[:endPos[2]]
			}

			line = strings.TrimSpace(line)
			for _, cutset := range opt.Cutset {
				line = strings.TrimLeft(line, cutset)
			}
			lines[i] = strings.TrimSpace(line)
		}
		text = strings.Join(lines, " ")
	}

	return t.translate(ctx, text, opt)
}

func (t *translator) TranslateWord(ctx context.Context, opt Option) ([]string, error) {
	var text string
	if err := t.vim.Eval("expand('<cword>')", &text); err != nil {
		return nil, err
	}
	return t.translate(ctx, text, opt)
}

func (t *translator) translate(ctx context.Context, text string, opt Option) ([]string, error) {
	if text == "" {
		return nil, nil
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a translator English into Japanese.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}


	ss := strings.Split(resp.Choices[0].Message.Content, "\n")
	for i := range ss {
		ss[i] = strings.TrimSpace(ss[i])
	}
	return ss, nil
}
