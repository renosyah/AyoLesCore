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
	BannerModule struct {
		db   *sql.DB
		Name string
	}

	AddBannerParam struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}

	OneBannerParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllBannerParam struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func NewBannerModule(db *sql.DB) *BannerModule {
	return &BannerModule{
		db:   db,
		Name: "module/banner",
	}
}

func (m BannerModule) All(ctx context.Context, param AllBannerParam) ([]model.BannerResponse, *Error) {
	var allResp []model.BannerResponse

	banner := &model.Banner{}
	data, err := banner.All(ctx, m.db, model.AllBanner{
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on all banner"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no banner found"
		}

		return []model.BannerResponse{}, NewErrorWrap(err, m.Name, "all/banner",
			message, status)
	}

	for _, banner := range data {
		allResp = append(allResp, banner.Response())
	}

	return allResp, nil
}

func (m BannerModule) Add(ctx context.Context, param AddBannerParam) (model.BannerResponse, *Error) {
	banner := &model.Banner{
		Title:    param.Title,
		Content:  param.Content,
		ImageURL: param.ImageURL,
	}

	id, err := banner.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add banner"

		return model.BannerResponse{}, NewErrorWrap(err, m.Name, "add/banner",
			message, status)
	}

	banner.ID = id

	return banner.Response(), nil
}

func (m *BannerModule) One(ctx context.Context, param OneBannerParam) (model.BannerResponse, *Error) {
	banner := &model.Banner{
		ID: param.ID,
	}

	data, err := banner.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add banner"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no banner found"
		}

		return model.BannerResponse{}, NewErrorWrap(err, m.Name, "one/banner",
			message, status)
	}

	return data.Response(), nil
}
