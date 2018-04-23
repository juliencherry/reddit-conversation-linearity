package main

import (
	// "env"
	"fmt"
	// "reddit"
	. "linearity"
)

func main() {

	/*envManager := env.Manager{}
	env, err := envManager.GetRedditEnv()
	if (err != nil) {
		fmt.Printf("Couldn’t get Reddit environment: %s\n", err.Error()); return
	}

	reddit := reddit.APIClient{}
	err = reddit.RetrieveAuthToken(env.Username, env.Password, env.ClientID, env.ClientSecret)
	if (err != nil) {
		fmt.Printf("Couldn’t retrieve authorization token: %s\n\n", err.Error()); return
	}*/

	// reddit := reddit.APIClient{}
	// reddit.GetInfo()

	testConversation := CommentWithChildren{
		Children: []Conversation{
			CommentWithChildren{
				Children: []Conversation{
					FinalComment{},
					FinalComment{},
				},
			},
			CommentWithChildren{
				Children: []Conversation{
					FinalComment{},
					FinalComment{},
				},
			},
		},
	}

	linearityCalc := Calculator{}
	fmt.Printf("Linearity: %f\n", linearityCalc.Linearity(testConversation))
}
