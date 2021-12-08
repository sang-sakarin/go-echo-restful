package pkg

import (
	"go-echo-restful/internal/news"
	"go-echo-restful/pkg/database"
)

func UseMigration() {
	database.DBConn.Migrator().AutoMigrate(&news.News{})
}
