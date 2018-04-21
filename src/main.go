package main

import (
	"env"
	"fmt"
	"reddit"
)

func main() {
	envManager := env.Manager{}
	env, err := envManager.GetRedditEnv()
	if (err != nil) {
		fmt.Printf("Couldn’t get Reddit environment: %s\n", err.Error()); return
	}

	reddit := reddit.APIClient{}
	authToken, err := reddit.GetAuthToken(env.Username, env.Password, env.ClientID, env.ClientSecret)
	if (err != nil) {
		fmt.Printf("Couldn’t get authorization token: %s\n", err.Error()); return
	}

	fmt.Printf("Authorization token: %+v\n", authToken)
}
