package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	// Load Env & API Key from .env
	godotenv.Load()
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	c := gogpt.NewClient(apiKey)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Davinci,
		MaxTokens: 20,
		Prompt:    "I can't stand homework",
		//Stream:    true,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)

	/*
		stream, err := c.CreateCompletionStream(ctx, req)
		if err != nil {
			return
		}
		defer stream.Close()

		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("Stream finished")
				return
			}

			if err != nil {
				fmt.Printf("Stream error: %v\n", err)
				return
			}

			//fmt.Printf("Stream response: %v\n", response)
			fmt.Printf(response.Choices[0].Text)

		}

	*/
}
