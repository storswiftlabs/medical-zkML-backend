package article

import (
	"medical-zkml-backend/internal/db"
	"medical-zkml-backend/internal/module"
)

func SaveArticle(article *module.Article) error {
	if err := db.SaveArticle(article); err != nil {
		return err
	}
	return nil
}
