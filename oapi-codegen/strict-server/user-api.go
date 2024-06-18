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

func (s StrictServer) DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	err := s.userService.DeleteUser(request.UserId)

	return DeleteUser204Response{}, err
}

func (s StrictServer) GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error) {
	user, err := s.userService.GetUser(request.UserId)

	return GetUser200JSONResponse(*user), err
}

func (s StrictServer) PutApiV1UsersUserId(ctx context.Context, request PutApiV1UsersUserIdRequestObject) (PutApiV1UsersUserIdResponseObject, error) {
	user, err := s.userService.UpdateUser(request.UserId, request.Body.Email, request.Body.Lastname, request.Body.Name)

	return PutApiV1UsersUserId200JSONResponse(*user), err
}
