package handler

import (
	"app/m/base"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

// 一些预处理或者需要协程处理的函数

// 定期改变商品 推荐、新品状态
func ChangeNews() {
	var cTm string
	var id, isNew int
	db := base.DB.Raw("select id,create_time,is_new from goods")
	rows, err := db.Rows()
	if err != nil {
		log.Error("查询失败: ", err)
		return
	}
	for rows.Next() {
		tm := time.Now()
		_ = rows.Scan(&id, &cTm, &isNew)
		t, _ := time.Parse(strTm, cTm)
		if tm.Sub(t) > 7*24 && isNew == 1 {
			db = base.DB.Exec("update goods set is_new=? where id=?", 0, id)
			fmt.Println("Is new: ", db.RowsAffected)
		}
	}
	_ = rows.Close()
}

func ChangeRec() {
	var id, isRec int
	db := base.DB.Raw("select id from goods order by id desc limit 1")
	_ = db.Row().Scan(&id)
	res := int(float64(id) * 0.06)

	for i := 0; i < res; i++ {
		nId := rand.Intn(id)
		db = base.DB.Raw("select is_rec from goods where id=?", nId)
		_ = db.Row().Scan(&isRec)
		if isRec == 0 {
			db = base.DB.Exec("update goods set is_rec=? where id=?", 1, nId)
			fmt.Println("Is recommend: ", db.RowsAffected)
		}
	}
}

// 修改未使用的红包卡券的状态
func ChangePackets() {
	var id int
	var deadline string
	db := base.DB.Raw("select id,deadline from packets where is_overdue=0 and is_used=0")
	rows, err := db.Rows()
	if err != nil {
		log.Error("查询失败: ", err)
		return
	}
	for rows.Next() {
		_ = rows.Scan(&id, &deadline)
		tm := time.Now().Format(strTm)
		if tm >= deadline {
			db = base.DB.Exec("update packets set is_overdue=? where id=?", 1, id)
			fmt.Println("Is overdue: ", db.RowsAffected)
		}
	}
	_ = rows.Close()
}
