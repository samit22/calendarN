package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// TwitterTokenResponse is the response structure from Twitter's OAuth2 token endpoint
type TwitterTokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func getAccessTokenOauth2() (string, error) {
	// Encode API key and secret
	credentials := consumerKey + ":" + consumerSecret
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))

	// Prepare the request
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}
	req.Header.Add("Authorization", "Basic "+encodedCredentials)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	var tokenResponse TwitterTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return "", err
	}

	fmt.Println("Access Token:", tokenResponse.AccessToken)

	return tokenResponse.AccessToken, nil
}

func sendTweet(accessToken, tweet string) error {
	d := map[string]string{
		"text": tweet,
	}
	byt, _ := json.Marshal(d)

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(byt))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Print the response (for demonstration purposes)
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return err
	}

	// Print the JSON response
	fmt.Println("Response:", result)
	return nil
}

// makeAuthenticatedRequest demonstrates using the access token to call a Twitter API endpoint
func makeAuthenticatedRequest(accessToken string) {
	// Example API call: Get account credentials
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/tweets", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response (for demonstration purposes)
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Print the JSON response
	fmt.Println("Response:", result)
}
