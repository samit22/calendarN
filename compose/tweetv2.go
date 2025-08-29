package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// TweetPayload represents the JSON payload for posting a tweet
type TweetPayload struct {
	Text string `json:"text"`
}

// TweetResponse represents the API response
type TweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

func CreateTweet(tweetText string) (string, error) {
	if len(tweetText) > 280 {
		return "", fmt.Errorf("tweet exceeds 280 characters")
	}
	apiURL := "https://api.twitter.com/2/tweets"
	payload := TweetPayload{Text: tweetText}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshaling payload: %w", err)
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("creating http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	oauthParams := generateOAuthParams(consumerKey, accessToken)
	signature := generateOAuthSignature(req, oauthParams, consumerSecret, accessTokenSecret)
	oauthHeader := buildOAuthHeader(oauthParams, signature)
	req.Header.Set("Authorization", oauthHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("sending http request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("tweet post status: %s", resp.Status)
	}
	var tweetResp TweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweetResp); err != nil {
		return "", fmt.Errorf("decoding tweet response: %v", err)
	}
	return tweetResp.Data.ID, nil
}

func generateOAuthParams(consumerKey, accessToken string) map[string]string {
	return map[string]string{
		"oauth_consumer_key":     consumerKey,
		"oauth_token":            accessToken,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        fmt.Sprintf("%d", time.Now().Unix()),
		"oauth_nonce":            fmt.Sprintf("%d", time.Now().UnixNano()),
		"oauth_version":          "1.0",
	}
}

func generateOAuthSignature(req *http.Request, oauthParams map[string]string, consumerSecret, tokenSecret string) string {
	params := make(url.Values)
	for k, v := range oauthParams {
		params.Add(k, v)
	}
	var paramPairs []string
	for k, v := range params {
		for _, val := range v {
			paramPairs = append(paramPairs, fmt.Sprintf("%s=%s", percentEncode(k), percentEncode(val)))
		}
	}
	sort.Strings(paramPairs)
	paramString := strings.Join(paramPairs, "&")

	baseString := fmt.Sprintf(
		"%s&%s&%s",
		req.Method,
		percentEncode(req.URL.String()),
		percentEncode(paramString),
	)

	signingKey := fmt.Sprintf("%s&%s", percentEncode(consumerSecret), percentEncode(tokenSecret))

	h := hmac.New(sha1.New, []byte(signingKey))
	h.Write([]byte(baseString))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}

func buildOAuthHeader(oauthParams map[string]string, signature string) string {
	var headerParts []string
	for k, v := range oauthParams {
		headerParts = append(headerParts, fmt.Sprintf(`%s="%s"`, k, percentEncode(v)))
	}
	headerParts = append(headerParts, fmt.Sprintf(`oauth_signature="%s"`, percentEncode(signature)))
	return "OAuth " + strings.Join(headerParts, ", ")
}

func percentEncode(s string) string {
	encoded := url.QueryEscape(s)
	encoded = strings.ReplaceAll(encoded, "+", "%20")
	return encoded
}
