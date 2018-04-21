package env

import (
	"fmt"
	"os"
)

type Manager struct {}

type RedditEnv struct {
	Username string
	Password string
	ClientID string
	ClientSecret string
}

func (Manager) GetRedditEnv() (RedditEnv, error) {
	env := RedditEnv{}
	err := fmt.Errorf("")

	env.Username, err = getEnv("REDDIT_USERNAME")
	if (err != nil) {
		return env, err
	}

	env.Password, err = getEnv("REDDIT_PASSWORD")
	if (err != nil) {
		return env, err
	}

	env.ClientID, err = getEnv("REDDIT_CLIENT_ID")
	if (err != nil) {
		return env, err
	}

	env.ClientSecret, err = getEnv("REDDIT_CLIENT_SECRET")
	if (err != nil) {
		return env, err
	}

	return env, nil
}

func getEnv(key string) (string, error) {
	val := os.Getenv(key)
	if (val == "") {
		return "", fmt.Errorf("environment variable “%s” unset", key)
	}

	return val, nil
}
