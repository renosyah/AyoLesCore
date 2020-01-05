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
		"student_detail":                 studentDetailField,
		"student_login":                  studentLoginField,
		"student_list":                   studentListField,
		"teacher_detail":                 teacherDetailField,
		"teacher_login":                  teacherLoginField,
		"teacher_list":                   teacherListField,
		"category_detail":                categoryDetailField,
		"category_list":                  categoryListField,
		"banner_detail":                  bannerDetailField,
		"banner_list":                    bannerListField,
		"course_detail":                  courseDetailField,
		"course_list":                    courseListField,
		"course_detail_list":             courseDetailListField,
		"classroom_list":                 classRoomListField,
		"classroom_detail":               classRoomDetailField,
		"classroom_detail_by_id":         classRoomDetailByStudentAndCourseIdField,
		"course_material_detail":         courseMaterialDetailField,
		"course_material_list":           courseMaterialListField,
		"course_material_detail_list":    courseMaterialDetailListField,
		"classroom_progress_list":        classroomProgressListField,
		"classroom_progress_detail":      classroomProgressDetailField,
		"course_exam_detail":             courseExamDetailField,
		"course_exam_list":               courseExamListField,
		"course_exam_answer_detail":      courseExamAnswerDetailField,
		"course_exam_answer_list":        courseExamAnswerListField,
		"classroom_exam_progress_detail": classroomExamProgressDetailField,
		"classroom_exam_progress_list":   classroomExamProgressListField,
		"classroom_exam_result_detail":   classroomExamResultDetailField,
		"classroom_exam_result_list":     classroomExamResultListField,
		"classroom_certificate_detail":   classroomCertificateDetailField,
		"classroom_certificate_list":     classroomCertificateListField,
		"course_qualification_detail":    courseQualificationDetailField,
		"class_qualification_detail":     classQualificationDetailField,
	}

	mutationFields := graphql.Fields{
		"student_register":                 studentCreateField,
		"student_update":                   studentUpdateField,
		"teacher_register":                 teacherCreateField,
		"category_register":                categoryCreateField,
		"banner_register":                  bannerCreateField,
		"course_register":                  courseCreateField,
		"course_detail_register":           courseDetailCreateField,
		"classroom_register":               classRoomCreateField,
		"course_material_register":         courseMaterialCreateField,
		"course_material_detail_register":  courseMaterialDetailCreateField,
		"classroom_progress_register":      classroomProgressRegisterField,
		"course_exam_register":             courseExamCreateField,
		"course_exam_answer_register":      courseExamAnswerCreateField,
		"classroom_exam_progress_register": classroomExamProgressRegisterField,
		"classroom_exam_progress_delete":   classroomExamProgressDeleteField,
		"classroom_certificate_register":   classroomCertificateRegisterField,
		"course_qualification_register":    courseQualificationCreateField,
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
