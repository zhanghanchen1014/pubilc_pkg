# Viper 配置中心项目

> 基于 Go 语言 + Viper 配置管理，实现多环境配置加载，支持 MySQL、Redis、Elasticsearch、MongoDB、RabbitMQ 等组件初始化。

> 该配置信息所用的服务器是高绍朴的线上服务器 14.103.175.131

> 数据库统一为 "health_db"

---

## 📦 功能特性
- ✅ 统一的配置管理（Viper）
- ✅ 支持多种存储服务：
    - MySQL（关系型数据库）
    - Redis（缓存与分布式锁）
    - Elasticsearch（全文搜索）
    - MongoDB（文档数据库）
    - RabbitMQ（消息队列）
- ✅ 按需初始化服务
- ✅ 多环境配置（dev、test、prod）
- ✅ 结构化的目录设计

---

## 📂 目录结构 目录结构
    .
    ├── config/ # 配置结构体定义
    │ ├── config.go # 配置项定义
    │ └── global.go # 全局变量
    │
    ├── inits/ # 各组件初始化
    │ ├── config.go # 读取配置
    │ ├── es.go # Elasticsearch 初始化
    │ ├── init.go # 总入口初始化
    │ ├── mongodb.go # MongoDB 初始化
    │ ├── mysql.go # MySQL 初始化
    │ └── redis.go # Redis 初始化
    │
    ├── dev.yaml # 开发环境配置文件
    ├── go.mod
    ├── main.go # 程序入口
    └── README.md

---

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
 Host: "156724.103.175.131"
 Port: 5672
 User: "2301a"
 Database: "2301a"

