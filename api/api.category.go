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

func (m CategoryModule) All(ctx context.Context, param AllCategoryParam) ([]model.CategoryResponse, *Error) {
	var allResp []model.CategoryResponse

	data, err := (&model.Category{}).All(ctx, m.db, model.AllCategory{
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

		return []model.CategoryResponse{}, NewErrorWrap(err, m.Name, "all/category",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}
func (m CategoryModule) Add(ctx context.Context, param AddCategoryParam) (model.CategoryResponse, *Error) {
	category := &model.Category{
		Name:     param.Name,
		ImageURL: param.ImageURL,
	}

	id, err := category.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add category"

		return model.CategoryResponse{}, NewErrorWrap(err, m.Name, "add/category",
			message, status)
	}

	category.ID = id

	return category.Response(), nil
}

func (m CategoryModule) One(ctx context.Context, param OneCategoryParam) (model.CategoryResponse, *Error) {
	category := &model.Category{
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

		return model.CategoryResponse{}, NewErrorWrap(err, m.Name, "one/category",
			message, status)
	}

	return data.Response(), nil
}

func (m CategoryModule) Update(ctx context.Context, param AddCategoryParam, id uuid.UUID) (model.CategoryResponse, *Error) {
	var emptyUUID uuid.UUID

	category := &model.Category{
		ID:       id,
		Name:     param.Name,
		ImageURL: param.ImageURL,
	}

	i, err := category.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update category"

		return model.CategoryResponse{}, NewErrorWrap(err, m.Name, "update/category",
			message, status)
	}

	return category.Response(), nil
}

func (m CategoryModule) Delete(ctx context.Context, id uuid.UUID) (model.CategoryResponse, *Error) {
	var emptyUUID uuid.UUID

	category := &model.Category{
		ID: id,
	}

	i, err := category.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete category"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no category found"
		}

		return model.CategoryResponse{}, NewErrorWrap(err, m.Name, "delete/category",
			message, status)
	}

	return category.Response(), nil
}
