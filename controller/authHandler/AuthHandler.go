package authhandler

import (
	// "net/http"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	return
	// 尝试从 cookie 中获取 JWT
	// tokenCookie, err := c.Cookie("token")
	// if err != nil {
	// 	// 如果获取失败，说明未登录或登录已过期
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
	// 	return
	// }
	// // 尝试从 Redis 缓存中获取 JWT
	// tokenFromCache, err := redisClient.Get(claims.Username).Result()
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	// 	return
	// }

	// // 验证 JWT 是否正确
	// token, err := jwt.ParseWithClaims(tokenFromCache, &Claims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return jwtSecret, nil
	// })
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	// 	return
	// }

	// // 获取 JWT 的负载
	// claims, ok := token.Claims.(*Claims)
	// if !ok || !token.Valid {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	// 	return
	// }

	// // 将用户名存入上下文
	// c.Set("username", claims.Username)

	// // 继续处理请求
	// c.Next()
}
