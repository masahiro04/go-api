package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	dbTransaction "clean_architecture/golang/adapters/dao.dbTransaction"

	_ "github.com/lib/pq"

	"fmt"

	blogRW "clean_architecture/golang/adapters/dao.blogRW"
	server "clean_architecture/golang/adapters/gin.server"
	logger "clean_architecture/golang/adapters/logrus.logger"
	validator "clean_architecture/golang/adapters/validator"

	infra "clean_architecture/golang/infrastructure"
	uc "clean_architecture/golang/usecases"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// the command to run the server
var rootCmd = &cobra.Command{
	Use:   "hrbase-pro",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\n", Build, Version)
	},
}

func main() {
	log.SetFlags(log.Llongfile)
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Println(err)
	}

	rootCmd.AddCommand(versionCmd)
	cobra.OnInitialize(infra.CobraInitialization)

	infra.LoggerConfig(rootCmd)
	infra.ServerConfig(rootCmd)
	infra.DatabaseConfig(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}
}

func run() {
	// Gin
	ginServer := infra.NewServer(
		viper.GetInt("server.port"),
		infra.DebugMode,
	)
	// DB
	conn := "host=localhost port=5432 user=masahirookubo dbname=golang_api sslmode=disable" // local向け
	//conn := "postgres://postgresql:postgresql@db:5432/golang_api?sslmode=disable" // docker向け

	//conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	viper.GetString("db.host"),
	//	viper.GetInt("db.port"),
	//	viper.GetString("db.user"),
	//	viper.GetString("db.password"),
	//	viper.GetString("db.name"))
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	// Loggar
	routerLogger := logger.NewLogger("TEST",
		viper.GetString("log.level"),
		viper.GetString("log.format"),
	)

	server.NewRouterWithLogger(
		uc.HandlerConstructor{
			Logger:        routerLogger,
			BlogRW:        blogRW.New(db),
			Validator:     validator.New(),
			DBTransaction: dbTransaction.New(db),
		}.New(),
		routerLogger,
	).SetRoutes(ginServer.Router)

	ginServer.Start()
}
