# Infant_mom_API
### APP后台接口

##### 以下接口不需要 token 认证

> 请求地址: http://127.0.0.1:8080/v1/infant

> GoodsDetails :

 | 参数 | 类型 |  备注 |
 | :-----| ----: | :----: |
 | id | int | 用户ID |
 |name |    string   |  商品名称 |
 |nutrient | map[string]string | 营养元素含量|
 |picture | string | 商品图片 |
 |price  |  float64 | 商品价格 |
 |classify | string | 类别 |
 |content | string  | 描述 |
     

**登录接口**
* url:      /login
* method:   POST
* request:
    
    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | phone | string | Y | 手机号码 |
    | password | string | Y | 密码(MD5加密) |

* 请求示例：
```json
{
	"phone":"13800000000",
	"password":"test123456"
}
```

* response: 

    | 参数 | 类型 |  备注 |
    | :-----| ----: | :----: |
    | code | int | 状态码(200/500/400) |
    | data | struct | 返回的数据集(一个结构体) |
    | message | string | 提示信息 |
    
    > data: 
       
     | 参数 | 类型 |  备注 |
     | :-----| ----: | :----: |
     | user | struct | 用户信息 |
     | token | string | Token值 |
      
     > user:
          
     | 参数 | 类型 |  备注 |
     | :-----| ----: | :----: |
     | id | int | 用户ID |
     | user_name | string | 用户名 |
     | mobile | string | 用户电话 |
     | head_portrait | string | 用户头像图片地址 |
     
