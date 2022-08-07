package test

import (
	"blogbackend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "root:password@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	//data := make([]*models.User, 0)
	//err = db.Find(&data).Error
	data := new(models.User)
	err = db.Where("name = ? AND password = ?", "harris", "password").First(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	//for _, v := range data {
	//	fmt.Printf("Problem ==> %v \n", v)
	//}
}
