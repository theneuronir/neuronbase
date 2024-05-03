package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/forms/validators"
	"github.com/pocketbase/pocketbase/models"
)

// PageUpsert is a [models.Page] upsert (create/update) form.
type PageUpsert struct {
	app  core.App
	dao  *daos.Dao
	page *models.Page

	Id    string `form:"id" json:"id"`
	Title string `form:"title" json:"title"`
}

// NewPageUpsert creates a new [PageUpsert] form with initializer
// config created from the provided [core.App] and [models.Page] instances
// (for create you could pass a pointer to an empty Page - `&models.Page{}`).
//
// If you want to submit the form as part of a transaction,
// you can change the default Dao via [SetDao()].
func NewPageUpsert(app core.App, page *models.Page) *PageUpsert {
	form := &PageUpsert{
		app:  app,
		dao:  app.Dao(),
		page: page,
	}

	// load defaults
	form.Id = page.Id
	form.Title = page.Title

	return form
}

// SetDao replaces the default form Dao instance with the provided one.
func (form *PageUpsert) SetDao(dao *daos.Dao) {
	form.dao = dao
}

// Validate makes the form validatable by implementing [validation.Validatable] interface.
func (form *PageUpsert) Validate() error {
	return validation.ValidateStruct(form,
		validation.Field(
			&form.Id,
			validation.When(
				form.page.IsNew(),
				validation.Length(models.DefaultIdLength, models.DefaultIdLength),
				validation.Match(idRegex),
				validation.By(validators.UniqueId(form.dao, form.page.TableName())),
			).Else(validation.In(form.page.Id)),
		),
		validation.Field(
			&form.Title,
			validation.Required,
			validation.Length(1, 255),
		),
	)
}

// Submit validates the form and upserts the form page model.
//
// You can optionally provide a list of InterceptorFunc to further
// modify the form behavior before persisting it.
func (form *PageUpsert) Submit(interceptors ...InterceptorFunc[*models.Page]) error {
	if err := form.Validate(); err != nil {
		return err
	}

	// custom insertion id can be set only on create
	if form.page.IsNew() && form.Id != "" {
		form.page.MarkAsNew()
		form.page.SetId(form.Id)
	}

	form.page.Title = form.Title

	return runInterceptors(form.page, func(page *models.Page) error {
		return form.dao.SavePage(page)
	}, interceptors...)
}
