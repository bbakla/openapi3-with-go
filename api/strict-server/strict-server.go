package strict_server

import (
	"context"
)

type StrictServer struct {
	userService *UserService
}

func NewServer() StrictServer {
	return StrictServer{userService: NewUserService()}

}

func (s StrictServer) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {
	user, _ := s.userService.CreateUser(*request.Body)

	return CreateUser201JSONResponse{
		Id:       user.Id,
		Name:     user.Name,
		Lastname: user.Lastname,
		Email:    user.Email,
	}, nil
}

func (StrictServer) DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s StrictServer) GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
