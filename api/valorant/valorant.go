package valorant

import "net/http"

type API struct {
	httpClient  *http.Client
	clientID    string
	redirectURI string
	userAgent   string
}

func NewValorantAPI(
	httpClient *http.Client,
	clientID string,
	redirectURI string,
	userAgent string,
) *API {
	return &API{
		httpClient:  httpClient,
		clientID:    clientID,
		redirectURI: redirectURI,
		userAgent:   userAgent,
	}
}
