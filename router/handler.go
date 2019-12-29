package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
)

type (
	HandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, *api.Error)
)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errs []string

	// Ignore error from form parsing as it's insignificant.
	r.ParseForm()

	data, err := fn(w, r)
	if err != nil {
		log.Println(err)
		errs = append(errs, err.Error())
		w.WriteHeader(err.StatusCode)
		resp := api.Response{
			Status: http.StatusText(err.StatusCode),
			Data:   data,
			BaseResponse: api.BaseResponse{
				Errors: errs,
			},
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Println(err)
			return
		}
	} else {
		resp := api.Response{
			Status: http.StatusText(200),
			Data:   data,
			BaseResponse: api.BaseResponse{
				Errors: errs,
			},
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Println(err)
			return
		}
	}
}

func InitSchema() (graphql.Schema, error) {
	queryFields := graphql.Fields{
		"student_detail":   studentDetailField,
		"student_login":    studentLoginField,
		"student_list":     studentListField,
		"teacher_detail":   teacherDetailField,
		"teacher_login":    teacherLoginField,
		"teacher_list":     teacherListField,
		"category_detail":  categoryDetailField,
		"category_list":    categoryListField,
		"banner_detail":    bannerDetailField,
		"banner_list":      bannerListField,
		"course_detail":    courseDetailField,
		"course_list":      courseListField,
		"classroom_list":   classRoomListField,
		"classroom_detail": classRoomDetailField,
	}

	mutationFields := graphql.Fields{
		"student_register":       studentCreateField,
		"teacher_register":       teacherCreateField,
		"category_register":      categoryCreateField,
		"banner_register":        bannerCreateField,
		"course_register":        courseCreateField,
		"course_detail_register": courseDetailCreateField,
		"classroom_register":     classRoomCreateField,
	}

	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: queryFields,
		},
	)

	mutationType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: mutationFields,
		},
	)

	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
}
