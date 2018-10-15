package user

import (
	"fmt"
	"net/http"

	"mygithub/Gin/gin-restful-learn/apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Debugf("username is:[%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		// 返回给用户的是 {"code":20102,"message":"The user was not found. This is add message."}
		err = errno.New(errno.ErrUserNotFpund, fmt.Errorf("username can not found in db:xx.xx.xx.xx")).Add("This is add message.")
		// 输出到日志文件
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		// 输出到运行日志
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		// 输出到后台日志
		err = fmt.Errorf("password is empty")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
