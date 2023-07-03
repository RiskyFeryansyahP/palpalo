package valorant

type PlayerDetailAcct struct {
	Type     int    `json:"type"`
	State    string `json:"state"`
	Adm      bool   `json:"adm"`
	GameName string `json:"game_name"`
	TagLine  string `json:"tag_line"`
}

type PlayerDetailResponse struct {
	PlayerDetailAcct PlayerDetailAcct `json:"acct"`
	Country          string           `json:"country"`
	// Player UUID
	Sub           string `json:"sub"`
	EmailVerified bool   `json:"email_verified"`
	Age           int    `json:"age"`
	JTI           string `json:"jti"`
}
