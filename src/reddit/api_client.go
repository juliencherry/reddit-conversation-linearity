package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type APIClient struct {
	authToken string
}

type AccessToken struct {
	Name string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	Scope string
}

type ListingJSON struct {
	Data {
		Children {

		}`json:"children"`
	} `json:"data"`
}

type ThingJSON struct {
	Kind string
	Data interface{}
}

func (t *ThingJSON) UnmarshalJSON(b []byte) {

	type IntermediateThingJSON struct {
		Kind string
		Data json.RawMessage
	}

	var data IntermediateThingJSON

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	*t.Data.Kind = data.Kind

	switch data.Kind {
	default:
		err = json.Unmarshal(data.Data, &t.Data)
	case "more":
		err = json.Unmarshal(data.Data, &t.Data)
	}

	return nil
}

type CommentData struct {
	Author string
	Body string `json:"body"`
	BodyHTML string `json:"body_html"`
	Replies ListingJSON `json:"replies"`
}

type MoreData struct {
	ParentID string
	Children []string
}

func (a *APIClient) RetrieveAuthToken(username string, password string, clientID string, clientSecret string) (error) {
	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("user", username)
	form.Add("password", password)

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(form.Encode()))
	if (err != nil) {
		return err
	}

	req.PostForm = form
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("User-Agent", "Conversation Linearity Calculator")

	client := http.Client{}
	resp, err := client.Do(req)
	if (err != nil) {
		return err
	}

	if (resp.StatusCode != 200) {
		fmt.Printf("%+v", resp)
		return fmt.Errorf("server responded with status %s", resp.Status)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var token AccessToken
	err = json.Unmarshal(body, &token)
	if (err != nil) {
		return err
	}

	a.authToken = token.Name
	return nil
}

func (a APIClient) GetInfo() {
	req, _ := http.NewRequest("GET", "https://www.reddit.com/r/mildlyinteresting/comments/8e2by0.json", strings.NewReader(""))

	req.Header.Add("User-Agent", "Conversation Linearity Calculator")
	// req.Header.Add("Authorization", fmt.Sprintf("bearer %s", a.authToken))

	client := http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var dataAsJSON []json.RawMessage
	_ = json.Unmarshal(body, &dataAsJSON)

	justComments := dataAsJSON[1]

	var listing ListingJSON
	_ = json.Unmarshal(justComments, &listing)

	firstComment := listing.Data.Children[0]
	for _, comment := range firstComment.Data.Replies.Data.Children {
		fmt.Printf("%s says: %s\n", comment.Data.Author, comment.Data.Body)
	}
}
