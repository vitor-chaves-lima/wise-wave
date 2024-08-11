package dto

type UserSessionData struct {
	IdToken      string
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    int32
}