* 返回实例：
````json
{
    "code": 200,
    "data": {
        "user": {
            "id": 1,
            "user_name": "laofee",
            "mobile": "13800000000",
            "head_portrait": "static/head/github头像-1.gif"
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0Ijoic3RhdGljL2hlYWQvZ2l0aHVi5aS05YOPLTEuZ2lmIiwiZXhwIjoxNTg3MDI2ODQ2LCJuYmYiOjE1ODY0MjEwNDZ9.yEHU8n41Ug8N8afopW1FqWAf2mrMJC8y8Ts2qPa-ffI"
    },
    "message": "登录成功!"
}
````
     
**注册接口**
* url:      /register
* method:   POST
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | user_name | string | Y | 用户名 |
    | password | string | Y | 密码(MD5加密) |
    | phone | string | Y | 手机号码 |
    
* 请求示例：
```json
{
	"user_name":"test",
	"password":"qwer1234",
	"phone":"12345678901"
}
```
   
* response: 

     | 参数 | 类型 | 备注 |
     | :-----| ----: | :----: | 
     | code | int | 状态码 |
     | data | string | 返回的数据 |
     | message | string | 提示信息 |
    
* 返回实例:
````json
{
    "code": 200,
    "data": null,
    "message": "注册成功!"
}
````

**首页**
* url:      /home
* method:   GET
* request: 无需请求参数
* response: 
    
     | 参数 | 类型 | 备注 |
     | :-----| ----: | :----: | 
     | code | int | 状态码 |
     | data | struct | 返回的数据 |
     | message | string | 提示信息 |
     
     > data:
  
    | 参数 | 类型 | 备注 |
    | :-----| ----: | :----: | 
    | recommend | []GoodsDetails | 推荐商品 |
    | new_goods | []GoodsDetails | 返回的数据 |
    | others | []GoodsDetails | 提示信息 |  

* 返回实例:
```json
{
    "code": 200,
    "data": {
        "recommend": [
            {
                "id": 4,
                "name": "文化街",
                "nutrient": {
                    "元素A": "1.8%",
                    "元素B": "13%",
                    "元素C": "0.38mg/g"
                },
                "picture": "",
                "price": 6.2,
                "classify": "C4",
                "content": "他雅虎日本的风格与爱放火"
            }
        ],
        "new_goods": null,
        "others": [
            {
                "id": 1,
                "name": "商品",
                "nutrient": null,
                "picture": "",
                "price": 2.5,
                "classify": "C1",
                "content": "描述描述描述描述描述"
            },
            {
                "id": 2,
                "name": "蔬菜哦",
                "nutrient": {
                    "胡萝卜素": "1.5%"
                },
                "picture": "",
                "price": 3,
                "classify": "C2",
                "content": "我无法测天气文风和任务一万二"
            },
            {
                "id": 3,
                "name": "水果盘",
                "nutrient": {
                    "Ca": "30%",
                    "VC": "16%"
                },
                "picture": "",
                "price": 5.2,
                "classify": "C3",
                "content": "和忽悠人法规和更好看个人通过五热管"
            }
        ]
    },
    "message": "获取商品信息成功!"
}
```

**查询接口**
* url:      /search?text=商品名称&type=1
* method:   GET
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | text | string | Y | 需要查询的内容(名称或类别或营养元素) |
    | type | number | Y | 查询类别(默认为1，即名称查询；2:类别;3:营养元素) |
* response: 

     | 参数 | 类型 | 备注 |
     | :-----| ----: | :----: | 
     | code | int | 状态码 |
     | data | GoodsDetails | 返回的数据 |
     | message | string | 提示信息 |

* 返回示例:
```json
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "name": "商品",
            "nutrient": {
                "A": "16%",
                "Fe": "30%"
            },
            "picture": "",
            "price": 2.5,
            "classify": "C1",
            "content": "描述描述描述描述描述"
        }
    ],
    "message": "查询成功!"
}
```

##### 以下接口需要 token 认证

> 需要认证的接口地址: http://127.0.0.1:8080/v1/infant/auth

**上传头像**
* url:      /upload?file=abcd.gif
* method:   POST
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | file | string | Y | 头像图片地址 |

* response: 

     | 参数 | 类型 | 备注 |
     | :-----| ----: | :----: | 
     | code | int | 状态码 |
     | data | string | 返回的数据 |
     | message | string | 提示信息 |
     
* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "上传成功!"
}
```

**刷新Token**
* url:      /refresh
* method:   GET
* request: 无需参数，如果用户已登录，每次进入项目时自动调用该接口

* response: 

     | 参数 | 类型 | 备注 |
     | :-----| ----: | :----: | 
     | code | int | 状态码 |
     | data | string | 返回的数据 |
     | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "刷新成功!"
}
```

**我的首页**
* url:      /mine
* method:   GET
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 无需参数

* response: 

     | 参数 | 类型 | 备注 |
     | :-----| ----: | :----: | 
     | code | int | 状态码 |
     | data | struct | 返回的数据 |
     | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": {
        "user": {
            "id": 1,
            "user_name": "laofee",
            "mobile": "13800000000",
            "head_portrait": "./static/head/github头像-1.gif"
        },
        "attention": 3,   // 关注的商品
        "favorite": 2,    // 喜爱的商品
        "packets": 2,     // 红包卡券
        "browse": 3,      // 浏览记录
        "pending": 0,     // 待付款
        "received": 0,    // 待收货
        "evaluate": 0,    // 待评价
        "refund": 0       // 退款/售后
    },
    "message": "获取信息成功!"
}
```

**我的订单**
* url:      /mine/order?title=all
* method:   GET
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | title | string | Y | 查看类型(all：全部；pending：待付款；received：待收货；evaluate：待评价；refund：退款/售后) |

* response: 

     | 参数 | 类型 | 备注 |
     | :-----| ----: | :----: | 
     | code | int | 状态码 |
     | data | struct | 返回的数据 |
     | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": {
        "info": [
            {
                "id": 1,
                "name": "商品",
                "nutrient": null,
                "picture": "",
                "price": 2.5,
                "number": 1,
                "content": "",
                "browse_time": ""
            },
            {
                "id": 3,
                "name": "水果盘",
                "nutrient": null,
                "picture": "",
                "price": 5.2,
                "number": 2,
                "content": "",
                "browse_time": ""
            },
            {
                "id": 2,
                "name": "蔬菜哦",
                "nutrient": null,
                "picture": "",
                "price": 3,
                "number": 0,
                "content": "",
                "browse_time": ""
            },
            {
                "id": 4,
                "name": "文化街",
                "nutrient": null,
                "picture": "",
                "price": 6.2,
                "number": 0,
                "content": "",
                "browse_time": ""
            }
        ],
        "total": 4
    },
    "message": "获取订单信息成功!"
}
```

