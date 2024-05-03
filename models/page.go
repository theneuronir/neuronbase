package models

var (
	_ Model = (*Page)(nil)
)

type Page struct {
	BaseModel

	Title string `db:"title" json:"title"`
}

// TableName returns the Page model SQL table name.
func (m *Page) TableName() string {
	return "_pages"
}
