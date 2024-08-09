package dto

import "time"

type UserLoginRequest struct {
	Email    string `json:"email" valid:"required~Email can't be empty" example:"johnDoemail.com"`
	Password string `json:"password" valid:"required~Password can't be empty" example:"password"`
}

type UserRegisterRequest struct {
	Name     string `json:"name" valid:"required~Name can't be empty" example:"John Doe"`
	Email    string `json:"email" valid:"required~Email can't be empty, email~Email has to be a valid format" example:"johnDoemail.com"`
	Password string `json:"password" valid:"required~Password can't be empty" example:"password"`
}

type TokenResponse struct {
	Token string `json:"token" example:"random string"`
}

type CheckProfileRequest struct {
	Token string `json:"token" valid:"required~Token can't be empty" example:"random string"`
}

type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListUserResponse struct {
	Users []UserResponse `json:"users"`
}
