package main

import (
	"WEBOOK/internal/respository"
	"WEBOOK/internal/respository/dao"
	"WEBOOK/internal/service"
	"WEBOOK/internal/web"
	"WEBOOK/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	server := initWebService()
	db := initDB()
	initUserHandler(db, server)
	server.Run(":8080")
}

func initUserHandler(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDAO(db)
	ur := respository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	u := web.NewUserHandler(us)
	u.UserRegister(server)
}

func initWebService() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		/*		AllowOrigins:     []string{"https://foo.com"},
				AllowMethods:     []string{"PUT", "PATCH"},
				ExposeHeaders:    []string{"Content-Length"},*/
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}),
		func(ctx *gin.Context) {
			// 这里使用println函数输出一条信息，表明请求经过了这个自定义中间件。
			// 在实际应用中，可以在此处添加更复杂的逻辑，如日志记录、权限验证等。
			println("这是我的 Middleware")
		})

	login := &middleware.LoginMiddlewareBuilder{}
	// 存储数据的，也就是你 userId 存哪里
	// 直接存 cookie
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store), login.CheckLogin())
	return server

}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
