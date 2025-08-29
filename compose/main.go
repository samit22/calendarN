package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

var (
	consumerKey       = os.Getenv("CONSUMER_KEY")
	consumerSecret    = os.Getenv("CONSUMER_SECRET")
	accessToken       = os.Getenv("ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
)

func main() {
	ctx := context.Background()
	if err := validateEnv(); err != nil {
		exit("environment validation failed: %v", err)
	}
	tweetID, err := CreateTweet(ctx, &http.Client{}, getToday())
	if err != nil {
		exit("failed to create tweet: %v", err)
	}
	fmt.Printf("created tweet: %s\n", tweetID)
}

func validateEnv() error {
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		return fmt.Errorf("one or more required environment variables are missing")
	}

	return nil
}

var osExit = os.Exit

func exit(msg string, args ...any) {
	fmt.Printf(msg, args...)
	osExit(1)
}
