package openapigenonlyinterface

import (
	openapigengin "github.com/bbakla/openoapi-code-generator/oapi_generator_userapi"
	//openapigen_gin "github.com/bbakla/openapi3-with-go/openapi-generator/gin-server-gen/oapi-go-codegen"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserAPI struct {
	userService *UserService
}

func NewUserAPI() UserAPI {
	return UserAPI{userService: NewUserService()}
}

// Put /oapi-codegen/v1/users/:userId
func (api UserAPI) ApiV1UsersUserIdPut(c *gin.Context) {
	var body openapigengin.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		c.Error(err)
		return
	}

	user, _ := api.userService.UpdateUser(body.Id, body.Email, body.Email, body.Name)
	c.JSON(http.StatusCreated, user)

}

// Post /oapi-codegen/v1/users
// creates a user
func (api UserAPI) CreateUser(c *gin.Context) {
	var body openapigengin.CreateUser
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		c.Error(err)
		return
	}

	user, _ := api.userService.CreateUser(body)

	c.JSON(http.StatusCreated, user)
}

// Delete /oapi-codegen/v1/users/:userId
// delete a user
func (api UserAPI) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	i, err := strconv.ParseInt(userId, 10, 64)

	err = api.userService.DeleteUser(i)

	c.JSON(http.StatusNoContent, err)
}

// Get /oapi-codegen/v1/users/:userId
// get user
func (api UserAPI) GetUser(c *gin.Context) {
	userId := c.Param("userId")
	i, _ := strconv.ParseInt(userId, 10, 64)

	user, _ := api.userService.GetUser(i)

	c.JSON(http.StatusOK, user)
}
