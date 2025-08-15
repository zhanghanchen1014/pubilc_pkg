# JWT中间件使用说明

## 简介

本中间件提供了JWT（JSON Web Token）的完整实现，包括token的生成、验证、刷新等功能，适用于基于Gin框架的Go应用程序。

## 功能特性

- 生成JWT Token
- 验证JWT Token
- 刷新JWT Token
- 从上下文中获取用户信息
- 自定义错误处理

## 配置说明

在项目的配置文件中添加JWT相关配置：

```yaml
JWT:
  Secret: "your-secret-key-here"  # JWT密钥，建议使用强随机字符串
  ExpireTime: 86400               # Token过期时间（秒），默认为24小时
  Issuer: "your-application-name" # Token签发者
```

## 使用方法

### 1. 生成Token

```go
// 创建JWT实例
jwt := middleware.NewJWT()

// 生成Token
userID := uint(1)       // 用户ID
username := "testuser" // 用户名
token, err := jwt.CreateToken(userID, username)
if err != nil {
    // 处理错误
}
```

### 2. 验证Token

在需要JWT认证的路由组中使用中间件：

```go
// 创建需要认证的路由组
authGroup := router.Group("/api")

// 使用JWT中间件
authGroup.Use(middleware.JWTAuth())

// 添加受保护的路由
authGroup.GET("/user/info", UserInfoHandler)
```

### 3. 获取用户信息

在处理函数中获取用户信息：

```go
func UserInfoHandler(c *gin.Context) {
    // 获取用户ID
    userID, exists := middleware.GetUserID(c)
    if !exists {
        // 处理错误
        return
    }
    
    // 获取用户名
    username, exists := middleware.GetUsername(c)
    if !exists {
        // 处理错误
        return
    }
    
    // 获取完整的用户信息
    claims, exists := middleware.GetUserInfo(c)
    if !exists {
        // 处理错误
        return
    }
    
    // 使用用户信息...
}
```

### 4. 刷新Token

```go
// 创建JWT实例
jwt := middleware.NewJWT()

// 刷新Token
newToken, err := jwt.RefreshToken(oldToken)
if err != nil {
    // 处理错误
}
```

## 客户端使用说明

### 1. 发送请求时携带Token

在HTTP请求头中添加Authorization字段，格式为：

```
Authorization: Bearer <your-token-here>
```

### 2. 处理Token过期

当收到401状态码且错误信息为"token已过期"时，应调用刷新Token接口获取新的Token。

## 完整示例

请参考 `examples/jwt_example.go` 文件中的完整示例。