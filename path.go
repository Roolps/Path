package path

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type APIClient struct {
	Username string
	Password string

	baseURL string
	token   token
}

type token struct {
	Token               string `json:"access_token"`
	Flows               bool   `json:"flows"`
	Metrics             bool   `json:"metrics"`
	HasFirewall         bool   `json:"has_firewall"`
	HasDomains          bool   `json:"has_domains"`
	HasWAF              bool   `json:"has_waf"`
	CanAnnouncePrefixes bool   `json:"can_announce_prefixes"`
}

// set up api client
func (c *APIClient) New() error {
	c.baseURL = "https://api.path.net"
	return c.newToken()
}

// generate new access token
func (c *APIClient) newToken() error {
	query := url.Values{
		"username": []string{c.Username},
		"password": []string{c.Password},
	}
	_, err := c.requestHandler("/token", http.MethodPost, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(query.Encode()))
	return err
}

// makes http request to the api using api key secret as authorization
func (c *APIClient) requestHandler(endpoint string, method string, headers map[string]string, body []byte) ([]byte, error) {
	client := &http.Client{}
	reader := bytes.NewReader(body)
	req, err := http.NewRequest(method, c.baseURL+endpoint, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.token.Token)
	if headers["Content-Type"] == "" {
		req.Header.Add("Content-Type", "Application/Json")
	} else {
		req.Header.Add("Content-Type", headers["Content-Type"])
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(raw))
	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		type detail struct {
			Detail json.RawMessage `json:"detail"`
		}
		d := detail{}
		json.Unmarshal(raw, &d)
		if res.StatusCode == 422 {
			d := extractError(d.Detail)
			return nil, fmt.Errorf("ERROR %v: %v", res.StatusCode, d)
		}
		return nil, fmt.Errorf("ERROR %v: %v", res.StatusCode, string(d.Detail))
	}

	return raw, nil
}

// extract first of error array
func extractError(raw []byte) string {
	type detail struct {
		Loc  []string `json:"loc"`
		Msg  string   `json:"msg"`
		Type string   `json:"type"`
	}
	d := []detail{}
	json.Unmarshal(raw, &d)
	return fmt.Sprintf("%v: %v", d[0].Msg, strings.Join(d[0].Loc, ": "))
}
