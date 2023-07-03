package valorant

type AuthCookieRequest struct {
	ClientID     string `json:"client_id"`
	Nonce        int8   `json:"nonce"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
}

type AuthCookieResponse struct {
	Type       string `json:"type"`
	Country    string `json:"country"`
	ASIDCookie string `json:"asid_cookie"`
}

type AuthRequest struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type AuthResponseDetailParameter struct {
	URI string `json:"uri"`
}

type AuthResponseDetail struct {
	Parameters AuthResponseDetailParameter `json:"parameters"`
	Mode       string                      `json:"mode"`
	Country    string                      `json:"country"`
}

type AuthResponse struct {
	Response    AuthResponseDetail `json:"response"`
	Type        string             `json:"type"`
	Error       string             `json:"error"`
	AccessToken string             `json:"access_token"`
	IDToken     string             `json:"id_token"`
	SSIDCookie  string             `json:"ssid_cookie"`
}

type ReAuthorizeResponse struct {
	AccessToken string `json:"access_token"`
	SSIDCookie  string `json:"SSIDCookie"`
}

type EntitlementResponse struct {
	Token string `json:"entitlements_token"`
}
