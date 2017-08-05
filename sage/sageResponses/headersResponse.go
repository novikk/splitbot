package sageResponses

type HeadersResponse struct {
	Credentials struct {
		RefreshToken    string `json:"refreshToken"`
		AccessToken     string `json:"accessToken"`
		ResourceOwnerID string `json:"resourceOwnerId"`
		ExpirationDate  int    `json:"expirationDate"`
	} `json:"credentials"`
	Headers map[string]string `json:"headers"`
}
