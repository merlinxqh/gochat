package dao

import (
	"testing"
	"time"
)

func TestUserDb(t *testing.T) {
	user := User{
		Id:         1,
		UserName:   "admin",
		Password:   "123456",
		CreateTime: time.Now(),
	}
	user.Add()
}
