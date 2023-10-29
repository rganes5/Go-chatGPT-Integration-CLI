package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	clientgpt "github.com/rganes5/Go-chatGPT/Go-chatGPT-Intergration-CLI/clientgpt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type NullWriter int

func (NullWriter) Write([]byte) (int, error) {
	return 0, nil
}

func main() {
	log.SetOutput(new(NullWriter))
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		panic("Missing API KEY")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	//Using cobra for the command line arguments
	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console.",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Ask Anything: (type 'quit' to end):")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true

				default:
					clientgpt.GetResponse(client, ctx, question)
				}

			}
		},
	}
	rootCmd.Execute()
}
