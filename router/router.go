/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2022-12-21 16:32:01
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-31 23:31:31
 * @FilePath: /goblog-back/router/router.go
 * @Description: 路由文件
 */
package router

import (
	"goblog-end/config"
	"goblog-end/controller"
	authHandler "goblog-end/controller/authHandler"
	controllerLogin "goblog-end/controller/login"

	"github.com/gin-gonic/gin"

	"net/http"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func Start() {
	r := gin.Default()

	//开启中间件 允许使用跨域请求
	r.Use(config.Cors())

	// 用户注册、登录、个人信息 API 放到 /api/v1/users 分组中
	users := r.Group("/api/v1/users")
	{
		// users.POST("/register", registerHandler)
		users.POST("/login", controllerLogin.Login)

		// 需要认证的 API 放到 authGroup 中
		authGroup := users.Group("/")
		authGroup.Use(authHandler.AuthHandler)
		{
			// authGroup.GET("/me", getProfileHandler)
			// authGroup.PUT("/me", updateProfileHandler)
		}
	}

	// // 文章增删改查 API 放到 /api/v1/articles 分组中
	// articles := r.Group("/api/v1/articles")
	// {
	//     // 需要认证的 API 放到 authGroup 中
	//     authGroup := articles.Group("/")
	//     authGroup.Use(authMiddleware)
	//     {
	//         authGroup.GET("/", listArticlesHandler)
	//         authGroup.POST("/", createArticleHandler)
	//         authGroup.GET("/:id", getArticleHandler)
	//         authGroup.PUT("/:id", updateArticleHandler)
	//         authGroup.DELETE("/:id", deleteArticleHandler)
	//     }
	// }

	// 处理需要认证的请求
	// authGroup.GET("/profile", profileHandler)

	// 文章增删改查相关的API
	// router.GET("/articles", handleListArticles)
	// router.GET("/articles/:id", handleGetArticle)
	// router.POST("/articles", handleCreateArticle)
	// router.PUT("/articles/:id", handleUpdateArticle)
	// router.DELETE("/articles/:id", handleDeleteArticle)

	// 搜索文章
	// router.GET("/articles/search", handleSearchArticles)

	// 管理员系统相关的API
	// router.GET("/admin/articles", handleListAllArticles)
	// router.PUT("/admin/articles/:id", handleUpdateArticleAsAdmin)
	// router.DELETE("/admin/articles/:id", handleDeleteArticleAsAdmin)

	// 个人信息管理相关的API
	// router.GET("/profile", handleGetProfile)
	// router.PUT("/profile", handleUpdateProfile)

	// 创建评论
	// router.POST("/articles/:id/comments", handleCreateComment)

	// 获取文章的所有评论
	// router.GET("/articles/:id/comments", handleListComments)

	// 删除评论
	// router.DELETE("/articles/:id/comments/:commentId", handleDeleteComment)

	// 创建文章分类
	// router.POST("/categories", handleCreateCategory)

	// 获取所有文章分类
	// router.GET("/categories", handleListCategories)

	// 创建文章标签
	// router.POST("/tags", handleCreateTag)

	// 获取所有文章标签
	// router.GET("/tags", handleListTags)

	// 将文章添加到分类中
	// router.PUT("/articles/:id/categories/:categoryId", handleAddArticleToCategory)

	// 将文章从分类中移除
	// router.DELETE("/articles/:id/categories/:categoryId", handleRemoveArticleFromCategory)

	// 将文章添加到标签中
	// router.PUT("/articles/:id/tags/:tagId", handleAddArticleToTag)

	// 将文章从标签中移除
	// router.DELETE("/articles/:id/tags/:tagId", handleRemoveArticleFromTag)

	// 创建用户角色
	// router.POST("/roles", handleCreateRole)

	// 获取所有用户角色
	// router.GET("/roles", handleListRoles)

	// 创建用户
	// router.POST("/users", handleCreateUser)

	// 获取所有用户
	// router.GET("/users", handleListUsers)

	// 获取单个用户信息
	// router.GET("/users/:id", handleGetUser)

	// 更新用户信息
	// router.PUT("/users/:id", handleUpdateUser)

	// 删除用户
	// router.DELETE("/users/:id", handleDeleteUser)

	// 分配用户角色
	// router.PUT("/users/:id/roles/:roleId", handleAssignUserRole)

	// 撤销用户角色
	// router.DELETE("/users/:id/roles/:roleId", handleRevokeUserRole)

	r.GET("/page/category/", controller.Category)

	r.GET("/page/userinfo/:name", controller.PageUserInfo)

	r.GET("/post", controller.Post)
	r.GET("/post/:pid", controller.PostByPid)
	r.GET("/post/id", controller.PostByPid)

	// r.POST("/login/", controller.Login)

	r.POST("/writing", controller.Writing)

	r.POST("/api/upload/img", controller.Upload)
	r.GET("/api/qiniu", controller.QiniuToken)

	port := "8082"
	r.Run(":" + port)
}
