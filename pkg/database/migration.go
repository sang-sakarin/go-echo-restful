package database

import "go-echo-restful/internal/news"

func UseMigration() {
	DBConn.Migrator().AutoMigrate(&news.News{})
}
