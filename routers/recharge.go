package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task5/models"
	"task5/pkg/util"
)

func rechargeGet(c *gin.Context) {
	util.PageLog.Info("充值余额页面")
	c.HTML(http.StatusOK, "recharge.html", nil)
}

func recharge(c *gin.Context) {
	//获取查询的timePoint(起始时间), hour(持续时长), serviceType(预测物)
	id := c.PostForm("id")
	count := c.PostForm("count")

	_, err := strconv.Atoi(count)
	if err != nil {
		util.PageLog.Info("输入格式不正确")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "输入格式不正确",
			"code": http.StatusBadRequest,
		})
	} else {
		util.PageLog.Info("进行充值操作")
		if models.ExistUser(id) == true {
			//编写相关逻辑
			isDone := models.Recharge(id, count)
			if isDone == true {
				quotas := models.ListQuotaInfo()
				c.HTML(http.StatusOK, "quota_list.html", gin.H{
					"code": http.StatusOK,
					"msg":  "充值成功!",
					"data": &quotas,
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg":  "充值失败!",
					"code": http.StatusBadRequest,
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "该用户id不存在，充值失败!",
				"code": http.StatusBadRequest,
			})
		}

	}
}
