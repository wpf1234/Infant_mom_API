package handler

import (
	"app/m/base"
	"app/m/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// 注册步骤: 输入电话号码，获取验证码，填写详细信息，注册完成
func (g *Gin) Register(c *gin.Context) {
	var register model.RegisterReq
	err := c.BindJSON(&register)
	if err != nil {
		log.Error("获取请求数据失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    err,
			"message": "请求失败!",
		})
		return
	}

	t := time.Now().Unix()
	db := base.DB.Exec("insert into user set user_name=?,password=?,mobile=?,register_time=?",
		register.UserName, register.Password, register.Phone, t)
	err = db.Error
	if err != nil {
		log.Error("新增失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "注册失败!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "注册成功!",
	})
}
