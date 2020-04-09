package handler

import (
	"app/m/base"
	"app/m/model"
	"app/m/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// 获取商品详情
func (g *Gin) GetDetails(c *gin.Context) {
	var details model.GoodsDetails
	var goodsRes model.GoodsRes
	var userId int
	claims, ok := c.Get("claims")
	if !ok {
		log.Warn("Claims字段不存在!")
	} else {
		userId = claims.(*utils.MyClaims).ID // user_id
	}

	goodsId, err := strconv.Atoi(c.Query("goods"))
	if err != nil {
		log.Error("数据转换失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "请求失败!",
		})
		return
	}
	db := base.DB.Raw("select id,g_name,g_nutrient,g_picture,g_price,classify,content from goods where id=?", goodsId)
	if err := db.Error; err != nil {
		log.Error("查询商品信息失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "查询商品信息失败!",
		})
		return
	}
	var nutrient string
	_ = db.Row().Scan(&details.Id, &details.Name, &nutrient, &details.Picture, &details.Price,
		&details.Classify, &details.Content)
	_ = json.Unmarshal([]byte(nutrient), &details.Nutrient)
	fmt.Println("Goods: ", details)
	goodsRes.Details = details
	// 查看是否关注和喜爱
	if userId != 0 {
		goodsRes.Attention = false
		goodsRes.Favorite = false
		db = base.DB.Raw("select record_type from record where user_id=? and g_id=?",
			userId, goodsId)
		rows, err := db.Rows()
		if err != nil {
			log.Error("查询关注/喜爱状态失败: ", err)
			return
		}
		for rows.Next() {
			var state int
			_ = rows.Scan(&state)
			if state == 1 {
				goodsRes.Attention = true
			}
			if state == 2 {
				goodsRes.Favorite = true
			}
		}
		_ = rows.Close()
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    goodsRes,
		"message": "获取商品详情成功!",
	})

	// 查看详情就相当于进行了一次浏览
	if userId != 0 {
		browse(userId, goodsId)
	}

}

// 查看详情，记录浏览记录
func browse(userId, goodsId int) {
	var id, browseTm int
	// 获取最新一条浏览记录
	db := base.DB.Raw(`select id,browse_time from record where user_id=? and g_id=? and record_type=? 
						 order by browse_time desc limit 1`,
		userId, goodsId, 3)
	_ = db.Row().Scan(&id, &browseTm)
	tm := time.Now()
	t := time.Unix(int64(browseTm), 0)
	// 没有这一个商品的浏览记录
	if id == 0 {
		db = base.DB.Exec("insert into record set user_id=?,g_id=?,record_type=?,browse_time=?",
			userId, goodsId, 3, tm.Unix())
		if err := db.Error; err != nil {
			log.Error("新增失败: ", err)
			return
		}
	}
	// 查到相应记录，判断时间间隔
	if tm.Sub(t).Hours() > 24 {
		db = base.DB.Exec("insert into record set user_id=?,g_id=?,record_type=?,browse_time=?",
			userId, goodsId, 3, tm.Unix())
		if err := db.Error; err != nil {
			log.Error("新增失败: ", err)
			return
		}
	} else {
		db = base.DB.Exec("update record set browse_time=? where id=?",
			tm.Unix(), id)
		if err := db.Error; err != nil {
			log.Error("新增失败: ", err)
			return
		}
	}
}
