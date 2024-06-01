package not_strict_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	userService *UserService
}

func NewServer() Server {
	return Server{userService: NewUserService()}
}

func (s Server) CreateUser(ctx *gin.Context) {

	var body CreateUser
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	user, _ := s.userService.CreateUser(body)

	ctx.JSON(http.StatusOK, user)
}

func (s Server) DeleteUser(c *gin.Context, userId int64) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUser(c *gin.Context, userId int64) {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutApiV1UsersUserId(c *gin.Context, userId int64) {
	panic("implement me")
}
