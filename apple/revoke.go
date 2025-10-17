package apple

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type TokenHintType int32

const (
	// RevokeURL is the endpoint for revoke tokens
	RevokeURL string = "https://appleid.apple.com/auth/revoke"
	// RevokeContentType is the one expected by Apple
	RevokeContentType string = "application/x-www-form-urlencoded"
	// RevokeAcceptHeader is the content that we are willing to accept
	RevokeAcceptHeader string = "application/json"

	RefreshTokenTypeHint TokenHintType = 1
	AccessTokenTypeHint  TokenHintType = 2
)

type RevokeClient interface {
	RevokeToken(ctx context.Context, reqBody RevokeTokenRequest) error
}

type RVKClient struct {
	revokeURL string
	client    *http.Client
}

// New RVKClient object
func NewRVKClient() *RVKClient {
	client := &RVKClient{
		revokeURL: RevokeURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	return client
}

// RevokeToken sends the WebRevokeTokenRequest and gets revocation result
func (c *RVKClient) RevokeToken(ctx context.Context, reqBody RevokeTokenRequest) error {
	var typeHint string

	data := url.Values{}
	data.Set("client_id", reqBody.ClientID)
	data.Set("client_secret", reqBody.ClientSecret)
	data.Set("token", reqBody.Token)

	switch reqBody.TokenTypeHint {
	case RefreshTokenTypeHint:
		typeHint = "refresh_token"
	case AccessTokenTypeHint:
		typeHint = "access_token"
	default:
		return errors.New("Invalid token type hint")
	}

	data.Set("token_type_hint", typeHint)

	return sendRequest(ctx, c.client, c.revokeURL, data)
}

func sendRequest(ctx context.Context, client *http.Client, url string, data url.Values) error {
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", RevokeContentType)
	req.Header.Add("accept", RevokeAcceptHeader)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	return nil
}
