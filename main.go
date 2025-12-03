package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"gin/config"
	"gin/database"
	"gin/handlers"
	"gin/middleware"
)

func main() {
	// 初始化日志
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 初始化数据库
	config.InitDB()
	database.Migrate()

	// 创建 Gin 路由
	r := gin.Default()

	// 添加中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 公共路由
	api := r.Group("/api")
	{
		// 认证路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// 文章路由（公开）
		posts := api.Group("/posts")
		{
			posts.GET("", handlers.GetPosts)
			posts.GET("/:id", handlers.GetPost)
		}

		// 需要认证的路由
		authorized := api.Group("")
		authorized.Use(middleware.Auth())
		{
			// 文章管理
			authorized.POST("/posts", handlers.CreatePost)
			authorized.PUT("/posts/:id", handlers.UpdatePost)
			authorized.DELETE("/posts/:id", handlers.DeletePost)

			// 评论管理
			authorized.POST("/comments", handlers.CreateComment)
			authorized.DELETE("/comments/:id", handlers.DeleteComment)

			// 获取文章评论（需要认证）
			authorized.GET("/posts/:id/comments", handlers.GetPostComments)
		}
	}

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Blog API is running",
		})
	})

	// 启动服务器
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
