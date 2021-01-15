package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task5/models"
	"task5/pkg/util"
)

func listUserHandler(c *gin.Context) {
	util.PageLog.Info("列出所有用户")
	userList := models.ListUsers()
	c.HTML(http.StatusOK, "user_list.html", gin.H{
		"code": http.StatusOK,
		"data": userList,
	})
}

func loginGet(c *gin.Context) {
	util.PageLog.Info("登录请求")
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginPost(c *gin.Context) {
	util.DaoLog.Info("loginPost")
	userName := c.PostForm("username")
	password := c.PostForm("password")
	isExist := models.CheckUser(userName, password)
	if isExist == false {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户不存在",
		})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Name":     userName,
		"Password": password,
	})
}
