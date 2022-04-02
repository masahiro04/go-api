package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-api/adapters"
	"go-api/adapters/controllers"
	"go-api/adapters/dao/blogDao"
	"go-api/adapters/dao/tx"
	"go-api/adapters/dao/userDao"
	"go-api/adapters/firebase"
	"go-api/adapters/loggers"

	_ "github.com/lib/pq"

	"fmt"

	infra "go-api/infrastructure"

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
	if err != nil {
		logrus.Fatal(err)
	}

	// TODO(okubo): ここでエラー出てるので、修正する
	appliedCount, err := migrate.Exec(pg, "postgres", migrations, migrate.Up)
	if err != nil {
		logrus.Fatal(err)
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

	sslmode := "require"
	if os.Getenv("ENV") == "development" {
		sslmode = "disable"
	}

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB"),
		dbPort,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		sslmode,
	)

	// migrate
	err := ExecMigrations(conn)
	if err != nil {
		log.Printf(conn)
		log.Println(err)
		log.Printf("sentinel5")
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Firebase
	client := infra.NewFirebaseAuthClient()
	firebaseHandler := firebase.New(client)

	// Loggar
	routerLogger := loggers.NewLogger(
		os.Getenv("ENV"),
		viper.GetString("log.level"),
		viper.GetString("log.format"),
	)

	drivers := adapters.NewDriver(
		adapters.Driver{
			Logger:          routerLogger,
			BlogDao:         blogDao.New(db),
			UserDao:         userDao.New(db),
			FirebaseHandler: firebaseHandler,
			DBTransaction:   tx.New(db),
			// Validator: validator.New(),
		},
	)

	controllers.NewRouter(drivers).SetRoutes(ginServer.Router)

	ginServer.Start()
}
