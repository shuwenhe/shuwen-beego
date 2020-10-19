package dao

import (
	"testing"

	"github.com/shuwenhe/shuwen-beego/models"
)

func TestAddClass(t *testing.T) {
	cls := &models.Class{
		Name: "shuwen",
		Des:  "mathematics",
	}

	cls2 := &models.Class{
		Name: "shuwen",
		Des:  "algorithm",
	}
	AddClass(cls)
	AddClass(cls2)
}

func TestUpdateClass(t *testing.T) {
	cls := &models.Class{
		ID:   5,
		Name: "Richard",
		Des:  "son",
	}
	UpdateClass(cls)
}
