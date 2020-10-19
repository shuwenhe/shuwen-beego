package dao

import (
	"github.com/shuwenhe/shuwen-beego/models"
	"github.com/shuwenhe/shuwen-beego/utils"
)

func ArticlePage(pi int, ps int) ([]*models.Article, error) {
	articles := make([]*models.Article, 0, ps)
	sql := "select * from article limit ?,?"
	err := utils.Db.Select(&articles, sql, (pi-1)*ps, ps)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func ArticleCount() (int, error) {
	var count int
	sql := "SELECT COUNT(id) as count FROM article"
	err := utils.Db.Get(&count, sql)
	if err != nil {
		utils.NewResult(utils.Fail, err.Error())
		return 0, err
	}
	return count, nil
}
