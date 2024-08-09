package dto

type ClassroomCreateRequest struct {
	Name string `json:"name" valid:"required~Name can't be empty" example:"Classroom 1"`
}

type ClassroomUpdateRequest struct {
	Name string `json:"name" valid:"required~Name can't be empty" example:"Classroom 1"`
}

type ClassroomResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ClassroomWithPagination struct {
	Data []ClassroomResponse `json:"data"`
	Page Page                `json:"page"`
}
