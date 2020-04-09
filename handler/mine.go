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
	_ = db.Row().Scan(&mine.Attention)
	// 获取喜欢的商品数量
	db = base.DB.Raw("select count(1) from record where user_id=? and record_type=?",
		id, 2)
	_ = db.Row().Scan(&mine.Favorite)
	// 获取浏览记录
	db = base.DB.Raw("select count(1) from record where user_id=? and record_type=?",
		id, 3)
	_ = db.Row().Scan(&mine.Browse)
	// 获取红包卡券数量
	db = base.DB.Raw("select count(1) from packets where user_id=?", id)
	_ = db.Row().Scan(&mine.Packets)
	// 获取待付款数量 1
	db = base.DB.Raw("select count(1) from order where user_id=? and order_state", id, 1)
	_ = db.Row().Scan(&mine.Pending)
	// 获取待收货数量 2
	db = base.DB.Raw("select count(1) from order where user_id=? and order_state", id, 2)
	_ = db.Row().Scan(&mine.Received)
	// 获取待评价数量 3
	db = base.DB.Raw("select count(1) from order where user_id=? and order_state", id, 3)
	_ = db.Row().Scan(&mine.Evaluate)
	// 获取退款/售后
	db = base.DB.Raw("select count(1) from order where user_id=? and order_state", id, 4)
	_ = db.Row().Scan(&mine.Refund)

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
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from myorder where order_state!=0 and user_id=%d ) as b 
							on a.id=b.g_id `, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from myorder where order_state!=0 and user_id=%d order by order_time desc ) as b 
							on a.id=b.g_id `, id)
	case "pending":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from myorder where order_state=1 and user_id=%d ) as b 
							on a.id=b.g_id `, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num,order_time from myorder where order_state=1 and user_id=%d order by order_time desc ) as b 
							on a.id=b.g_id `, id)
	case "received":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from myorder where order_state=2 and user_id=%d ) as b 
							on a.id=b.g_id `, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num,order_time from myorder where order_state=2 and user_id=%d order by order_time desc ) as b 
							on a.id=b.g_id `, id)
	case "evaluate":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from myorder where order_state=3 and user_id=%d ) as b 
							on a.id=b.g_id `, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num,order_time from myorder where order_state=3 and user_id=%d order by order_time desc ) as b 
							on a.id=b.g_id `, id)
	case "refund":
		cntSql = fmt.Sprintf(`select count(1) from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num from myorder where order_state=4 and user_id=%d ) as b 
							on a.id=b.g_id `, id)
		infoSql = fmt.Sprintf(`select * from 
							(select id,g_name,g_picture,g_price from goods ) as a 
							left join 
							(select g_id,g_num,order_time from myorder where order_state=4 and user_id=%d order by order_time desc ) as b 
							on a.id=b.g_id `, id)
	default:
		break
	}

	db := base.DB.Raw(cntSql)
	_ = db.Row().Scan(&total)

	db = base.DB.Raw(infoSql)
	rows, err := db.Rows()
	if err != nil {
		log.Error("获取商品信息失败: ", err)
		return
	}
	for rows.Next() {
		var info model.GoodsInfo
		var gID int
		_ = rows.Scan(&info.Id, &info.Name, &info.Picture, &info.Price,
			&gID, &info.Number)
		goodsInfo = append(goodsInfo, info)
	}
	_ = rows.Close()

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
	db := base.DB.Raw("select id,address,linkman,phone,is_default from address where user_id=? and is_del=?", id, 0)
	rows, err := db.Rows()
	if err != nil {
		log.Error("获取收货地址失败: ", err)
		return
	}
	for rows.Next() {
		var defaultAddr int
		var addr model.Address
		_ = rows.Scan(&addr.Id, &addr.Address, &addr.Linkman, &addr.Phone, &defaultAddr)
		if defaultAddr == 1 {
			addr.DefaultAddr = true
		} else {
			addr.DefaultAddr = false
		}
		addresses = append(addresses, addr)
	}
	_ = rows.Close()

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
	db := base.DB.Exec("update address set address=?,linkman=?,phone=?,is_default=? where id=?",
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

	db := base.DB.Exec("insert into address set user_id=?,address=?,linkman=?,phone=?,is_default=?",
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
	db := base.DB.Exec("update address set is_del=? where id=?", 1, id)
	fmt.Println("删除: ", db.RowsAffected)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "删除成功!",
	})
}

// 关注 1   喜爱 2   浏览 3
func (g *Gin) Record(c *gin.Context) {
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
	title, _ := strconv.Atoi(c.Query("title"))
	var sql string
	switch title {
	case 1:
		sql = fmt.Sprintf(`select * from 
					(select id,g_id,browse_time from record where user_id=%d and record_type=1) as a 
					left join 
					(select id,g_name,g_nutrient,g_picture,g_price,content from goods) as b 
					on a.g_id=b.id`, id)
	case 2:
		// 喜爱
		sql = fmt.Sprintf(`select * from 
					(select id,g_id,browse_time from record where user_id=%d and record_type=2) as a 
					left join 
					(select id,g_name,g_nutrient,g_picture,g_price,content from goods) as b 
					on a.g_id=b.id`, id)
	case 3:
		// 浏览记录 3
		sql = fmt.Sprintf(`select * from 
					(select id,g_id,browse_time from record where user_id=%d and record_type=3 ) as a 
					left join 
					(select id,g_name,g_nutrient,g_picture,g_price,content from goods) as b 
					on a.g_id=b.id`, id)
	default:
		break
	}
	var goods []model.GoodsInfo
	db := base.DB.Raw(sql)
	rows, err := db.Rows()
	if err != nil {
		log.Error("查询失败: ", err)
		return
	}

	for rows.Next() {
		var g model.GoodsInfo
		var rId, gId, bTm int
		var nutrient string
		_ = rows.Scan(&rId, &gId, &bTm,
			&g.Id, &g.Name, &nutrient, &g.Picture, &g.Price, &g.Content)
		_ = json.Unmarshal([]byte(nutrient), &g.Nutrient)
		g.BrowseTime = time.Unix(int64(bTm), 0).Format("2006-01-02 15:04:05")
		goods = append(goods, g)
	}
	_ = rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    goods,
		"message": "获取数据成功!",
	})
}

// 红包卡券
func (g *Gin) Packets(c *gin.Context) {
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

	var packets []model.Packets
	db := base.DB.Raw("select id,name,picture,content,valuable,deadline from packets where user_id=? and is_overdue=? and is_used=?",
		id, 0, 0)
	rows, err := db.Rows()
	if err != nil {
		log.Error("查询信息失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    err,
			"message": "获取信息失败!",
		})
		return
	}

	for rows.Next() {
		var p model.Packets
		var deadline int
		_ = rows.Scan(&p.Id, &p.Name, &p.Picture, &p.Content, &p.Valuable, &deadline)
		p.Deadline = time.Unix(int64(deadline), 0).Format(strTm)
		packets = append(packets, p)
	}
	_ = rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    packets,
		"message": "获取信息成功!",
	})
}

// 我的钱包
