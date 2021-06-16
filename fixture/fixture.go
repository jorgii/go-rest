package fixture

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"gorm.io/gorm"
)

func Load(db *gorm.DB, dir string) error {
	base, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Printf("Fixtures dir: %s.\n", path.Join(base, dir))
	return filepath.WalkDir(path.Join(base, dir), func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".sql" {
			log.Printf("Loading data from %s.\n", path)
			sql, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("Failed reading data from %s. Skipping. %s\n", path, err)
				return nil
			}
			if err := db.Exec(string(sql)).Error; err != nil {
				fmt.Printf("Error executing SQL from %s. %s\n", path, err)
				return err
			}
		}
		return nil
	})
}