**收货地址**
###### 查询
* url:      /mine/address
* method:   GET
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 无需参数

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | struct | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "address": "呃唔翻噶似的九月份",
            "linkman": "王鹏飞",
            "phone": "13800000000",
            "default_addr": true
        },
        {
            "id": 2,
            "address": "发谁额个发uu复合物i哦额",
            "linkman": "王鹏飞",
            "phone": "13800000000",
            "default_addr": false
        }
    ],
    "message": "获取收货地址成功!"
}
```

###### 新增
* url:      /mine/address
* method:   POST
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    |address|     string| Y| 地址信息 |
    |linkman |    string |Y| 联系人 |
    |phone    |   string |Y| 电话 |
    |default_addr| bool | Y| 是否设为默认地址 |

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | string | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "新增成功!"
}
```

###### 修改
* url:      /mine/address
* method:   PUT
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    |address|     string| Y| 地址信息 |
    |linkman |    string |Y| 联系人 |
    |phone    |   string |Y| 电话 |
    |default_addr| bool | Y| 是否设为默认地址 |

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | string | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "修改成功!"
}
```

###### 删除
* url:      /mine/address?id=1
* method:   DELETE
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | id | number| Y| 收货地址ID |

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | string | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "删除成功!"
}
```

**查看 关注/喜爱/浏览 记录**
* url:      /mine/record?title=1
* method:   GET
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | title | number| Y| 查看的类型(关注 1   喜爱 2   浏览 3) |

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | struct | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "name": "商品",
            "nutrient": {
                "A": "16%",
                "Fe": "30%"
            },
            "picture": "",
            "price": 2.5,
            "number": 0,
            "content": "描述描述描述描述描述",
            "browse_time": "2020-04-09 11:13:41"
        },
        {
            "id": 2,
            "name": "蔬菜哦",
            "nutrient": {
                "胡萝卜素": "1.5%"
            },
            "picture": "",
            "price": 3,
            "number": 0,
            "content": "我无法测天气文风和任务一万二",
            "browse_time": "2020-04-09 11:13:41"
        },
        {
            "id": 3,
            "name": "水果盘",
            "nutrient": {
                "Ca": "30%",
                "VC": "16%"
            },
            "picture": "",
            "price": 5.2,
            "number": 0,
            "content": "和忽悠人法规和更好看个人通过五热管",
            "browse_time": "2020-04-09 11:13:41"
        }
    ],
    "message": "获取数据成功!"
}
```

**查看我的红包卡券**
* url:      /mine/packets
* method:   GET
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 无需参数

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | struct | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "name": "通用券",
            "picture": "",
            "valuable": 0,
            "content": "",
            "deadline": "1970-01-01 08:00:00"
        },
        {
            "id": 2,
            "name": "专用券",
            "picture": "",
            "valuable": 0,
            "content": "",
            "deadline": "1970-01-01 08:00:00"
        }
    ],
    "message": "获取信息成功!"
}
```

