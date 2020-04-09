package handler

import (
	"app/m/base"
	"app/m/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// 查询，一般是按照名称查询 (或者营养元素？)
func (g *Gin) Search(c *gin.Context) {
	var sql string
	var details []model.GoodsDetails
	tp := c.Query("type")   // 查询类型
	text := c.Query("text") // 查询的内容
	if tp == "1" {
		// 默认按照名称查询
		sql = fmt.Sprintf("select id,g_name,g_nutrient,g_picture,g_price,classify,content from goods where g_name like '%s'",
			"%"+text+"%")
	} else if tp == "2" {
		sql = fmt.Sprintf("select id,g_name,g_nutrient,g_picture,g_price,classify,content from goods where classify like '%s'",
			"%"+text+"%")
	} else {
		sql = fmt.Sprintf("select id,g_name,g_nutrient,g_picture,g_price,classify,content from goods where g_picture like '%s'",
			"%"+text+"%")
	}
	db := base.DB.Raw(sql)
	rows, err := db.Rows()
	if err != nil {
		log.Error("查询失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "查询失败!",
		})
		return
	}

	for rows.Next() {
		var d model.GoodsDetails
		var nutrient string
		_ = rows.Scan(&d.Id, &d.Name, &nutrient, &d.Picture, &d.Price,
			&d.Classify, &d.Content)
		_ = json.Unmarshal([]byte(nutrient), &d.Nutrient)
		details = append(details, d)
	}
	_ = rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    details,
		"message": "查询成功!",
	})
}
