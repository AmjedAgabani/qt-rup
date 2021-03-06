package oidc

import (
	"net/url"

	"github.com/pkg/browser"
)

func OpenBrowser(configuration OpenIDConfiguration) error {
	u, err := url.Parse(configuration.AuthorizationEndpoint)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("client_id", "gx35ttsecbzblr2ksi9a7l2beqsx8c")
	q.Set("redirect_uri", "http://localhost:30423")
	q.Set("response_type", "code")
	q.Set("scope", "channel:read:subscriptions")
	q.Set("claims", "preferred_username picture")
	u.RawQuery = q.Encode()
	return browser.OpenURL(u.String())
}
