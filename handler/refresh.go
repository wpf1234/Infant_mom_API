package handler

import (
	. "app/m/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Gin struct{}

func (g *Gin) Refresh(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, err := ParseToken(token)
	if err != nil {
		log.Error("Token解析失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "Token解析失败!",
		})
		return
	}
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(7*ExpireTime) * time.Hour).Unix()
	newToken, err := GetToken(claims)
	if err != nil {
		log.Error("刷新失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "Token刷新失败!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    newToken,
		"message": "刷新成功!",
	})
}
