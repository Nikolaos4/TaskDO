package db
import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Login string `gorm:"login"`
	//Name string `gorm:"name"`
	//LastName string `gorm:"lastname"`
	Password string `gorm:"password"`
	Tasks []Task `gorm:"foreignKey:UserID"`
}

type Task struct {
	gorm.Model
	Title string 
	Data time.Time 
	UserID uint
}



