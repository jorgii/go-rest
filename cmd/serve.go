package cmd

import (
	"gorest/database"
	"gorest/router"
	"log"

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
		db, err := database.ConnectDB()
		if err != nil {
			panic("Failed to connect to the database.")
		}
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		app := router.NewApp(db)

		log.Fatal(app.Listen("127.0.0.1:8080"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
