package cmd

import (
	"gorest/config"
	"gorest/database"
	"gorest/fixture"
	"log"

	"github.com/spf13/cobra"
)

var loadFixturesCmd = &cobra.Command{
	Use:   "loadfixtures",
	Short: "Load fixtures from a directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		db, err := database.ConnectDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
		if err != nil {
			panic("Failed to connect to the database.")
		}
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
		db = db.Begin()
		if err := fixture.Load(db, args[0]); err != nil {
			log.Println("Failed loading fixtures.")
			db.Rollback()
		}
		db.Commit()
		log.Println("Fixtures loaded successfully.")
	},
}

func init() {
	rootCmd.AddCommand(loadFixturesCmd)
}
