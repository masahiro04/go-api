package controllers

import (
	"fmt"
	"go-api/adapters"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

// driverという名前でも良いかもj
// Repositoryを返すプログラムでも良さそう
type RouterHandler struct {
	drivers *adapters.Driver
}

// NOTE(okubo): usecasesのinterfaceを適用することで、抽象→実装
// func NewRouter(driver *adapters.Driver) RouterHandler {
func NewRouter(driver *adapters.Driver) RouterHandler {
	return RouterHandler{
		drivers: driver,
	}
}

// func NewRouterWithLogger(
// 	i uc.Handler, fb uc.FirebaseHandler, logger uc.Logger) RouterHandler {
// 	return RouterHandler{
// 		ucHandler:       i,
// 		firebaseHandler: fb,
// 		Logger:          logger,
// 	}
// }

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(rH.errorCatcher())

	rH.signupRoutes(api)
	rH.blogsRoutes(api)
	rH.usersRoutes(api)
}

func (rH RouterHandler) signupRoutes(api *gin.RouterGroup) {
	signup := api.Group("/signup")
	signup.POST("", rH.signUp)
	// signup.POST("/complete", rH.cookieMiddleware(), rH.signUpComplete)
}

func (rH RouterHandler) blogsRoutes(api *gin.RouterGroup) {
	blogs := api.Group("/blogs")
	blogs.GET("", rH.blogsGetAll)
	blogs.GET("/:id", rH.blogGet)
	blogs.POST("", rH.blogPost)
	blogs.PUT("/:id", rH.blogPatch)
	blogs.PATCH("/:id", rH.blogPatch)
	blogs.DELETE("/:id", rH.blogDelete)
}

func (rH RouterHandler) usersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	users.GET("", rH.userGetAll)
	users.GET("/:id", rH.userGet)
	users.POST("", rH.userPost)
	users.PUT("/:id", rH.userPatch)
	users.PATCH("/:id", rH.userPatch)
	users.DELETE("/:id", rH.userDelete)
}

func (rH RouterHandler) errorCatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() > 399 {
			var errorMessages []map[string]string
			for _, m := range c.Errors.Errors() {
				message := make(map[string]string, len(c.Errors.Errors()))
				message["message"] = m
				errorMessages = append(errorMessages, message)
			}

			c.Render(
				c.Writer.Status(),
				render.JSON{
					Data: gin.H{
						"errors": errorMessages,
					},
				},
			)
		}
	}
}

// log is used to "partially apply" the title to the rH.logger.Log function
// so we can see in the logs from which route the log comes from
func (rH RouterHandler) log(title string) func(...interface{}) {
	return func(logs ...interface{}) {
		// rH.Logger.Log(title, logs)
	}
}

func (RouterHandler) MethodAndPath(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
}
