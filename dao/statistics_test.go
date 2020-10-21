package dao

import (
	"fmt"
	"testing"
)

func TestStatisticsAll(t *testing.T) {
	articles, _ := StatisticsAll()
	for _, mod := range articles {
		fmt.Println("mod = ", mod)
	}
}
