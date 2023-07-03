package valorant

import (
	"context"
	"encoding/json"
	"net/http"
)

// PlayerDetail
//
// Retrieve player information, including the player's UUID (generated using UUIDv4)
// and other details such as game name, tagline, age, and more.
func (a *API) PlayerDetail(ctx context.Context, ssidCookie string, accessToken string) (PlayerDetailResponse, error) {
	// TODO:
	// need to re-organize base url and path for each valorant api url
	rawURL := "https://auth.riotgames.com/userinfo"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, nil)

	if err != nil {
		return PlayerDetailResponse{}, err
	}

	a.defaultRequestHeader(req)

	req.Header.Set("Referer", req.URL.Host)
	req.Header.Set("cookie", ssidCookie)
	req.Header.Set("Authorization", generateAuthorizationKey(accessToken))

	resp, err := a.httpClient.Do(req)

	if err != nil {
		return PlayerDetailResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PlayerDetailResponse{}, ErrPlayerDetailRequestNotOK
	}

	var playerDetailResponse PlayerDetailResponse

	err = json.NewDecoder(resp.Body).Decode(&playerDetailResponse)

	if err != nil {
		return PlayerDetailResponse{}, err
	}

	return playerDetailResponse, nil
}
