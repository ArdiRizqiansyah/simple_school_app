package dto

type StudentCreateRequest struct {
	Name       string `json:"name" valid:"required~Name can't be empty" example:"Student 1"`
	Nis        string `json:"nis" valid:"required~NIS can't be empty" example:"1234567890"`
	PlaceBirth string `json:"place_birth" valid:"required~Place Birth can't be empty" example:"Jakarta"`
	DateBirth  string `json:"date_birth" valid:"required~Date Birth can't be empty" example:"2000-01-01"`
}

type StudentUpdateRequest struct {
	Name       string `json:"name" valid:"required~Name can't be empty" example:"Student 1"`
	Nis        string `json:"nis" valid:"required~NIS can't be empty" example:"1234567890"`
	PlaceBirth string `json:"place_birth" valid:"required~Place Birth can't be empty" example:"Jakarta"`
	DateBirth  string `json:"date_birth" valid:"required~Date Birth can't be empty" example:"2000-01-01"`
}

type StudentResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Nis        string `json:"nis"`
	PlaceBirth string `json:"place_birth"`
	DateBirth  string `json:"date_birth"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type StudentWithPagination struct {
	Data []StudentResponse `json:"data"`
	Page Page              `json:"page"`
}
