package handler

import (
	"app/m/base"
	"app/m/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

// 上传头像
func (g *Gin) UploadImage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("获取请求失败!")

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    err,
			"message": "获取数据失败!",
		})
		return
	}

	// 获取文件名
	fileName := header.Filename
	str := strings.Split(fileName, ".")
	layout := strings.ToLower(str[len(str)-1])
	if layout != "jpeg" && layout != "png" && layout != "jpg" && layout != "gif" {
		log.Error("文件格式不正确!")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "文件格式不正确!",
		})
		return
	}

	if header.Size > 10000000 {
		//判断大小是否大于10M
		log.Error("文件过大!")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "文件大于10M，请重新上传",
		})
		return
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Error("Token解析失败: ", err)
		return
	}
	id := claims.ID
	username := claims.UserName
	filePath := "static/head/" + username + "-" + fileName
	exist, err := utils.PathExists(filePath)
	if err != nil {
		log.Error(err)
		return
	}
	// 存在
	if exist {
		return
	}
	// 不存在
	out, err := os.Create(filePath)
	if err != nil {
		log.Error("创建文件失败!")
		c.JSON(500, gin.H{
			"code":    500,
			"data":    err,
			"message": "创建文件失败!",
		})
		return
	}
	_, err = io.Copy(out, file)
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{
			"code":    500,
			"data":    err,
			"message": "保存文件失败!",
		})
		return
	}
	file.Close()
	out.Close()

	db := base.DB.Exec("update user set head_portrait=? where id=?", filePath, id)
	log.Info("更新: ", db.RowsAffected)

	c.JSON(200, gin.H{
		"code":    200,
		"data":    filePath,
		"message": "上传图片成功!",
	})

}
