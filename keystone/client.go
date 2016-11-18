package keystone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const X_SUBJECT_TOKEN_HEADER = "X-Subject-Token"

type request struct {
	URL          string
	Method       string
	Body         []byte
	OkStatusCode int
}

type response struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
}

type Client struct {
	URL string
}

func NewClient(url string) (*Client, error) {
	if url == "" {
		return nil, fmt.Errorf("missing URL")
	}
	return &Client{URL: url}, nil
}

func (c *Client) doRequest(r request) (response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return response{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response{}, err
	}

	if resp.StatusCode != r.OkStatusCode {
		return response{}, fmt.Errorf("Error: %s details: %s\n", resp.Status, body)
	}

	return response{
		Body:       body,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header}, nil
}

func (c *Client) Tokens(auth Auth) (string, error) {
	jsonStr, err := json.Marshal(SingleAuth{Auth: auth})
	if err != nil {
		return "", fmt.Errorf("invalid auth request: ", err)
	}

	resp, err := c.doRequest(request{
		URL:          fmt.Sprintf("%s/v3/auth/tokens", c.URL),
		Method:       http.MethodPost,
		Body:         jsonStr,
		OkStatusCode: http.StatusCreated,
	})

	if err != nil {
		return "", err
	}

	// note: not unmarshalling response body right now
	// since dont need anything from it yet
	token := resp.Headers.Get(X_SUBJECT_TOKEN_HEADER)
	if token == "" {
		return "", errors.New("No token found in response")
	}

	return token, nil
}
