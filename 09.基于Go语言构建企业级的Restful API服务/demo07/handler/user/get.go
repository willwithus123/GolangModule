package user

import (
	. "ApiServer/demo07/handler"
	"github.com/gin-gonic/gin"
	"ApiServer/demo07/model"
	"ApiServer/demo07/pkg/errno"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
