package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type APIClient struct {}

type AccessToken struct {
	Name string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	Scope string
}

func (APIClient) GetAuthToken(username string, password string, clientID string, clientSecret string) (AccessToken, error) {
	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("user", username)
	form.Add("password", password)

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(form.Encode()))
	if (err != nil) {
		return AccessToken{}, err
	}

	req.PostForm = form
	req.SetBasicAuth(clientID, clientSecret)

	client := http.Client{}
	resp, err := client.Do(req)
	if (err != nil) {
		return AccessToken{}, err
	}

	if (resp.StatusCode != 200) {
		return AccessToken{}, fmt.Errorf("server responded with status %s", resp.Status)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var token AccessToken
	err = json.Unmarshal(body, &token)
	if (err != nil) {
		return AccessToken{}, err
	}

	return token, nil;
}
