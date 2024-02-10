package service

import (
	"hexagonal/internal/core/model/request"
	"hexagonal/internal/core/model/response"
)

type UserService interface {
	SignUp(request *request.SignUpRequest) *response.Response
}
