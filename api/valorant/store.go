package valorant

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// DetailStoreByPlayerID
//
// Get detailed player store information based on the player's ID and region.
// This will display the night market, bundles, and offers available to the player.
func (a *API) DetailStoreByPlayerID(
	ctx context.Context,
	accessToken string,
	entitlementToken string,
	shard string,
	playerID string,
) (DetailStoreResponse, error) {
	rawURL := fmt.Sprintf("https://pd.%s.a.pvp.net/store/v2/storefront/%s", shard, playerID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)

	if err != nil {
		return DetailStoreResponse{}, err
	}

	req.Header.Set("Referer", req.URL.Host)
	req.Header.Set("Authorization", generateAuthorizationKey(accessToken))
	req.Header.Set("X-Riot-Entitlements-JWT", entitlementToken)

	resp, err := a.httpClient.Do(req)

	if err != nil {
		return DetailStoreResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return DetailStoreResponse{}, ErrStoreDetailNotOK
	}

	var storeDetailResponse DetailStoreResponse

	err = json.NewDecoder(resp.Body).Decode(&storeDetailResponse)

	if err != nil {
		return DetailStoreResponse{}, err
	}

	return storeDetailResponse, nil
}
