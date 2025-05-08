package response

type LoginUserResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type GetUserInfoResponse struct{}
