package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-api/adapters/controllers"
	"go-api/adapters/dao/blogDao"
	"go-api/adapters/dao/userDao"
	"go-api/adapters/loggers"

	_ "github.com/lib/pq"

	"fmt"

	infra "go-api/infrastructure"
	uc "go-api/usecases"

	migrate "github.com/rubenv/sql-migrate"
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
	Use:   "go-api",
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

func ExecMigrations(postgresURL string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	pg, err := sql.Open("postgres", postgresURL)

	logrus.WithError(err).Fatal()

	appliedCount, err := migrate.Exec(pg, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	log.Printf("Applied %v migrations", appliedCount)
	return nil
}

func run() {
	// Gin
	ginServer := infra.NewServer(
		viper.GetInt("server.port"),
		infra.DebugMode,
	)

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	// DB
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB"),
		dbPort,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// migrate
	err := ExecMigrations(conn)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Loggar
	routerLogger := loggers.NewLogger("TEST",
		viper.GetString("log.level"),
		viper.GetString("log.format"),
	)

	controllers.NewRouterWithLogger(
		uc.HandlerConstructor{
			Logger:  routerLogger,
			BlogDao: blogDao.New(db),
			UserDao: userDao.New(db),
			// Validator: validator.New(),
			// DBTransaction: dbTransaction.New(db),
		}.New(),
		routerLogger,
	).SetRoutes(ginServer.Router)

	ginServer.Start()
}
