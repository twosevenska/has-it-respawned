package steampowered

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	resty "gopkg.in/resty.v0"
)

const steamURL = "http://api.steampowered.com"

// ResponseError represents a plaza error
type ResponseError struct {
	StatusCode int             `json:"-"`
	Message    string          `json:"message"`
	ErrorCode  int             `json:"error_code"`
	Debug      json.RawMessage `json:"debug,omitempty"`
	Path       json.RawMessage `json:"path,omitempty"`
}

// Error implements the error interface
func (e ResponseError) Error() string {
	return e.Message
}

type Client struct {
	client *resty.Client
	URL    string
	Key    string
}

// New creates and returns a new Mailchimp API client
func New(steamKey string) Client {
	return Client{
		client: resty.New().
			SetLogger(ioutil.Discard).
			SetTimeout(5*time.Second).
			SetHeader("Content-Type", "application/json"),
		URL: steamURL,
		Key: steamKey,
	}
}

// GetMembers reads a list of members from Mailchimp
func (c *Client) GetGames(userID string) (GamesList, error) {
	var cResp Response
	var cErr ResponseError
	u := fmt.Sprintf("%s/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=1&include_played_free_games=1", c.URL, c.Key, userID)
	resp, err := c.client.R().
		SetResult(&cResp).
		SetError(&cErr).
		Get(u)

	if err != nil {
		return GamesList{}, err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		cErr.StatusCode = resp.StatusCode()
		err = cErr
		return GamesList{}, err
	}

	return cResp.GL, nil
}
