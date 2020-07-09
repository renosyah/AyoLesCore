package router

import (
	"github.com/graphql-go/graphql"
)

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
		"course_exam_solution_detail":    courseExamSolutionDetailField,
		"course_exam_solution_list":      courseExamSolutionListField,
	}

	mutationFields := graphql.Fields{
		"student_register":                 studentCreateField,
		"student_update":                   studentUpdateField,
		"student_delete":                   studentDeleteField,
		"teacher_register":                 teacherCreateField,
		"teacher_update":                   teacherUpdateField,
		"teacher_delete":                   teacherDeleteField,
		"category_register":                categoryCreateField,
		"category_update":                  categoryUpdateField,
		"category_delete":                  categoryDeleteField,
		"banner_register":                  bannerCreateField,
		"banner_update":                    bannerUpdateField,
		"banner_delete":                    bannerDeleteField,
		"course_register":                  courseCreateField,
		"course_update":                    courseUpdateField,
		"course_delete":                    courseDeleteField,
		"course_detail_register":           courseDetailCreateField,
		"course_detail_update":             courseDetailUpdateField,
		"course_detail_delete":             courseDetailDeleteField,
		"classroom_register":               classRoomCreateField,
		"classroom_update":                 classRoomUpdateField,
		"classroom_delete":                 classRoomDeleteField,
		"course_material_register":         courseMaterialCreateField,
		"course_material_update":           courseMaterialUpdateField,
		"course_material_delete":           courseMaterialDeleteField,
		"course_material_detail_register":  courseMaterialDetailCreateField,
		"course_material_detail_update":    courseMaterialDetailUpdateField,
		"course_material_detail_delete":    courseMaterialDetailDeleteField,
		"classroom_progress_register":      classroomProgressRegisterField,
		"classroom_progress_update":        classroomProgressUpdateField,
		"classroom_progress_delete":        classroomProgressDeleteField,
		"course_exam_register":             courseExamCreateField,
		"course_exam_update":               courseExamUpdateField,
		"course_exam_delete":               courseExamDeleteField,
		"course_exam_answer_register":      courseExamAnswerCreateField,
		"course_exam_answer_update":        courseExamAnswerUpdateField,
		"course_exam_answer_delete":        courseExamAnswerDeletelField,
		"classroom_exam_progress_register": classroomExamProgressRegisterField,
		"classroom_exam_progress_update":   classroomExamProgressUpdateField,
		"classroom_exam_progress_delete":   classroomExamProgressDeleteField,
		"classroom_certificate_register":   classroomCertificateRegisterField,
		"course_qualification_register":    courseQualificationCreateField,
		"course_qualification_update":      courseQualificationUpdateField,
		"course_qualification_delete":      courseQualificationDeleteField,
		"course_exam_solution_register":    courseExamSolutionCreateField,
		"course_exam_solution_update":      courseExamSolutionUpdateField,
		"course_exam_solution_delete":      courseExamSolutionDeleteField,
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
