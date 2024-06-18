package non_strict_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NonStrictServer struct {
	userService *UserService
}

func NewServer() NonStrictServer {
	return NonStrictServer{userService: NewUserService()}
}

func (s NonStrictServer) CreateUser(ctx *gin.Context) {

	var body CreateUser
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	user, _ := s.userService.CreateUser(body)

	ctx.JSON(http.StatusCreated, user)
}

func (s NonStrictServer) DeleteUser(c *gin.Context, userId int64) {
	err := s.userService.DeleteUser(userId)
	if err != nil {
		return
	}

	c.JSON(http.StatusNoContent, err)
}

func (s NonStrictServer) GetUser(c *gin.Context, userId int64) {
	user, _ := s.userService.GetUser(userId)

	c.JSON(http.StatusOK, user)
}

func (s NonStrictServer) PutApiV1UsersUserId(ctx *gin.Context, userId int64) {
	var body User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	user, _ := s.userService.UpdateUser(userId, body.Email, body.Lastname, body.Name)

	ctx.JSON(http.StatusOK, user)
}
