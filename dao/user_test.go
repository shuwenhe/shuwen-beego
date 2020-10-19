package dao

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	num := "1"
	user, _ := Login(num)
	fmt.Println("user=", user)
}
