package models

import (
	"time"

	"github.com/shuwenhe/shuwen-beego/utils"
)

type Article struct {
	ID      int       `json:"id,omitempty"`
	Cid     int       `json:"cid,omitempty"`
	Title   string    `json:"title,omitempty"`
	Author  string    `json:"author,omitempty"`
	Content string    `json:"content,omitempty"`
	Hits    string    `json:"hits,omitempty"`
	Ctime   time.Time `json:"ctime,omitempty"`
}

func ArticlePage(pi int, ps int) ([]Article, error) {
	mods := make([]Article, 0, ps)
	sql := ""
	utils.Db.Select(&mods, sql)

}
