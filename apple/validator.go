package apple

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	// ValidationURL is the endpoint for verifying tokens
	ValidationURL string = "https://appleid.apple.com/auth/token"
	// ContentType is the one expected by Apple
	ContentType string = "application/x-www-form-urlencoded"
	// UserAgent is required by Apple or the request will fail
	UserAgent string = "go-signin-with-apple"
	// AcceptHeader is the content that we are willing to accept
	AcceptHeader string = "application/json"
)

// ValidationClient is an interface to call the validation API
type ValidationClient interface {
	VerifyWebToken(ctx context.Context, reqBody WebValidationTokenRequest, result any) error
	VerifyAppToken(ctx context.Context, reqBody AppValidationTokenRequest, result any) error
	VerifyRefreshToken(ctx context.Context, reqBody ValidationRefreshRequest, result any) error
}

type Client struct {
	validationURL string
	client        *http.Client
}

func New() *Client {
	client := &Client{
		validationURL: ValidationURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	return client
}

func NewWithURL(url string) *Client {
	client := &Client{
		validationURL: url,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	return client
}

func (c *Client) VerifyWebToken(ctx context.Context, reqBody WebValidationTokenRequest,
	result any,
) error {
	data := url.Values{}
	data.Set("client_id", reqBody.ClientID)
	data.Set("client_secret", reqBody.ClientSecret)
	data.Set("code", reqBody.Code)
	data.Set("redirect_uri", reqBody.RedirectURI)
	data.Set("grant_type", "authorization_code")

	return doRequest(ctx, c.client, &result, c.validationURL, data)
}

func (c *Client) VerifyAppToken(ctx context.Context, reqBody AppValidationTokenRequest,
	result any,
) error {
	data := url.Values{}
	data.Set("client_id", reqBody.ClientID)
	data.Set("client_secret", reqBody.ClientSecret)
	data.Set("code", reqBody.Code)
	data.Set("grant_type", "authorization_code")

	return doRequest(ctx, c.client, &result, c.validationURL, data)
}

func (c *Client) VerifyRefreshToken(ctx context.Context, reqBody ValidationRefreshRequest,
	result any,
) error {
	data := url.Values{}
	data.Set("client_id", reqBody.ClientID)
	data.Set("client_secret", reqBody.ClientSecret)
	data.Set("refresh_token", reqBody.RefreshToken)
	data.Set("grant_type", "refresh_token")

	return doRequest(ctx, c.client, &result, c.validationURL, data)
}

func doRequest(ctx context.Context, client *http.Client, result any, url string,
	data url.Values,
) error {
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", ContentType)
	req.Header.Add("accept", AcceptHeader)
	req.Header.Add("user-agent", UserAgent) // apple requires a user agent

	b, _ := httputil.DumpRequest(req, true)
	log.Info().Msgf("doRequest request %s", string(b))

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	b, _ = httputil.DumpResponse(res, true)
	log.Info().Msgf("doRequest response %s", string(b))

	defer func() {
		_ = res.Body.Close()
	}()

	return json.NewDecoder(res.Body).Decode(result)
}
