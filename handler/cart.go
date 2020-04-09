package handler

// 购物车
import (
	"app/m/base"
	"app/m/model"
	"app/m/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Gin struct{}

const (
	strTm = "2006-01-02 15:04:05"
)

// 获取购物车信息
func (g *Gin) GetCart(c *gin.Context) {
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
	id := claims.(*utils.MyClaims).ID // user_id
	var carts []model.Cart
	sql := fmt.Sprintf(`select * from 
					(select id,g_id,g_num from cart where user_id=%d and is_del=%d ) as a 
					left join 
					(select id,g_name,g_picture,g_price,content from goods ) as b 
					on a.g_id=b.id `, id, 0)
	db := base.DB.Raw(sql)
	rows, err := db.Rows()
	if err != nil {
		log.Error("查询数据失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "查询数据失败!",
		})
		return
	}

	for rows.Next() {
		var c model.Cart
		var gId, goodId int
		rows.Scan(&c.Id, &gId, &c.GoodsNum,
			&goodId, &c.GoodsName, &c.GoodsPicture, &c.GoodsPrice, &c.GoodsContent)
		carts = append(carts, c)
	}
	rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    carts,
		"message": "查询成功!",
	})
}

// 添加购物车
func (g *Gin) AddCart(c *gin.Context) {
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
	var gNum int
	gId, _ := strconv.Atoi(c.Query("goods"))
	if gId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "未选择商品!",
		})
		return
	}
	if num := c.Query("num"); num != "" {
		gNum, _ = strconv.Atoi(num)
	} else {
		gNum = 1
	}

	db := base.DB.Raw("select g_num from cart where user_id=? and g_id=?", userId, gId)
	err := db.Error
	if err != nil {
		log.Error("查询失败: ", err)
		return
	}
	var gn int
	_ = db.Row().Scan(&gn)
	if gn != 0 {
		db = base.DB.Exec("update cart set g_num=? where user_id=? and g_id=?",
			gNum+gn, userId, gId)
		fmt.Println("Goods update num: ", db.RowsAffected)
	} else {
		db = base.DB.Exec("insert into cart set user_id=?,g_id=?,g_num=?",
			userId, gId, gNum)
		err = db.Error
		if err != nil {
			log.Error("新增失败: ", err)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"data":    err,
				"message": "加入购物车失败!",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "加入购物车成功!",
	})
}

// 修改数量
func (g *Gin) PutCart(c *gin.Context) {
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
	gId, _ := strconv.Atoi(c.Query("goods"))
	num, _ := strconv.Atoi(c.Query("num"))
	if gId == 0 && num == 0 {
		log.Error("数据为空")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "请求失败!",
		})
		return
	}
	db := base.DB.Exec("update cart set g_num=? where user_id=? and g_id=?",
		num, userId, gId)
	err := db.Error
	if err != nil {
		log.Error("更改失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "更改失败!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "更改成功!",
	})
}

// 从购物车中删除
func (g *Gin) DelCart(c *gin.Context) {
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
	gId, _ := strconv.Atoi(c.Query("goods"))
	if gId == 0 {
		log.Error("数据为空")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "请求失败!",
		})
		return
	}
	db := base.DB.Exec("update cart set is_del=? where user_id=? and g_id=?",
		1, userId, gId)
	err := db.Error
	if err != nil {
		log.Error("删除失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "删除失败!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "从购物车中成功移除!",
	})
}
