package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	ctx := context.Background()
	os.Setenv("OPENAI_API_KEY", "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")
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	prompt := "단미가 현태에게 언제 오냐고 물어보는데 어떻게 대답하는게 좋을까?"
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
