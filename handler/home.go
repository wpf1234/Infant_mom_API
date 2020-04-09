package handler

import (
	"app/m/base"
	"app/m/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// 主页
func (g *Gin) Home(c *gin.Context) {
	var rec, newGoods, others []model.GoodsDetails
	var nutrient string
	var home model.Home
	// 获取推荐商品
	db := base.DB.Raw("select id,g_name,g_nutrient,g_picture,g_price,classify,content from goods where is_rec=1")
	rows, err := db.Rows()
	if err != nil {
		log.Error("查询推荐失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "获取推荐商品失败!",
		})
		return
	}
	for rows.Next() {
		var details model.GoodsDetails
		_ = rows.Scan(&details.Id, &details.Name, &nutrient, &details.Picture, &details.Price,
			&details.Classify, &details.Content)
		_ = json.Unmarshal([]byte(nutrient), &details.Nutrient)
		rec = append(rec, details)
	}

	// 获取新品
	db = base.DB.Raw("select id,g_name,g_nutrient,g_picture,g_price,classify,content from goods where is_new=1")
	rows, err = db.Rows()
	if err != nil {
		log.Error("查询新品失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "获取新品失败!",
		})
		return
	}
	for rows.Next() {
		var details model.GoodsDetails
		_ = rows.Scan(&details.Id, &details.Name, &nutrient, &details.Picture, &details.Price,
			&details.Classify, &details.Content)
		_ = json.Unmarshal([]byte(nutrient), &details.Nutrient)
		newGoods = append(newGoods, details)
	}

	// 获取其他商品
	db = base.DB.Raw("select id,g_name,g_nutrient,g_picture,g_price,classify,content from goods where is_new=0 and is_rec=0")
	rows, err = db.Rows()
	if err != nil {
		log.Error("查询商品失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "获取商品失败!",
		})
		return
	}
	for rows.Next() {
		var details model.GoodsDetails
		_ = rows.Scan(&details.Id, &details.Name, &nutrient, &details.Picture, &details.Price,
			&details.Classify, &details.Content)
		_ = json.Unmarshal([]byte(nutrient), &details.Nutrient)
		others = append(others, details)
	}
	_ = rows.Close()

	home.Recommend = rec
	home.NewGoods = newGoods
	home.Others = others

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    home,
		"message": "获取商品信息成功!",
	})
}
