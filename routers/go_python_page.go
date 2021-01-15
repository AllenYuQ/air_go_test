package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"task5/models"
	"task5/pkg/util"
)

//初始化rocket
func init() {
	util.StartRocketMQ()
}

func goPythonPost(c *gin.Context) {
	//获取查询的timePoint(起始时间), hour(持续时长), serviceType(预测物)
	hasDone := c.PostForm("hasDone")
	util.PageLog.Info("python端的数据抓取和模型计算完毕，开始生产消息" + hasDone)

	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  "接受消息成功！",
		"code": http.StatusOK,
	})

	/*
		mlmets := models.ListMlmets()
		util.RocketLog.Info("将消息发送给rocket")
		for _, mlmet := range(mlmets) {
			data, _ := json.Marshal(&mlmet)
			util.SendMessageToRocketChan(data, 1,"test")
		}*/
	//将预测数据发送给rocket mq
	mlouts := models.ListMloutsBetweenInterval("24")
	util.RocketLog.Info("将消息发送给rocket")
	for _, mlout := range mlouts {
		data, _ := json.Marshal(&mlout)
		util.SendMessageToRocketChan(data, 1, "test")
	}
	//消息发送完毕
}
