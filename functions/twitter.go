package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// EnvLoad load twitter Key & Token
func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//TwitterClient for creating Twitter client
func TwitterClient() *twitter.Client {
	EnvLoad()
	var (
		consumerKey    = os.Getenv("consumerKey")
		consumerSecret = os.Getenv("consumerSecret")
		accessToken    = os.Getenv("accessToken")
		accessSecret   = os.Getenv("accessSecret")
	)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client
}