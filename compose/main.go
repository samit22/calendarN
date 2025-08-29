package main

import (
	"fmt"
	"os"
)

var (
	consumerKey       = os.Getenv("CONSUMER_KEY")
	consumerSecret    = os.Getenv("CONSUMER_SECRET")
	accessToken       = os.Getenv("ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
)

func main() {
	if err := validateEnv(); err != nil {
		fmt.Printf("environment validation failed: %v", err)
		os.Exit(1)
	}
	tweetID, err := CreateTweet(getToday())
	if err != nil {
		fmt.Printf("failed to tweet %v", err)
		os.Exit(1)
	}
	fmt.Printf("created tweet: %s\n", tweetID)
}

func validateEnv() error {
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		return fmt.Errorf("one or more required environment variables are missing")
	}

	return nil
}
