# Viper 配置中心项目

> 基于 Go 语言 + Viper 配置管理，实现多环境配置加载，支持 MySQL、Redis、Elasticsearch、MongoDB、RabbitMQ 等组件初始化，并提供JWT认证中间件。

## 📦 功能特性

- ✅ 统一的配置管理（Viper）
- ✅ 支持多种存储服务：
  - MySQL（关系型数据库）
  - Redis（缓存与分布式锁）
  - Elasticsearch（全文搜索）
  - MongoDB（文档数据库）
  - RabbitMQ（消息队列）
- ✅ JWT认证中间件
- ✅ 按需初始化服务
- ✅ 多环境配置（dev、test、prod）
- ✅ 结构化的目录设计

## 📂 目录结构

```
.
├── config/                # 配置结构体定义
│   ├── config.go         # 配置项定义
│   └── global.go         # 全局变量
│
├── inits/                # 各组件初始化
│   ├── config.go         # 读取配置
│   ├── es.go             # Elasticsearch 初始化
│   ├── init.go           # 总入口初始化
│   ├── mongodb.go        # MongoDB 初始化
│   ├── mysql.go          # MySQL 初始化
│   └── rabbitmq.go       # RabbitMQ 初始化
│   └── redis.go          # Redis 初始化
│
├── middleware/           # 中间件
│   ├── jwt.go            # JWT认证中间件
│   └── README.md         # JWT中间件使用说明
│
├── pkg/                  # 通用工具包
│   ├── alipay.go         # 支付宝支付
│   ├── aliyun.go         # 阿里云服务
│   ├── md5.go            # MD5加密
│   └── oss.go            # 对象存储
│
├── untils/               # 工具函数
│   ├── es.go             # ES工具函数
│   ├── pay.go            # 支付相关工具
│   └── sms.go            # 短信相关工具
│
├── dev.yaml              # 开发环境配置文件
├── go.mod                # Go模块定义
├── go.sum                # 依赖版本锁定
├── main.go               # 程序入口
└── README.md             # 项目说明
```

## ⚙️ 配置文件示例（`dev.yaml`）

```yaml
Mysql:
  Host: "14.103.175.131"
  Port: 3306
  User: "root"
  Password: "34862bca5fdc07f8d93ccb39994ad26d"
  Database: "health_db"
Redis:
  Host: "14.103.175.131"
  Port: 6379
  Password: "34862bca5fdc07f8d93ccb39994ad26d"
  Database: 0
Elasticsearch:
  Host: "14.103.175.131"
  Port: 9200
MongoDB:
  Host: "14.103.175.131"
  Port: 27017
  User: "gaoshaopu"
  Password: "gaoshaopu"
  Database: "health_db"
RabbitMQ:
  Host: "14.103.175.131"
  Port: 5672
  User: "2301a"
  Database: "2301a"
```

## 🚀 快速开始

### 安装

```bash
# 克隆项目
git clone https://github.com/zhanghanchen1014/viper.git

# 进入项目目录
cd viper

# 安装依赖
go mod tidy
```

### 配置

1. 复制 `dev.yaml` 为 `dev.yaml.local`，并根据实际环境修改配置
2. 修改 `inits/config.go` 中的配置文件路径

### 使用

```go
package main

import (
    _ "github.com/zhanghanchen1014/pubilc_pkg/inits" // 自动初始化所有组件
    "github.com/zhanghanchen1014/pubilc_pkg/config"
)

func main() {
    // 使用已初始化的组件
    // MySQL
    config.DB.Create(&YourModel{})
    
    // Redis
    config.Rdb.Set(config.Ctx, "key", "value", 0)
    
    // Elasticsearch
    config.Els.Info()
    
    // 更多组件使用方法...
}
```

## 🔐 JWT中间件使用

### 生成Token

```go
jwt := middleware.NewJWT()
token, err := jwt.CreateToken(userID, username)
```

### 验证Token

```go
// 在路由组中使用JWT中间件
authGroup := router.Group("/api")
authGroup.Use(middleware.JWTAuth())
```

### 获取用户信息

```go
// 在处理函数中获取用户信息
userID, _ := middleware.GetUserID(c)
username, _ := middleware.GetUsername(c)
```

更多详细用法请参考 `middleware/README.md` 和 `examples/jwt_example.go`。

## 📚 组件使用示例

### MySQL

```go
// 创建记录
config.DB.Create(&User{Name: "张三"})

// 查询记录
var user User
config.DB.First(&user, 1) // 查询ID为1的用户
```

### Redis

```go
// 设置键值
config.Rdb.Set(config.Ctx, "key", "value", time.Hour)

// 获取键值
val, err := config.Rdb.Get(config.Ctx, "key").Result()
```

### Elasticsearch

```go
// 搜索文档
res, err := config.Els.Search(
    config.Els.Search.WithIndex("your_index"),
    config.Els.Search.WithBody(strings.NewReader(`{"query":{"match":{"title":"搜索关键词"}}}`)),
)
```

### MongoDB

```go
// 获取集合
collection := MongoDBInit()

// 插入文档
collection.InsertOne(context.Background(), bson.D{{"name", "张三"}, {"age", 30}})
```

## 📄 许可证

[MIT](LICENSE)

