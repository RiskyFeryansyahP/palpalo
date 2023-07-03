package valorant

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	cookiePkg "github.com/confus1on/palpalo/pkg/cookie"
)

// AuthCookies
//
// Authorize the client to send a request to the server and obtain the cookie value `ASID`.
// This way, when the client includes the `ASID` value in subsequent requests to the server,
// it will not receive a "forbidden" response.
func (a *API) AuthCookies(ctx context.Context) (AuthCookieResponse, error) {
	// TODO:
	// need to re-organize base url and path for each valorant api url
	rawURL := "https://auth.riotgames.com/api/v1/authorization"

	payload := AuthCookieRequest{
		ClientID:     a.clientID,
		Nonce:        1,
		RedirectURI:  a.redirectURI,
		ResponseType: "token id_token",
		Scope:        "account openid",
	}

	reqBody, err := json.Marshal(payload)

	if err != nil {
		return AuthCookieResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, bytes.NewBuffer(reqBody))

	if err != nil {
		return AuthCookieResponse{}, err
	}

	a.defaultRequestHeader(req)

	req.Header.Set("Referer", req.URL.Host)

	resp, err := a.httpClient.Do(req)

	if err != nil {
		return AuthCookieResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AuthCookieResponse{}, ErrAuthCookieRequestNotOK
	}

	var authCookieResp AuthCookieResponse

	err = json.NewDecoder(resp.Body).Decode(&authCookieResp)

	if err != nil {
		return AuthCookieResponse{}, err
	}

	cookies := resp.Header.Values("set-cookie")

	ASIDCookie := cookiePkg.GetCookieByPrefix(cookies, "asid")

	authCookieResp.ASIDCookie = ASIDCookie

	return authCookieResp, nil
}

// Authorize
//
// Authorize the Valorant user by providing the username and password to the server.
// In response, the server will provide the client with an Access Token, ID Token,
// and SSID (Session ID that will be used as a cookie value).
// With these credentials, the client can consume other APIs such as player details,
// player store, and more.
func (a *API) Authorize(
	ctx context.Context,
	payload AuthRequest,
	asidCookie string,
) (AuthResponse, error) {
	// TODO:
	// need to re-organize base url and path for each valorant api url
	rawURL := "https://auth.riotgames.com/api/v1/authorization"

	reqBody, err := json.Marshal(payload)

	if err != nil {
		return AuthResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, rawURL, bytes.NewBuffer(reqBody))

	if err != nil {
		return AuthResponse{}, err
	}

	a.defaultRequestHeader(req)

	req.Header.Set("cookie", asidCookie)
	req.Header.Set("Referer", req.URL.Host)

	resp, err := a.httpClient.Do(req)

	if err != nil {
		return AuthResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AuthResponse{}, ErrAuthRequestNotOK
	}

	var authResponse AuthResponse

	err = json.NewDecoder(resp.Body).Decode(&authResponse)

	if err != nil {
		return AuthResponse{}, err
	}

	if authResponse.Error != "" {
		return AuthResponse{}, ErrValorantAPIAuthFailure
	}

	// NOTE
	// Since our `access_token` and `id_token` are located in the response body field `URI`,
	// we need to parse the URL query to retrieve them.
	// Once we parse the query, we can obtain the `access_token` and `id_token`.
	// example response:
	/*
		{
			"type": "...",
			"response": {
				"mode": "...",
				"parameters": {
				"uri": "https://playvalorant.com/opt_in#access_token=...&scope=...&id_token=...&token_type=Bearer&session_state=...&expires_in=3600"
				}
			},
			"country": "..."
		}
	*/
	u, err := url.Parse(authResponse.Response.Parameters.URI)

	if err != nil {
		return AuthResponse{}, err
	}

	urlValues, err := url.ParseQuery(u.Fragment)

	if err != nil {
		return AuthResponse{}, err
	}

	accessToken := urlValues.Get("access_token")
	idToken := urlValues.Get("id_token")

	authResponse.AccessToken = accessToken
	authResponse.IDToken = idToken

	cookies := resp.Header.Values("set-cookie")

	SSIDCookie := cookiePkg.GetCookieByPrefix(cookies, "ssid")

	authResponse.SSIDCookie = SSIDCookie

	return authResponse, nil
}

// Entitlement
//
// Obtain an entitlement token that will be used in the header, such as `X-Riot-Entitlements-JWT`,
// for making remote requests to the Valorant server.
// This token is required to authenticate and authorize the client for accessing certain
// resources or functionalities within the Valorant system.
func (a *API) Entitlement(
	ctx context.Context,
	ssidCookie string,
	accessToken string,
) (EntitlementResponse, error) {
	// TODO:
	// need to re-organize base url and path for each valorant api url
	rawURL := "https://entitlements.auth.riotgames.com/api/token/v1"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, nil)

	if err != nil {
		return EntitlementResponse{}, err
	}

	a.defaultRequestHeader(req)

	req.Header.Set("cookie", ssidCookie)
	req.Header.Set("Referer", req.URL.Host)
	req.Header.Set("Authorization", generateAuthorizationKey(accessToken))

	resp, err := a.httpClient.Do(req)

	if err != nil {
		return EntitlementResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return EntitlementResponse{}, ErrEntitlementRequestNotOK
	}

	var entitlementResponse EntitlementResponse

	err = json.NewDecoder(resp.Body).Decode(&entitlementResponse)

	if err != nil {
		return EntitlementResponse{}, err
	}

	return entitlementResponse, nil
}

// ReAuthorize
//
// Reauthorize a user when their access token and entitlement token have expired.
// This way, the user does not need to provide their username and password again for authentication.
func (a *API) ReAuthorize(ctx context.Context, ssidCookie string) (string, error) {
	// TODO:
	// need to re-organize base url and path for each valorant api url
	rawURL := "https://auth.riotgames.com/authorize?redirect_uri=https%3A%2F%2Fplayvalorant.com%2Fopt_in&client_id=play-valorant-web-prod&response_type=token%20id_token&nonce=1"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)

	if err != nil {
		return "", err
	}

	a.defaultRequestHeader(req)

	req.Header.Set("cookie", ssidCookie)
	req.Header.Set("Referer", req.URL.Host)

	// NOTE
	// Since we need to modify the `http.Client.CheckRedirect` behavior,
	// I prefer not to modify it in-place within our `*httpClient` object.
	// This modification is only required for re-authentication requests.
	// Therefore, I suggest re-assigning the modified `httpClient` to a new variable.
	client := a.httpClient

	var newAccessToken string

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirectUrl := req.URL.String()

		u, err := url.Parse(redirectUrl)

		if err != nil {
			return err
		}

		fragment, err := url.QueryUnescape(u.Fragment)

		if err != nil {
			return err
		}

		values, err := url.ParseQuery(fragment)

		if err != nil {
			return err
		}

		// got the new access token
		newAccessToken = values.Get("access_token")

		return http.ErrUseLastResponse
	}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return newAccessToken, nil
}

// defaultRequestHeader
// generate default request header that will be sent to the valorant api server
// the default header is `content-type` and `user-agent`
func (a API) defaultRequestHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", a.userAgent)
}

// generateAuthorizationKey
// generating Authorization bearer with access token
// example: Authorization: Bearer ...
func generateAuthorizationKey(accessToken string) string {
	return fmt.Sprintf("Bearer %s", accessToken)
}
