package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (g *Gin) ChangeState(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	goodsId, err := strconv.Atoi(c.Query("goods_id"))
	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		log.Error("数据转换失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "请求失败!",
		})
		return
	}
	// 关注，喜欢或取消
	var id int
	if state == 1 || state == 2 {
		err := sqlDB.Exec("insert into record set user_id=?,g_id=?,record_type=?",
			userId, goodsId, state).Error
		if err != nil {
			log.Error("添加商品失败: ", err)
			return
		}
	} else if state == -1 {
		db := sqlDB.Raw("select id from record where user_id=? and g_id=? and record_type=?",
			userId, goodsId, 1)
		if err := db.Error; err != nil {
			log.Error("查询失败: ", err)
			return
		}
		db.Row().Scan(&id)
		err := sqlDB.Exec("update record set record_type=? where id=?", state, id).Error
		if err != nil {
			log.Error("更新失败: ", err)
			return
		}
	} else if state == -2 {
		db := sqlDB.Raw("select id from record where user_id=? and g_id=? and record_type=?",
			userId, goodsId, 2)
		if err := db.Error; err != nil {
			log.Error("查询失败: ", err)
			return
		}
		db.Row().Scan(&id)
		err := sqlDB.Exec("update record set record_type=? where id=?", state, id).Error
		if err != nil {
			log.Error("更新失败: ", err)
			return
		}
	} else if state == -3 {
		db := sqlDB.Raw("select id from record where user_id=? and g_id=? and record_type=?",
			userId, goodsId, 3)
		if err := db.Error; err != nil {
			log.Error("查询失败: ", err)
			return
		}
		db.Row().Scan(&id)
		err := sqlDB.Exec("update record set record_type=? where id=?", state, id).Error
		if err != nil {
			log.Error("更新失败: ", err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "修改成功!",
	})

}
