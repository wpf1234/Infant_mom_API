package model

type LoginRes struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type User struct {
	Id           int    `json:"id"`
	UserName     string `json:"user_name"`
	Mobile       string `json:"mobile"`
	HeadPortrait string `json:"head_portrait"`
}

// 我的页面
type MineRes struct {
	User      User `json:"user"`
	Attention int  `json:"attention"` // 关注
	Favorite  int  `json:"favorite"`  // 喜爱
	Packets   int  `json:"packets"`   // 红包卡券
	Browse    int  `json:"browse"`    // 浏览
	Pending   int  `json:"pending"`   // 待付款
	Received  int  `json:"received"`  // 待收货
	Evaluate  int  `json:"evaluate"`  // 待评价
	Refund    int  `json:"refund"`    // 退款/售后
}

type MyOrder struct {
	Info  []GoodsInfo `json:"info"`
	Total int         `json:"total"`
}

type GoodsInfo struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Nutrient interface{} `json:"nutrient"`
	Picture  string      `json:"picture"`
	Price    float64     `json:"price"`
	Number   int         `json:"number"`
}

type Address struct {
	Id          int    `json:"id"`
	Address     string `json:"address"`
	Linkman     string `json:"linkman"`
	Phone       string `json:"phone"`
	DefaultAddr bool   `json:"default_addr"`
}

// 商品详情
type GoodsDetails struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Nutrient interface{} `json:"nutrient"`
	Picture  string      `json:"picture"`
	Price    float64     `json:"price"`
}

type GoodsRes struct {
	Details   GoodsDetails `json:"details"`
	Attention bool         `json:"attention"`
	Favorite  bool         `json:"favorite"`
}
