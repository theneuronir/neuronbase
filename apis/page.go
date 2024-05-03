package apis

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/search"
)

// bindPageApi registers the page api endpoints and the corresponding handlers.
func bindPageApi(app core.App, rg *echo.Group) {
	api := pageApi{app: app}

	subGroup := rg.Group("/pages", ActivityLogger(app))
	subGroup.GET("", api.list, RequireAdminAuth())
	// subGroup.POST("", api.create, RequireAdminAuthOnlyIfAny(app))
	subGroup.GET("/:id", api.view, RequireAdminAuth())
	// subGroup.PATCH("/:id", api.update, RequireAdminAuth())
	// subGroup.DELETE("/:id", api.delete, RequireAdminAuth())
}

type pageApi struct {
	app core.App
}

func (api *pageApi) list(c echo.Context) error {
	fieldResolver := search.NewSimpleFieldResolver(
		"id", "created", "updated", "title",
	)

	pages := []*models.Page{}

	result, err := search.NewProvider(fieldResolver).
		Query(api.app.Dao().PageQuery()).
		ParseAndExec(c.QueryParams().Encode(), &pages)

	if err != nil {
		return NewBadRequestError("", err)
	}

	event := new(core.PagesListEvent)
	event.HttpContext = c
	event.Pages = pages
	event.Result = result

	return api.app.OnPagesListRequest().Trigger(event, func(e *core.PagesListEvent) error {
		if e.HttpContext.Response().Committed {
			return nil
		}

		return e.HttpContext.JSON(http.StatusOK, e.Result)
	})
}

func (api *pageApi) view(c echo.Context) error {
	id := c.PathParam("id")
	if id == "" {
		return NewNotFoundError("", nil)
	}

	page, err := api.app.Dao().FindPageById(id)
	if err != nil || page == nil {
		return NewNotFoundError("", err)
	}

	event := new(core.PageViewEvent)
	event.HttpContext = c
	event.Page = page

	return api.app.OnPageViewRequest().Trigger(event, func(e *core.PageViewEvent) error {
		if e.HttpContext.Response().Committed {
			return nil
		}

		return e.HttpContext.JSON(http.StatusOK, e.Page)
	})
}

func (api *pageApi) create(c echo.Context) error {
	page := &models.Page{}

	form := forms.NewPageUpsert(api.app, page)

	// load request
	if err := c.Bind(form); err != nil {
		return NewBadRequestError("Failed to load the submitted data due to invalid formatting.", err)
	}

	event := new(core.PageCreateEvent)
	event.HttpContext = c
	event.Page = page

	// create the page
	submitErr := form.Submit(func(next forms.InterceptorNextFunc[*models.Page]) forms.InterceptorNextFunc[*models.Page] {
		return func(p *models.Page) error {
			event.Page = p

			return api.app.OnPageBeforeCreateRequest().Trigger(event, func(e *core.PageCreateEvent) error {
				if err := next(e.Page); err != nil {
					return NewBadRequestError("Failed to create admin.", err)
				}

				return api.app.OnPageAfterCreateRequest().Trigger(event, func(e *core.PageCreateEvent) error {
					if e.HttpContext.Response().Committed {
						return nil
					}

					return e.HttpContext.JSON(http.StatusOK, e.Page)
				})
			})
		}
	})

	return submitErr
}

// func (api *pageApi) update(c echo.Context) error {
// 	id := c.PathParam("id")
// 	if id == "" {
// 		return NewNotFoundError("", nil)
// 	}

// 	admin, err := api.app.Dao().FindAdminById(id)
// 	if err != nil || admin == nil {
// 		return NewNotFoundError("", err)
// 	}

// 	form := forms.NewAdminUpsert(api.app, admin)

// 	// load request
// 	if err := c.Bind(form); err != nil {
// 		return NewBadRequestError("Failed to load the submitted data due to invalid formatting.", err)
// 	}

// 	event := new(core.AdminUpdateEvent)
// 	event.HttpContext = c
// 	event.Admin = admin

// 	// update the admin
// 	submitErr := form.Submit(func(next forms.InterceptorNextFunc[*models.Admin]) forms.InterceptorNextFunc[*models.Admin] {
// 		return func(m *models.Admin) error {
// 			event.Admin = m

// 			return api.app.OnAdminBeforeUpdateRequest().Trigger(event, func(e *core.AdminUpdateEvent) error {
// 				if err := next(e.Admin); err != nil {
// 					return NewBadRequestError("Failed to update admin.", err)
// 				}

// 				return api.app.OnAdminAfterUpdateRequest().Trigger(event, func(e *core.AdminUpdateEvent) error {
// 					if e.HttpContext.Response().Committed {
// 						return nil
// 					}

// 					return e.HttpContext.JSON(http.StatusOK, e.Admin)
// 				})
// 			})
// 		}
// 	})

// 	return submitErr
// }

// func (api *pageApi) delete(c echo.Context) error {
// 	id := c.PathParam("id")
// 	if id == "" {
// 		return NewNotFoundError("", nil)
// 	}

// 	admin, err := api.app.Dao().FindAdminById(id)
// 	if err != nil || admin == nil {
// 		return NewNotFoundError("", err)
// 	}

// 	event := new(core.AdminDeleteEvent)
// 	event.HttpContext = c
// 	event.Admin = admin

// 	return api.app.OnAdminBeforeDeleteRequest().Trigger(event, func(e *core.AdminDeleteEvent) error {
// 		if err := api.app.Dao().DeleteAdmin(e.Admin); err != nil {
// 			return NewBadRequestError("Failed to delete admin.", err)
// 		}

// 		return api.app.OnAdminAfterDeleteRequest().Trigger(event, func(e *core.AdminDeleteEvent) error {
// 			if e.HttpContext.Response().Committed {
// 				return nil
// 			}

// 			return e.HttpContext.NoContent(http.StatusNoContent)
// 		})
// 	})
// }
