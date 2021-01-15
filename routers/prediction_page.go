package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task5/models"
	"task5/pkg/util"
)

func serviceGet(c *gin.Context) {
	util.PageLog.Info("获取查询服务")
	c.HTML(http.StatusOK, "index.html", nil)
}

func servicePost(c *gin.Context) {
	//获取查询的timePoint(起始时间), hour(持续时长), serviceType(预测物)
	timePoint := c.PostForm("time")
	substance := c.PostForm("substance")
	hour := c.PostForm("hour")
	_, err := strconv.Atoi(hour)
	if err != nil {
		util.PageLog.Info("输入格式不正确")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "输入格式不正确",
			"code": http.StatusBadRequest,
		})
	} else {
		util.DaoLog.Info("按照时间节点和污染物条件查询")
		mlouts := models.ListMlouts(timePoint, hour, substance)
		c.HTML(http.StatusOK, "prediction_list.html", gin.H{
			"code": http.StatusOK,
			"data": mlouts,
		})
	}
}
