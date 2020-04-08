# Infant_mom_API
### APP后台接口

##### 以下接口不需要 token 认证

> no auth host: http://127.0.0.1:8080/v1/infant

**登录接口**
* url: /login
* method: post
* request:
    
    | 参数 | 类型 | 是否必须 | 备注 |
    | :-----| ----: | :----: | :----: |
    | phone | string | Y | 手机号码 |
    | password | string | Y | 密码(MD5加密) |

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
      
      user:
          
      | 参数 | 类型 |  备注 |
      | :-----| ----: | :----: |
      | id | int | 用户ID |
      | user_name | string | 用户名 |
      | mobile | string | 用户电话 |
      | head_portrait | string | 用户头像图片地址 |

##### 以下接口需要 token 认证

> with auth host: http://127.0.0.1:8080/v1/infant/auth