package dao

import (
	"fmt"
	"testing"
)

func TestArticlePage(t *testing.T) {
	articles, _ := ArticlePage(1, 1)
	for _, article := range articles {
		fmt.Println("article=", article)
	}
}