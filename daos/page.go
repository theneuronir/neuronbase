package daos

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
)

// PageQuery returns a new Page select query.
func (dao *Dao) PageQuery() *dbx.SelectQuery {
	return dao.ModelQuery(&models.Page{})
}

// FindPageById finds the page with the provided id.
func (dao *Dao) FindPageById(id string) (*models.Page, error) {
	model := &models.Page{}

	err := dao.PageQuery().
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(model)

	if err != nil {
		return nil, err
	}

	return model, nil
}

// TotalPages returns the number of existing page records.
func (dao *Dao) TotalPages() (int, error) {
	var total int

	err := dao.PageQuery().Select("count(*)").Row(&total)

	return total, err
}

// DeletePage deletes the provided Page model.
func (dao *Dao) DeletePage(page *models.Page) error {
	return dao.Delete(page)
}

// SavePage upserts the provided Page model.
func (dao *Dao) SavePage(page *models.Page) error {
	return dao.Save(page)
}
