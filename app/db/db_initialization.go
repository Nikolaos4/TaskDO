package db
import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Faild to connect to database: ", err)
	}

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Faild to connect to database: ", err)
	}

	if err := DB.AutoMigrate(&User{}, &Task{}); err != nil {
		log.Fatal("Faild to migrate model: ", err)
	}
}