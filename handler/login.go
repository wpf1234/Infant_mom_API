package handler

import (
	"app/m/base"
	"app/m/model"
	"app/m/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// 登录流程：输入账号密码，验证账号密码，错误---请注册，正确---返回token
func (g *Gin) Login(c *gin.Context) {
	var id int
	var username, headPortrait string
	var login model.LoginReq
	err := c.BindJSON(&login)
	if err != nil {
		log.Error("数据解析失败： ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "数据解析失败!",
		})
		return
	}
	db := base.DB.Raw("select id,user_name,head_portrait from user where mobile=? and password=?",
		login.Phone, login.Password)
	db.Row().Scan(&id, &username, &headPortrait)
	if id == 0 {
		log.Warn("没有该用户的信息!")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "用户信息不存在，请前往注册!",
		})
		return
	}
	// 存在用户信息，生成token
	claims := &utils.MyClaims{
		ID:           id,
		UserName:     username,
		Password:     login.Password,
		Mobile:       login.Phone,
		HeadPortrait: headPortrait,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                                             // 签名生效时间
			ExpiresAt: time.Now().Add(time.Duration(7*utils.ExpireTime) * time.Hour).Unix(), // 过期时间
		},
	}
	token, err := utils.GetToken(claims)
	if err != nil {
		log.Error("Token生成失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "Token生成失败!",
		})
		return
	}

	db = base.DB.Exec("update user set login_time=? where id=?", time.Now().Unix(), id)
	log.Info("更新: ", db.RowsAffected)
	data := model.LoginRes{
		User: model.User{
			Id:           id,
			UserName:     username,
			Mobile:       login.Phone,
			HeadPortrait: headPortrait,
		},
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": "登录成功!",
	})
}

// 使用第三方平台登录
