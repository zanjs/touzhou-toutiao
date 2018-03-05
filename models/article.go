package models

import (
	"time"

	"anla.io/taizhou-ir/db"
	"github.com/houndgo/suuid"
	"github.com/theplant/batchputs"
)

type (
	// ArticleType is
	articleType struct {
		New    int
		Bubble int
	}
)

// ArticleType 发帖类型
var ArticleType = articleType{
	Bubble: 1, // 冒泡， 用户发帖
	New:    2, // 专家发帖
}

type (

	// Article is
	Article struct {
		IDModel
		TimeAllModel
		User          User          `gorm:"Table:user;foreignkey:UserID;AssociationForeignKey:id" json:"user,omitempty"`
		Pics          []*ArticlePic `json:"pics,omitempty"`
		UserID        string        `json:"user_id"`
		Title         string        `json:"title" gorm:"type:varchar(100)"`
		Content       string        `json:"content" gorm:"type:text"`
		ViewCount     int           `json:"view_count"`
		CommentCount  uint          `json:"comment_count"`
		Categories    []Category    `gorm:"many2many:article_category;ForeignKey:ID;AssociationForeignKey:ID" json:"categories"`
		LastUserID    uint          `json:"last_user_id"` //最后一个回复话题的人
		LastUser      User          `json:"last_user"`
		LastCommentAt *time.Time    `json:"last_comment_at"`
		Disabled      bool          `json:"disabled" gorm:"default:'0'"`
		Comments      []*Comment    `json:"comments" gorm:"-"`
	}
)

//BeforeSave is
// func (a *Article) BeforeSave(scope *gm.Scope) (err error) {
// 	a.ID = suuid.New().String()
// 	return err
// }

// Create is
func (a Article) Create(m *Article) error {
	var err error
	m.ID = suuid.New().String()
	rows := [][]interface{}{}

	// pics := &m.Pics
	createTime := time.Now()
	for i := 0; i < len(m.Pics); i++ {
		m.Pics[i].ArticleID = m.ID
		uid := suuid.New().String()
		m.Pics[i].ID = uid
		rows = append(rows, []interface{}{
			m.ID,
			createTime,
			m.Pics[i].Src,
			uid,
		})
	}

	columns := []string{"article_id", "created_at", "src", "id"}
	dialect := "mysql"

	err = batchputs.Put(gorm.MysqlConn().DB(), dialect, "article_pics", "article_id", columns, rows)
	if err != nil {
		panic(err)
	}

	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err
}

// GetAll is find
func (a Article) GetAll(page *PageModel) ([]Article, error) {
	var (
		data []Article
		err  error
	)

	if page.Num < 1 {
		page.Num = 1
	}

	pageSize := 2
	offset := (page.Num - 1) * pageSize

	tx := gorm.MysqlConn().Begin()

	if err = tx.Preload("User").Preload("Pics").Find(&data).Count(&page.Count).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	if err = tx.Offset(offset).Limit(pageSize).Preload("User").Preload("Pics").Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	// tx.Model(&data).Related(&pics)
	// tx.Model(&data).Related(&pics, "Pics")
	// tx.Model(&data).Association("Pics")

	// fmt.Println(pics)

	tx.Commit()

	return data, err
}
