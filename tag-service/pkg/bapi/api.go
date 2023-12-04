package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type API struct {
	URL string
}

const (
	APP_KEY    = "senyas"
	APP_SECRET = "senyas-tour"
)

type AccessToken struct {
	Token string `json:"token"`
}

func (a *API) getAsscessToken(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s?/app_key=%s&app_secret=%s", "auth", APP_KEY, APP_SECRET)
	body, err := a.httpGet(ctx, url)
	if err != nil {
		fmt.Println("a.httpGet err:", err)
		return "", err
	}
	var accessToken AccessToken
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		fmt.Println("json.Unmarchal err:", err)
		return "", err
	}
	return accessToken.Token, nil
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", a.URL, path)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io.ReadAll err:", err)
		return nil, err
	}
	return body, nil

}

func NewAPI(url string) *API {
	return &API{url}
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAsscessToken(ctx)
	if err != nil {
		fmt.Println("a.getAccessToken err", err)
		return nil, err
	}
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		fmt.Println("a.httpGet err:", err)
		return nil, err
	}
	return body, nil
}
