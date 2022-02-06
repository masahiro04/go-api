package controllers

import (
	uc "clean_architecture/golang/usecases"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type RouterHandler struct {
	ucHandler uc.Handler
	Logger    uc.Logger
}

func NewRouter(i uc.Handler) RouterHandler {
	return RouterHandler{
		ucHandler: i,
	}
}

func NewRouterWithLogger(i uc.Handler, logger uc.Logger) RouterHandler {
	return RouterHandler{
		ucHandler: i,
		Logger:    logger,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(rH.errorCatcher())

	rH.blogsRoutes(api)
	rH.usersRoutes(api)
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
	// blogs.PUT("/:id", rH.blogPatch)
	// blogs.PATCH("/:id", rH.blogPatch)
	// blogs.DELETE("/:id", rH.blogDelete)
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
		rH.Logger.Log(title, logs)
	}
}

func (RouterHandler) MethodAndPath(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
}
