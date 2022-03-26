package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CobraInitialization() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Println(err)
	}

	viper.AutomaticEnv()

	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No configuration file found")
	}
}

func LoggerConfig(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().String("log.level", "info", "one of debug, info, warn, error or fatal")
	rootCmd.PersistentFlags().String("log.format", "text", "one of text or json")
	rootCmd.PersistentFlags().Bool("log.line", false, "enable filename and line in logs")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Println("No configuration file found")
	}
}

func ServerConfig(cmd *cobra.Command) {
	cmd.Flags().String("server.host", "127.0.0.1", "host on which the server should listen")
	cmd.Flags().Int("server.port", 8080, "port on which the server should listen")
	cmd.Flags().Bool("server.debug", false, "debug mode for the server")
	cmd.Flags().String("server.fronthost", os.Getenv("FRONT_HOST"), "allowed origins for the server")

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		log.Println("No configuration file found")
	}
}

func DatabaseConfig(cmd *cobra.Command) {
	cmd.Flags().String("db.host", "db", "host on which the db should listen")
	cmd.Flags().Int("db.port", 5432, "port on which the db should listen")
	cmd.Flags().String("db.user", "postgresql", "user on which the db should listen")
	cmd.Flags().String("db.password", "postgresql", "password on which the db should listen")
	cmd.Flags().String("db.name", "api", "name on which the db should listen")

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		log.Println("No configuration file found")
	}
}
