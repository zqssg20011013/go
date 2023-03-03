package pojo

import (
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	_ "goProject/config"
	_ "net/http"
	"time"
)

type User struct {
	Id       uint32 `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int32  `gorm:"size:3"`
	Birthday time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}
