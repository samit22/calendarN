package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dghubble/oauth1"
)

var (
	consumerKey      = os.Getenv("CONSUMER_KEY")
	consumerSecret   = os.Getenv("CONSUMER_SECRET")
	requestTokenURL  = "https://api.twitter.com/oauth/request_token?oauth_callback=oob&x_auth_access_type=write"
	baseAuthorizeURL = "https://api.twitter.com/oauth/authorize"
	accessTokenURL   = "https://api.twitter.com/oauth/access_token"
	tweetURL         = "https://api.twitter.com/2/tweets"
)

func main() {
	// token, err := getAccessTokenOauth2()
	// if err != nil {
	// 	fmt.Printf("Error getting access token: %v", err)
	// 	return
	// }
	// err = sendTweet(token, getToday())
	// if err != nil {
	// 	fmt.Printf("Error sending tweet: %v", err)
	// 	return
	// }
	// return
	// Step 1: Obtain a request token
	config := oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: requestTokenURL,
			AuthorizeURL:    baseAuthorizeURL,
			AccessTokenURL:  accessTokenURL,
		},
	}

	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		log.Fatalf("Error getting request token: %v", err)
	}
	fmt.Printf("Got OAuth token: %s\n", requestToken)

	// Step 2: Redirect user to Twitter for authorization
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		log.Fatalf("Error getting authorization URL: %v", err)
	}
	fmt.Printf("Please go here and authorize: %s\n", authorizationURL.String())

	// Step 3: Get the PIN from the user and exchange the request token for an access token
	fmt.Print("Paste the PIN here: ")
	var verifier string
	fmt.Scan(&verifier)

	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, verifier)
	if err != nil {
		log.Fatalf("Error getting access token: %v", err)
	}

	// Step 4: Use the access token to make a request to the Twitter API
	httpClient := config.Client(oauth1.NoContext, oauth1.NewToken(accessToken, accessSecret))

	payload := map[string]string{"text": getToday()}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshalling JSON payload: %v", err)
	}

	resp, err := httpClient.Post(tweetURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error making request to Twitter API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Request returned an error: %d %s", resp.StatusCode, body)
	}

	// Step 5: Parse and print the JSON response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON response: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(jsonResponse, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling JSON response: %v", err)
	}

	fmt.Printf("Response code: %d\n", resp.StatusCode)
	fmt.Println(string(prettyJSON))
}
