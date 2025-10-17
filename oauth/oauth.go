package oauth

import (
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/skyrocket-qy/erx"
	social "github.com/skyrocket-qy/line-login-sdk-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"

	"github.com/skyrocket-qy/iam/configs"
)

type OauthClient struct {
	LineCli   *social.Client
	GoogleCfg *oauth2.Config
	FbCfg     *oauth2.Config
	Cfg       *configs.OauthConfig
}

func NewOuathClient(oauthCfg *configs.OauthConfig) (*OauthClient, error) {
	var oauthClient *OauthClient
	if oauthCfg != nil {
		oauthClient = &OauthClient{}
		_, err := linebot.New(oauthCfg.LineChannelSecret, oauthCfg.LineChannelAccessToken)
		if err != nil {
			return nil, erx.W(err)
		}
		lineClient, err := social.New(oauthCfg.LineChannelID, oauthCfg.LineChannelSecret)
		if err != nil {
			return nil, erx.W(err)
		}
		googleCfg := &oauth2.Config{
			ClientID:     oauthCfg.GoogleID,
			ClientSecret: oauthCfg.GoogleSecret,
			RedirectURL:  oauthCfg.GoogleRedirectURL,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		}
		fbCfg := &oauth2.Config{
			ClientID:     oauthCfg.FbID,
			ClientSecret: oauthCfg.FbSecret,
			RedirectURL:  oauthCfg.FbRedirectURL,
			Scopes:       []string{"public_profile"},
			Endpoint:     facebook.Endpoint,
		}
		oauthClient.LineCli = lineClient
		oauthClient.GoogleCfg = googleCfg
		oauthClient.FbCfg = fbCfg

		oauthClient.Cfg = oauthCfg
	}

	return oauthClient, nil

}

type GoogleUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	Hd            string `json:"hd"`
}

type FBUser struct {
	ID      string
	Name    string
	Email   string
	Picture string
}

type AppleUser struct {
	Iss   string `json:"iss"`   //"iss": "https://appleid.apple.com",
	Aud   string `json:"aud"`   //"aud": "com.Rich.MegaRich",
	Exp   uint64 `json:"exp"`   //"exp": 1231245234,
	Iat   uint64 `json:"iat"`   //"iat": 1231245234,
	Sub   string `json:"sub"`   //"sub": "000091.4219285fe6ed4100992f093decc05009.0955",
	CHash string `json:"cHash"` //"c_hash": "EkfJCbJdLKzDOO1qhCe_CA",
	Name  string `json:"name"`
}
