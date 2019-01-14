package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// DB connect Db
var DB *gorm.DB

//Database define Database struct
type Database struct {
	Adapter string
	DNS     string
	LogMode bool
}

//Config define Config struct
type Config struct {
	Database Database
}

func loadEnv() Config {
	rootDir := os.Getenv("GOPATH") + "/src/github.com/hbl-duytv/repo-1/internal/"
	if err := godotenv.Load(rootDir + ".env"); err != nil {
		log.Fatal(err)
	}
	config := Config{
		Database{
			Adapter: os.Getenv("DB_DRIVER"),
			DNS:     os.Getenv("DB_DNS"),
			LogMode: os.Getenv("DB_LOG_MODE") == "true",
		},
	}
	return config
}

func getDB() *gorm.DB {
	config := loadEnv()
	db, err := gorm.Open(config.Database.Adapter, config.Database.DNS)
	if err != nil {
		log.Fatalln(err)
	}
	db.LogMode(config.Database.LogMode)
	return db
}

// InitDB init DB
func InitDB() {
	DB = getDB()
	fmt.Println("connect DB success!")
}
