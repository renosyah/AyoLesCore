package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/renosyah/AyoLesCore/api"
)

var (
	studentModule  *api.StudentModule
	categoryModule *api.CategoryModule
	bannerModule   *api.BannerModule
	db             *sql.DB
)

func Init(d *sql.DB) {
	db = d
	studentModule = api.NewStudentModule(db)
	categoryModule = api.NewCategoryModule(db)
	bannerModule = api.NewBannerModule(db)
}

// ParseBodyData parse json-formatted request body into given struct.
func ParseBodyData(ctx context.Context, r *http.Request, data interface{}) error {
	bBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "read")
	}

	err = json.Unmarshal(bBody, data)
	if err != nil {
		return errors.Wrap(err, "json")
	}

	valid, err := govalidator.ValidateStruct(data)
	if err != nil {
		return errors.Wrap(err, "validate")
	}

	if !valid {
		return errors.New("invalid data")
	}

	return nil
}
