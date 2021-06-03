package cmd

import (
	"gorest/config"
	"gorest/database"
	"gorest/handler"
	"gorest/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		db, err := database.ConnectDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
		if err != nil {
			panic("Failed to connect to the database.")
		}
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		app := fiber.New()
		app.Use(logger.New())
		h := handler.New(db, cfg)
		router.SetupRoutes(app, h)

		log.Fatal(app.Listen(cfg.ListenAddr))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
