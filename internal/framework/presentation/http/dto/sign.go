package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	User         *UserResponse `json:"user"`
}

type SignRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type SignRefreshResponse struct {
	AccessToken string `json:"access_token"`
}
