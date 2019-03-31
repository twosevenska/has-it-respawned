package steampowered

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	resty "gopkg.in/resty.v0"
)

// These values should be loaded from env or a secrets storage
const steamURL = "http://api.steampowered.com"
const steamKey = "REPLACE_WITH_API_KEY"

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
}

// New creates and returns a new Mailchimp API client
func New() Client {
	return Client{
		client: resty.New().
			SetLogger(ioutil.Discard).
			SetTimeout(5*time.Second).
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", fmt.Sprintf("apikey %s", steamKey)),
		URL: steamURL,
	}
}

// GetMembers reads a list of members from Mailchimp
func (c *Client) GetGames(userID string) (GamesList, error) {
	var cResp Response
	var cErr ResponseError
	u := fmt.Sprintf("%s/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=1&include_played_free_games=1", c.URL, steamKey, userID)
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
