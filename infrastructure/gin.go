package infrastructure

import (
	"log"
	"strconv"
	"time"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinServerMode int

const (
	DebugMode GinServerMode = iota
	ReleaseMode
	TestMode
)

// GinServer : the struct gathering all the server details
type GinServer struct {
	port   int
	Router *gin.Engine
}

func NewServer(port int, mode GinServerMode) GinServer {
	s := GinServer{}
	s.port = port

	s.Router = gin.New()

	switch mode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	case TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	s.Router.Use(gin.Recovery())

	SetCors(s.Router)

	return s
}

// SetCors is a helper to set current engine cors
func SetCors(engine *gin.Engine) {
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           50 * time.Second,
	}))

}

// Start the server
func (s GinServer) Start() {
	if err := s.Router.Run(":" + strconv.Itoa(int(s.port))); err != nil {
		log.Print(err)
	}
}
