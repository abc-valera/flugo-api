package dto

import "github.com/abc-valera/flugo-api/internal/domain"

// UserResponse type is returned back with response. It omits unnecessary data from the database's user type.
type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Status   string `json:"status"`
	Bio      string `json:"bio"`
}

// NewUserResponse returns new UserResponse from default user type
func NewUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		user.Username,
		user.Email,
		user.Fullname,
		user.Status,
		user.Bio,
	}
}

type UsersResponse []*UserResponse

func NewUsersResponse(users domain.Users) UsersResponse {
	usersResponse := make(UsersResponse, len(users))
	for i, user := range users {
		usersResponse[i] = NewUserResponse(user)
	}
	return usersResponse
}

type UpdateMyPasswordRequest struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

type UpdateMyFullnameRequest struct {
	Fullname string `json:"fullname"`
}

type UpdateMyStatusRequest struct {
	Status string `json:"status"`
}

type UpdateMyBioRequest struct {
	Bio string `json:"bio"`
}

type DeleteMeRequest struct {
	Password string `json:"password"`
}
