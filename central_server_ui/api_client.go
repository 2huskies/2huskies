package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/2huskies/structs"
)

type ApiClient struct {
	URL *url.URL
}

func newApiClient(rawurl string) (*ApiClient, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse api URL '%s': %s", rawurl, err)
	}
	u.Path = ""
	u.RawQuery = ""
	api = &ApiClient{
		URL: u,
	}
	return api, nil
}

//return nil, nil means permission denied
func (a *ApiClient) verify_user(username string, password string) (*structs.UserCheckResult, error) {
	uc := structs.UserCheck{username, password}
	buf, err := json.Marshal(uc)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(buf)
	req, err := a.NewRequest("POST", "verify_user", reader)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 401 {
		return nil, nil
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid response status, want: 200, got: %d", resp.StatusCode)
	}
	dec := json.NewDecoder(resp.Body)
	result := &structs.UserCheckResult{}
	err = dec.Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *ApiClient) NewRequest(method string, path string, body io.Reader) (*http.Request, error) {
	u := &url.URL{
		Scheme: a.URL.Scheme,
		Host:   a.URL.Host,
		Path:   path,
	}

	return http.NewRequest(method, u.String(), body)
}
