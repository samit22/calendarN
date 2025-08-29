package main

import (
	"os"
	"testing"
)

func setEnvVars(vars map[string]string) func() {
	originals := map[string]string{
		"CONSUMER_KEY":        os.Getenv("CONSUMER_KEY"),
		"CONSUMER_SECRET":     os.Getenv("CONSUMER_SECRET"),
		"ACCESS_TOKEN":        os.Getenv("ACCESS_TOKEN"),
		"ACCESS_TOKEN_SECRET": os.Getenv("ACCESS_TOKEN_SECRET"),
	}
	for k, v := range vars {
		os.Setenv(k, v)
	}
	return func() {
		for k, v := range originals {
			os.Setenv(k, v)
		}
	}
}

func TestValidateEnv_AllSet(t *testing.T) {
	cleanup := setEnvVars(map[string]string{
		"CONSUMER_KEY":        "key",
		"CONSUMER_SECRET":     "secret",
		"ACCESS_TOKEN":        "token",
		"ACCESS_TOKEN_SECRET": "tokensecret",
	})
	defer cleanup()

	// reload globals
	consumerKey = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")
	accessToken = os.Getenv("ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")

	if err := validateEnv(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestValidateEnv_MissingVars(t *testing.T) {
	tests := []struct {
		name string
		envs map[string]string
	}{
		{"MissingConsumerKey", map[string]string{
			"CONSUMER_KEY":        "",
			"CONSUMER_SECRET":     "secret",
			"ACCESS_TOKEN":        "token",
			"ACCESS_TOKEN_SECRET": "tokensecret",
		}},
		{"MissingConsumerSecret", map[string]string{
			"CONSUMER_KEY":        "key",
			"CONSUMER_SECRET":     "",
			"ACCESS_TOKEN":        "token",
			"ACCESS_TOKEN_SECRET": "tokensecret",
		}},
		{"MissingAccessToken", map[string]string{
			"CONSUMER_KEY":        "key",
			"CONSUMER_SECRET":     "secret",
			"ACCESS_TOKEN":        "",
			"ACCESS_TOKEN_SECRET": "tokensecret",
		}},
		{"MissingAccessTokenSecret", map[string]string{
			"CONSUMER_KEY":        "key",
			"CONSUMER_SECRET":     "secret",
			"ACCESS_TOKEN":        "token",
			"ACCESS_TOKEN_SECRET": "",
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cleanup := setEnvVars(tc.envs)
			defer cleanup()
			consumerKey = os.Getenv("CONSUMER_KEY")
			consumerSecret = os.Getenv("CONSUMER_SECRET")
			accessToken = os.Getenv("ACCESS_TOKEN")
			accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")

			if err := validateEnv(); err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
func TestExit_CallsOsExit(t *testing.T) {
	// Save original osExit and restore after test
	origExit := osExit
	defer func() { osExit = origExit }()

	called := false
	var code int
	osExit = func(c int) {
		called = true
		code = c
		panic("os.Exit called")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic from os.Exit, got none")
		}
		if !called {
			t.Errorf("expected osExit to be called")
		}
		if code != 1 {
			t.Errorf("expected exit code 1, got %d", code)
		}
	}()

	exit("test message %d", 42)
}
