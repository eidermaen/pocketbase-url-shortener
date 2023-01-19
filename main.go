package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"log"
	"net/http"
	"strings"
)

var _ models.Model = (*Link)(nil)

// Struct for working with records right in the code
type Link struct {
	models.BaseModel

	Slug   string `db:"slug"`
	Url    string `db:"url"`
	Clicks int    `db:"clicks"`
}

// Name of the collection
func (l *Link) TableName() string {
	return "link"
}

func LinkQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Link{})
}

// Query for searching for records by their slug
func FindLinkBySlug(dao *daos.Dao, slug string) (*Link, error) {
	link := &Link{}

	err := LinkQuery(dao).
		AndWhere(dbx.NewExp("LOWER(slug)={:slug}", dbx.Params{
			"slug": strings.ToLower(slug),
		})).
		Limit(1).
		One(link)

	if err != nil {
		return nil, err
	}

	return link, nil
}

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Add(http.MethodGet, "/:slug", func(c echo.Context) error {
			slug := c.PathParam("slug")
			link, err := FindLinkBySlug(app.Dao(), slug)

			if err != nil {
				return apis.NewNotFoundError("Not found", err)
			}

			currentCount := link.Clicks
			link.Clicks = currentCount + 1
			if err = app.Dao().Save(link); err != nil {
				return apis.NewApiError(500, "Internal server error", nil)
			}

			return c.Redirect(http.StatusMovedPermanently, link.Url)
		},
			func(next echo.HandlerFunc) echo.HandlerFunc {
				apis.ActivityLogger(app)
				return next
			})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
