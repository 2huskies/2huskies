package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

type UserCheck struct {
	UserName string
	Password string
}

func (a *ApiClient) check_user(username string, password string) (bool, error) {
	uc := UserCheck{username, password}
	buf, err := json.Marshal(uc)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(buf)
	req, err := a.NewRequest("POST", "check_user", reader)
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	//200 OK else false
	return resp.StatusCode == 200, nil
}

func (a *ApiClient) NewRequest(method string, path string, body io.Reader) (*http.Request, error) {
	u := &url.URL{
		Scheme: a.URL.Scheme,
		Host:   a.URL.Host,
		Path:   path,
	}

	return http.NewRequest(method, u.String(), body)
}
