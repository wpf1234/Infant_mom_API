package handler

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

// 我的页面---首页
func (g *Gin) Mine(c *gin.Context) {
	var mine model.MineRes
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
	id := claims.(*utils.MyClaims).ID
	mine.User.Id = id
	mine.User.UserName = claims.(*utils.MyClaims).UserName
	mine.User.Mobile = claims.(*utils.MyClaims).Mobile
	mine.User.HeadPortrait = claims.(*utils.MyClaims).HeadPortrait
	// 获取关注的商品数
	db := base.DB.Raw("select count(1) from record where user_id=? and record_type=?",
		id, 1)
	db.Row().Scan(&mine.Attention)
	// 获取喜欢的商品数量
	db = base.DB.Raw("select count(1) from record where user_id=? and record_type=?",
		id, 2)
	db.Row().Scan(&mine.Favorite)
	// 获取浏览记录
	db = base.DB.Raw("select count(1) from record where user_id=? and record_type=?",
		id, 3)
	db.Row().Scan(&mine.Browse)
	// 获取红包卡券数量
	db = base.DB.Raw("select count(1) from packets where user_id=?", id)
	db.Row().Scan(&mine.Packets)
	// 获取待付款数量 1
	db = base.DB.Raw("select count(1) from order where user_id=? and order_state", id, 1)
	db.Row().Scan(&mine.Pending)
	// 获取待收货数量 2
	db = base.DB.Raw("select count(1) from order where user_id=? and order_state", id, 2)
	db.Row().Scan(&mine.Received)
	// 获取待评价数量 3
	db = base.DB.Raw("select count(1) from order where user_id=? and order_state", id, 3)
	db.Row().Scan(&mine.Evaluate)
	// 获取退款/售后
	db = base.DB.Raw("select count(1) from order where user_id=? and order_state", id, 4)
	db.Row().Scan(&mine.Refund)

	fmt.Println(mine)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    mine,
		"message": "获取信息成功!",
	})
}

// 我的订单
func (g *Gin) MyOrder(c *gin.Context) {
	var myOrder model.MyOrder
	var total int
	var cntSql, infoSql string
	title := c.Query("title")
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
	id := claims.(*utils.MyClaims).ID
	goodsInfo := []model.GoodsInfo{}
	switch title {
	case "all":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state!=0 and user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state!=0 and user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
	case "pending":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state=1 user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state=1 and user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
	case "received":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state=2 user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state=2 and user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
	case "evaluate":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state=3 user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state=3 and user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
	case "refund":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state=4 user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_nutrient,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from order where order_state=4 and user_id=%d) as b 
							where a.id=b.g_id order by order_time desc`, id)
	default:
		break
	}

	db := sqlDB.Raw(cntSql)
	db.Row().Scan(&total)
	db = sqlDB.Raw(infoSql)
	rows, err := db.Rows()
	if err != nil {
		log.Error("获取商品信息失败: ", err)
		return
	}
	for rows.Next() {
		var info model.GoodsInfo
		var gID int
		rows.Scan(&info.Id, &info.Name, &info.Nutrient, &info.Picture, &info.Price,
			&gID, &info.Number)
		goodsInfo = append(goodsInfo, info)
	}
	myOrder.Info = goodsInfo
	myOrder.Total = total
	fmt.Println("My order: ", myOrder)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    myOrder,
		"message": "获取订单信息成功!",
	})
}

// 收获地址
func (g *Gin) GetAddress(c *gin.Context) {
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
	id := claims.(*utils.MyClaims).ID

	var addresses []model.Address
	db := sqlDB.Raw("select id,address,linkman,phone,is_default from address where user_id=? and is_del=?", id, 0)
	rows, err := db.Rows()
	if err != nil {
		log.Error("获取收货地址失败: ", err)
		return
	}
	for rows.Next() {
		var defaultAddr int
		var addr model.Address
		rows.Scan(&addr.Id, &addr.Address, &addr.Linkman, &addr.Phone, &defaultAddr)
		if defaultAddr == 1 {
			addr.DefaultAddr = true
		} else {
			addr.DefaultAddr = false
		}
		addresses = append(addresses, addr)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    addresses,
		"message": "获取收货地址成功!",
	})
}

// 更新收货地址
func (g *Gin) PutAddress(c *gin.Context) {
	var addr model.Address
	var defaultAddr int
	err := c.BindJSON(&addr)
	if err != nil {
		log.Error("获取请求参数失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "请求失败!",
		})
		return
	}

	if addr.DefaultAddr == true {
		defaultAddr = 1
	} else {
		defaultAddr = 0
	}
	db := sqlDB.Exec("update address set address=?,linkman=?,phone=?,is_default=? where id=?",
		addr.Address, addr.Linkman, addr.Phone, defaultAddr)
	err = db.Error
	if err != nil {
		log.Error("更新失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "更新收货地址失败!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "更新收货地址成功!",
	})
}

// 新增收获地址
func (g *Gin) PostAddress(c *gin.Context) {
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
	id := claims.(*utils.MyClaims).ID

	var addr model.Address
	err := c.BindJSON(&addr)
	if err != nil {
		log.Error("获取请求参数失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "请求失败!",
		})
		return
	}
	var isDefault int
	if addr.DefaultAddr == true {
		isDefault = 1
	} else {
		isDefault = 0
	}

	db := sqlDB.Exec("insert into address set user_id=?,address=?,linkman=?,phone=?,is_default=?",
		id, addr.Address, addr.Linkman, addr.Phone, isDefault)
	err = db.Error
	if err != nil {
		log.Error("新增失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "新增收货地址失败!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "新增收货地址成功!",
	})
}

//删除收货地址
func (g *Gin) DelAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Error("获取请求参数失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "请求失败!",
		})
		return
	}
	db := sqlDB.Exec("update address set is_del=? where id=?", 1, id)
	fmt.Println("删除: ", db.RowsAffected)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "删除成功!",
	})
}
