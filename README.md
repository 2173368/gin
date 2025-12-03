# 个人博客系统后端

基于 Go 语言、Gin 框架和 GORM 开发的个人博客系统后端，支持完整的 CRUD 操作、用户认证和评论功能。

## 功能特性

- ✅ 用户注册和登录（JWT 认证）
- ✅ 博客文章的创建、读取、更新和删除
- ✅ 文章评论功能
- ✅ 用户权限控制（只能操作自己的内容和文章）
- ✅ SQLite 数据库存储
- ✅ 完整的错误处理和日志记录

## 技术栈

- **Go** - 编程语言
- **Gin** - Web 框架
- **GORM** - ORM 库
- **SQLite** - 数据库
- **JWT** - 用户认证
- **bcrypt** - 密码加密

## 项目结构

```
├── config/          # 配置文件
│   └── database.go  # 数据库配置
├── database/         # 数据库相关
│   └── migrate.go   # 数据库迁移
├── handlers/         # 请求处理程序
│   ├── auth.go      # 认证相关
│   ├── post.go      # 文章相关
│   └── comment.go   # 评论相关
├── middleware/       # 中间件
│   └── auth.go      # JWT 认证中间件
├── models/          # 数据模型
│   ├── user.go      # 用户模型
│   ├── post.go      # 文章模型
│   └── comment.go   # 评论模型
├── main.go          # 应用入口
└── README.md        # 项目说明
```

## 安装和运行

1. 确保已安装 Go 1.16+
2. 克隆项目并进入目录
3. 安装依赖：
   ```bash
   go mod tidy
   ```
4. 运行项目：
   ```bash
   go run main.go
   ```
5. 服务器将在 `http://localhost:8080` 启动

## API 接口文档

### 认证相关

#### 用户注册
- **POST** `/api/auth/register`
- 请求体：
  ```json
  {
    "username": "testuser",
    "password": "password123",
    "email": "test@example.com"
  }
  ```

#### 用户登录
- **POST** `/api/auth/login`
- 请求体：
  ```json
  {
    "username": "testuser",
    "password": "password123"
  }
  ```
- 响应：包含 JWT token

### 文章相关

#### 获取所有文章
- **GET** `/api/posts` - 公开接口

#### 获取单篇文章
- **GET** `/api/posts/:id` - 公开接口

#### 创建文章（需要认证）
- **POST** `/api/posts`
- 请求头：`Authorization: Bearer <JWT_TOKEN>`
- 请求体：
  ```json
  {
    "title": "文章标题",
    "content": "文章内容"
  }
  ```

#### 更新文章（需要认证）
- **PUT** `/api/posts/:id`
- 请求头：`Authorization: Bearer <JWT_TOKEN>`

#### 删除文章（需要认证）
- **DELETE** `/api/posts/:id`
- 请求头：`Authorization: Bearer <JWT_TOKEN>`

### 评论相关

#### 创建评论（需要认证）
- **POST** `/api/comments`
- 请求头：`Authorization: Bearer <JWT_TOKEN>`
- 请求体：
  ```json
  {
    "content": "评论内容",
    "post_id": 1
  }
  ```

#### 获取文章评论（需要认证）
- **GET** `/api/posts/:post_id/comments`
- 请求头：`Authorization: Bearer <JWT_TOKEN>`

#### 删除评论（需要认证）
- **DELETE** `/api/comments/:id`
- 请求头：`Authorization: Bearer <JWT_TOKEN>`

### 健康检查

- **GET** `/health` - 检查服务状态

## 数据库设计

### Users 表
- id (主键)
- username (唯一)
- password (加密存储)
- email (唯一)
- created_at
- updated_at

### Posts 表
- id (主键)
- title
- content
- user_id (外键)
- created_at
- updated_at

### Comments 表
- id (主键)
- content
- user_id (外键)
- post_id (外键)
- created_at

## 安全特性

- 密码使用 bcrypt 加密存储
- JWT token 用于用户认证和授权
- 用户只能操作自己的内容和文章
- 输入验证和错误处理

## 开发说明

1. 数据库使用 SQLite，数据文件为 `blog.db`
2. JWT 密钥在 `handlers/auth.go` 中配置
3. 所有 API 返回标准 JSON 格式
4. 错误响应包含详细的错误信息

## 后续改进

- [ ] 添加文章分类和标签
- [ ] 支持文章分页
- [ ] 添加用户头像上传
- [ ] 实现文章搜索功能
- [ ] 添加邮件通知
- [ ] 支持多种数据库（MySQL/PostgreSQL）
- [ ] 添加单元测试
- [ ] 实现 API 版本控制
- [ ] 添加 Swagger 文档
- [ ] 支持 Docker 部署