package service

import (
	"fmt"
	"goProject/pojo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:zqayy20011013@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
)

func insert(user pojo.User) error {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
func show(id int) (pojo.User, error) {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var user pojo.User
	if result := db.First(&user, id); result.Error != nil {
		return user, result.Error
	}
	fmt.Print("show method")
	fmt.Print(user.Name)
	return user, nil
}
func search(name string) (pojo.User, error) {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var users pojo.User
	if result := db.Where("name=?", name).First(&users); result.Error != nil {
		return users, result.Error
	}
	fmt.Print(users)
	return users, nil
}
func translationsearch(name string) (pojo.User, error) {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	tx := db.Begin()
	var users pojo.User
	if result := tx.Where("name=?", name).First(&users); result.Error != nil {
		tx.Rollback()
	}
	tx.Commit()
	return users, nil
}
func translationsearchauto(name string) (pojo.User, error) {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var users pojo.User
	db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("name=?", name).Find(&users); result.Error != nil {
			return result.Error
		}
		return nil
	})
	return users, nil
}
func delete(id int) error {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("id=?", id).Delete(&pojo.User{}); result.Error != nil {
			return result.Error
		}
		return nil
	})
	return nil
}
