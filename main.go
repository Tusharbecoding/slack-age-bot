package main

import (
	"fmt"
	"context"
	"log"
	"os"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent)  {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}
}

func main()  {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5529699535137-5510423351942-ZjdDV2biRRFeQEcWljLKCcMc")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05FKLQ1GM7-5547745943014-c45fda0ee40a1517a69913d16138dcde17634dc8c846c6b565b8b40bb060d18e")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvent())

	bot.Command("My yob is <year>", &slacker.CommandDefinition) {
		Description: "yob calculator",
		Example: "My yob is 2002",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}