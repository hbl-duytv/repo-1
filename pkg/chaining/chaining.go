package chaining

import (
	"fmt"

	"github.com/hbl-duytv/repo-1/internal"
	"github.com/jinzhu/gorm"
)

func UserStatus(status int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}

func UserNameLike(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(UserStatus(1)).Where("firstname LIKE ?", fmt.Sprint("%", name, "%"))
	}
}

func UserIDs(ids []uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN (?)", ids)
	}
}

type User struct {
	ID        uint64
	Firstname string
	Lastname  string
}

func (u *User) AfterFind() {
	u.Firstname = fmt.Sprint("found* ", u.Firstname)
	return
}

func C() {
	var users []User
	internal.DB.Scopes(UserStatus(1)).Find(&users)
	fmt.Println("Active users", users)
	internal.DB.Scopes(UserNameLike("duy")).Find(&users)
	fmt.Println("Active user with name", users)
	internal.DB.Scopes(UserIDs([]uint64{})).Find(&users)
	fmt.Println("User with ids", users)
}
