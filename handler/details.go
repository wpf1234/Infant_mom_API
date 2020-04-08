package handler

import (
	"app/m/model"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func (g *Gin) GetDetails(c *gin.Context) {
	var details model.GoodsDetails
	var goodsRes model.GoodsRes
	userId, err := strconv.Atoi(c.Query("user_id"))
	goodsId, err := strconv.Atoi(c.Query("goods_id"))
	if err != nil {
		log.Error("数据转换失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "请求失败!",
		})
		return
	}
	db := sqlDB.Raw("select * from goods where id=?", goodsId)
	if err := db.Error; err != nil {
		log.Error("查询商品信息失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "查询商品信息失败!",
		})
		return
	}

	db.Row().Scan(&details.Id, &details.Name, &details.Nutrient, &details.Picture, &details.Price)
	fmt.Println("Goods: ", details)
	// 查看是否关注和喜爱
	goodsRes.Attention = false
	goodsRes.Favorite = false
	db = sqlDB.Raw("select record_type from record where user_id=? and g_id=?",
		userId, goodsId)
	rows, err := db.Rows()
	if err != nil {
		log.Error("查询关注/喜爱状态失败: ", err)
		return
	}
	for rows.Next() {
		var state int
		rows.Scan(&state)
		if state == 1 {
			goodsRes.Attention = true
		}
		if state == 2 {
			goodsRes.Favorite = true
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": goodsRes,
		"message":"获取商品详情成功!",
	})

	// 查看详情就相当于进行了一次浏览
	browse(userId,goodsId)
}

func browse(userId, goodsId int) {
	var id,browseTm int
	// 获取最新一条浏览记录
	db := sqlDB.Raw(`select id,browse_time from record where user_id=? and g_id=? and record_type=? 
						 order by browse_time desc limit 1`,
		userId, goodsId, 3)
	db.Row().Scan(&id,&browseTm)
	tm := time.Now()
	t := time.Unix(int64(browseTm), 0)
	// 没有这一个商品的浏览记录
	if id == 0 {
		db = sqlDB.Exec("insert into record set user_id=?,g_id=?,record_type=?,browse_time=?",
			userId, goodsId, 3, tm.Unix())
		if err := db.Error; err != nil {
			log.Error("新增失败: ", err)
			return
		}
	}
	// 查到相应记录，判断时间间隔
	if tm.Sub(t).Hours() > 24 {
		db = sqlDB.Exec("insert into record set user_id=?,g_id=?,record_type=?,browse_time=?",
			userId, goodsId, 3, tm.Unix())
		if err := db.Error; err != nil {
			log.Error("新增失败: ", err)
			return
		}
	}else {
		db = sqlDB.Exec("update record set browse_time=? where id=?",
			tm.Unix(),id)
		if err := db.Error; err != nil {
			log.Error("新增失败: ", err)
			return
		}
	}
}
