package user

import (
	"fmt"
	. "mygithub/Gin/gin-restful-learn/apiserver/handler"
	"mygithub/Gin/gin-restful-learn/apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest

	// Bind()
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, gin.H{"error": errno.ErrBind})
		return
	}

	// Param() http://127.0.0.1:8080/v1/user/admin2
	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	// Query() http://127.0.0.1:8080/v1/user/admin2?desc=""
	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	// GetHeader() -H "Content-Type: application/json"
	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	// -d'{"username":"admin","password":"admin"}'
	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		//// 返回给用户的是 {"code":20102,"message":"The user was not found. This is add message."}
		//err = errno.New(errno.ErrUserNotFpund, fmt.Errorf("username can not found in db:xx.xx.xx.xx")).Add("This is add message.")
		//// 输出到日志文件
		//log.Errorf(err, "Get an error")
		SendResponse(c, errno.New(errno.ErrUserNotFpund, fmt.Errorf("username can not found in db:xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		// 输出到后台日志
		//err = fmt.Errorf("password is empty")
		SendResponse(c, fmt.Errorf("password is empty"), nil)
		return
	}

	//if errno.IsErrUserNotFound(err) {
	//	// 输出到运行日志
	//	log.Debug("err type is ErrUserNotFound")
	//}

	rsp := CreateResponse{
		Username: r.Username,
	}

	//code, message := errno.DecodeErr(err)
	//c.JSON(http.StatusOK, gin.H{"code": code, "message": message})

	SendResponse(c, nil, rsp)
}