**购物车**
###### 查询
* url:      /cart
* method:   GET
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 无需参数

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | struct | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "goods_name": "商品",
            "goods_picture": "",
            "goods_price": 2.5,
            "goods_content": "描述描述描述描述描述",
            "goods_num": 2
        },
        {
            "id": 2,
            "goods_name": "蔬菜哦",
            "goods_picture": "",
            "goods_price": 3,
            "goods_content": "我无法测天气文风和任务一万二",
            "goods_num": 1
        },
        {
            "id": 3,
            "goods_name": "水果盘",
            "goods_picture": "",
            "goods_price": 5.2,
            "goods_content": "和忽悠人法规和更好看个人通过五热管",
            "goods_num": 3
        }
    ],
    "message": "查询成功!"
}
```

###### 新增
* url:      /cart?goods=3
* method:   POST
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    |goods| number | Y| 需要添加的商品ID |
    |num | number | N | 商品数量(默认为1) |

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | string | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "加入购物车成功!"
}
```

###### 修改
* url:      /cart
* method:   PUT
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    |goods| number| Y| 商品ID |
    |num | number |Y| 商品数量 |

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | string | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "修改成功!"
}
```

###### 删除
* url:      /mine/address?id=1
* method:   DELETE
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | id | number| Y| 收货地址ID |

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | string | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "从购物车中成功移除!"
}
```

**修改 关注/喜爱/浏览 状态**
* url:      /change?goods=1&state=-1
* method:   PUT
* header:  
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | goods | number| Y| 商品ID |
    | state | number| Y| 商品状态(关注 1/取消 -1   喜爱 2/取消-2   浏览 3/删除 -3) |

* response: 
  
   | 参数 | 类型 | 备注 |
   | :-----| ----: | :----: | 
   | code | int | 状态码 |
   | data | struct | 返回的数据 |
   | message | string | 提示信息 |

* 返回示例：
```json
{
    "code": 200,
    "data": null,
    "message": "修改成功!"
}
```

> 特殊接口，既可以在无token值的情况下使用，也可以在有token验证的下使用

**商品详情**
* url:      http://127.0.0.1:8080/v1/infant/auth/detail?goods=1
            http://127.0.0.1:8080/v1/infant/detail?goods=1
* method:   GET
* header: 
```json
{
  "token":
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGFvZmVlIiwicGFzc3dvcmQiOiJ0ZXN0MTIzNDU2IiwibW9iaWxlIjoiMTM4MDAwMDAwMDAiLCJoZWFkX3BvcnRyYWl0IjoiLi9zdGF0aWMvaGVhZC9naXRodWLlpLTlg48tMS5naWYiLCJleHAiOjE1ODY5NTIxOTEsIm5iZiI6MTU4NjMyOTU1OX0.DiOUt2YiRWYpugs0UWXY9HBDd3xJafUGrvuQNMiKgV0"
}
```
* request: 

    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | goods | number | Y | 需要查看详情的商品ID |
    | user_id | number | N | 用户ID(用于记录浏览记录，没有登录则不需要该参数) |

* response: 

     | 参数 | 类型 | 备注 |
     | :-----| ----: | :----: | 
     | code | int | 状态码 |
     | data | GoodsDetails | 返回的数据 |
     | message | string | 提示信息 |

* 返回示例：
> 无token验证
```json 
{
    "code": 200,
    "data": {
        "details": {
            "id": 1,
            "name": "商品",
            "nutrient": {
                "A": "16%",
                "Fe": "30%"
            },
            "picture": "",
            "price": 2.5,
            "classify": "C1",
            "content": "描述描述描述描述描述"
        },
        "attention": false,  // 没有关注
        "favorite": false    // 不在喜欢的商品中
    },
    "message": "获取商品详情成功!"
}
```

> 有token验证
```json 
{
    "code": 200,
    "data": {
        "details": {
            "id": 1,
            "name": "商品",
            "nutrient": {
                "A": "16%",
                "Fe": "30%"
            },
            "picture": "",
            "price": 2.5,
            "classify": "C1",
            "content": "描述描述描述描述描述"
        },
        "attention": true,      // 关注状态 
        "favorite": true        // 喜欢状态
    },
    "message": "获取商品详情成功!"
}
```