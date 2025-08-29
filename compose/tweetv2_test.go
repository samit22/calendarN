package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{Transport: fn}
}

func TestCreateTweet_Success(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(func(req *http.Request) *http.Response {
		if req.Method != "POST" {
			t.Errorf("expected POST, got %s", req.Method)
		}
		if req.URL.String() != "https://api.twitter.com/2/tweets" {
			t.Errorf("unexpected URL: %s", req.URL.String())
		}
		body, _ := io.ReadAll(req.Body)
		if !bytes.Contains(body, []byte(`"text":"hello world"`)) {
			t.Errorf("unexpected body: %s", string(body))
		}
		return &http.Response{
			StatusCode: http.StatusCreated,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"id": "12345",
					"text": "hello world"
				}
			}`)),
			Header: make(http.Header),
		}
	})

	id, err := CreateTweet(ctx, client, "hello world")
	assert.Nil(t, err)
	assert.Equal(t, "12345", id)
}

func TestCreateTweet_TooLong(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(func(req *http.Request) *http.Response {
		t.Fatal("should not send request for long tweet")
		return nil
	})
	longText := make([]byte, 281)
	for i := range longText {
		longText[i] = 'a'
	}
	_, err := CreateTweet(ctx, client, string(longText))
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "tweet exceeds 280 characters")
}

func TestCreateTweet_RequestError(t *testing.T) {
	client := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(bytes.NewBufferString(`{"error":"bad request"}`)),
			Header:     make(http.Header),
		}
	})
	_, err := CreateTweet(nil, client, "fail tweet")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "creating http request: net/http: nil Context")
}

func TestCreateTweet_HTTPError(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(bytes.NewBufferString(`{"error":"bad request"}`)),
			Header:     make(http.Header),
		}
	})
	_, err := CreateTweet(ctx, client, "fail tweet")
	if err == nil || !bytes.Contains([]byte(err.Error()), []byte("tweet post status")) {
		t.Errorf("expected tweet post status error, got %v", err)
	}
}

func TestCreateTweet_BadJSON(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(`not json`)),
			Header:     make(http.Header),
		}
	})
	_, err := CreateTweet(ctx, client, "bad json")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "decoding tweet response")
}

func TestCreateTweet_ReadBodyError(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       errorReader{},
			Header:     make(http.Header),
		}
	})
	_, err := CreateTweet(ctx, client, "test")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "reading response body")
}

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) { return 0, io.ErrUnexpectedEOF }
func (errorReader) Close() error             { return nil }

func TestBuildOAuthHeader(t *testing.T) {
	params := map[string]string{
		"oauth_consumer_key": "ckey",
		"oauth_token":        "atoken",
	}
	header := buildOAuthHeader(params, "sig")
	if !strings.HasPrefix(header, "OAuth ") {
		t.Errorf("header should start with OAuth: %s", header)
	}
	if !strings.Contains(header, `oauth_consumer_key="ckey"`) {
		t.Errorf("missing consumer key: %s", header)
	}
	if !strings.Contains(header, `oauth_token="atoken"`) {
		t.Errorf("missing token: %s", header)
	}
	if !strings.Contains(header, `oauth_signature="sig"`) {
		t.Errorf("missing signature: %s", header)
	}
}

func TestGenerateOAuthSignature_Deterministic(t *testing.T) {
	req, _ := http.NewRequest("POST", "https://api.twitter.com/2/tweets", nil)
	oauthParams := map[string]string{
		"oauth_consumer_key":     "ckey",
		"oauth_token":            "atoken",
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        "1234567890",
		"oauth_nonce":            "nonce",
		"oauth_version":          "1.0",
	}
	sig1 := generateOAuthSignature(req, oauthParams, "csecret", "tsecret")
	sig2 := generateOAuthSignature(req, oauthParams, "csecret", "tsecret")
	assert.Equal(t, sig1, sig2, "signatures should be equal")
}
func TestGenerateOAuthParams_TimestampNonce(t *testing.T) {
	params := generateOAuthParams("ckey", "atoken")
	// Check that timestamp is a valid integer and nonce is not empty
	if _, err := time.ParseDuration(params["oauth_timestamp"] + "s"); err != nil {
		t.Errorf("invalid timestamp: %s", params["oauth_timestamp"])
	}
	if params["oauth_nonce"] == "" {
		t.Error("nonce should not be empty")
	}
}

func TestPercentEncode(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"abc", "abc"},
		{"a b", "a%20b"},
		{"!*'();:@&=+$,/?#[]", "%21%2A%27%28%29%3B%3A%40%26%3D%2B%24%2C%2F%3F%23%5B%5D"},
	}
	for _, c := range cases {
		got := percentEncode(c.in)
		if got != c.want {
			t.Errorf("percentEncode(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGenerateOAuthParams(t *testing.T) {
	params := generateOAuthParams("ckey", "atoken")
	if params["oauth_consumer_key"] != "ckey" {
		t.Errorf("expected ckey, got %s", params["oauth_consumer_key"])
	}
	if params["oauth_token"] != "atoken" {
		t.Errorf("expected atoken, got %s", params["oauth_token"])
	}
	if params["oauth_signature_method"] != "HMAC-SHA1" {
		t.Errorf("unexpected signature method: %s", params["oauth_signature_method"])
	}
	if params["oauth_version"] != "1.0" {
		t.Errorf("unexpected version: %s", params["oauth_version"])
	}
	if params["oauth_timestamp"] == "" || params["oauth_nonce"] == "" {
		t.Error("timestamp or nonce not set")
	}
}
