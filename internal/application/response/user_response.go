package response

type LoginUserResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type GetUserInfoResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}
