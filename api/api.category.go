package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

type (
	CategoryModule struct {
		db   *sql.DB
		Name string
	}

	AddCategoryParam struct {
		Name     string `json:"name"`
		ImageURL string `json:"image_url"`
	}

	OneCategoryParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllCategoryParam struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func NewCategoryModule(db *sql.DB) *CategoryModule {
	return &CategoryModule{
		db:   db,
		Name: "module/category",
	}
}

func (m CategoryModule) All(ctx context.Context, param AllCategoryParam) ([]model.CategoryModelResponse, *Error) {
	var allResp []model.CategoryModelResponse

	data, err := (&model.CategoryModel{}).All(ctx, m.db, model.AllCategory{
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all category"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no category found"
		}

		return []model.CategoryModelResponse{}, NewErrorWrap(err, m.Name, "all/category",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}
func (m CategoryModule) Add(ctx context.Context, param AddCategoryParam) (model.CategoryModelResponse, *Error) {
	category := &model.CategoryModel{
		Name:     param.Name,
		ImageURL: param.ImageURL,
	}

	data, err := category.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add category"

		return model.CategoryModelResponse{}, NewErrorWrap(err, m.Name, "add/category",
			message, status)
	}

	return data.Response(), nil
}

func (m CategoryModule) One(ctx context.Context, param OneCategoryParam) (model.CategoryModelResponse, *Error) {
	category := &model.CategoryModel{
		ID: param.ID,
	}

	data, err := category.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one category"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no category found"
		}

		return model.CategoryModelResponse{}, NewErrorWrap(err, m.Name, "one/category",
			message, status)
	}

	return data.Response(), nil
}
