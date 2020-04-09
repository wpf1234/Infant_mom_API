package handler

import (
	"app/m/base"
	"app/m/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// 改变关注，喜欢等状态
func (g *Gin) ChangeState(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		log.Error("Claims字段不存在!")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "没有获取到信息!",
		})
		return
	}
	userId := claims.(*utils.MyClaims).ID // user_id
	goodsId, err := strconv.Atoi(c.Query("goods"))
	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		log.Error("数据转换失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    err,
			"message": "请求失败!",
		})
		return
	}
	if goodsId == 0 && state == 0 {
		log.Error("数据为空")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "请求失败!",
		})
		return
	}
	// 关注，喜欢或取消
	var id int
	if state == 1 || state == 2 {
		err := base.DB.Exec("insert into record set user_id=?,g_id=?,record_type=?",
			userId, goodsId, state).Error
		if err != nil {
			log.Error("添加商品失败: ", err)
			return
		}
	} else if state == -1 {
		db := base.DB.Raw("select id from record where user_id=? and g_id=? and record_type=?",
			userId, goodsId, 1)
		if err := db.Error; err != nil {
			log.Error("查询失败: ", err)
			return
		}
		_ = db.Row().Scan(&id)
		err := base.DB.Exec("update record set record_type=? where id=?", state, id).Error
		if err != nil {
			log.Error("更新失败: ", err)
			return
		}
	} else if state == -2 {
		db := base.DB.Raw("select id from record where user_id=? and g_id=? and record_type=?",
			userId, goodsId, 2)
		if err := db.Error; err != nil {
			log.Error("查询失败: ", err)
			return
		}
		_ = db.Row().Scan(&id)
		err := base.DB.Exec("update record set record_type=? where id=?", state, id).Error
		if err != nil {
			log.Error("更新失败: ", err)
			return
		}
	} else if state == -3 {
		db := base.DB.Raw("select id from record where user_id=? and g_id=? and record_type=?",
			userId, goodsId, 3)
		if err := db.Error; err != nil {
			log.Error("查询失败: ", err)
			return
		}
		_ = db.Row().Scan(&id)
		err := base.DB.Exec("update record set record_type=? where id=?", state, id).Error
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
