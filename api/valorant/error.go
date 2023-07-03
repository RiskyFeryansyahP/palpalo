package valorant

import "fmt"

var (
	ErrAuthCookieRequestNotOK   = fmt.Errorf("failed to request auth cookies is not ok")
	ErrAuthRequestNotOK         = fmt.Errorf("failed to request authorization is not ok")
	ErrEntitlementRequestNotOK  = fmt.Errorf("failed to request entitlement is not ok")
	ErrPlayerDetailRequestNotOK = fmt.Errorf("failed to request player detail is not ok")
	ErrValorantAPIAuthFailure   = fmt.Errorf("failed to authorize username and password")
)
