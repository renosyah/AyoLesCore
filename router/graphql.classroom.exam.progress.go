package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	classroomExamProgressType = graphql.NewObject(graphql.ObjectConfig{
		Name: "classroom_exam_progress",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"classroom_id": &graphql.Field{
				Type: graphql.String,
			},
			"course_exam_id": &graphql.Field{
				Type: graphql.String,
			},
			"course_exam_answer_id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		classroom_exam_progress_list(
			classroom_id:"",
			order_by:"create_at",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			classroom_id,
			course_exam_id,
			course_exam_answer_id
		 }
	} */

	classroomExamProgressListField = &graphql.Field{
		Type: graphql.NewList(classroomExamProgressType),
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"order_by": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"order_dir": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			classRoomID, errUUID := uuid.FromString(p.Args["classroom_id"].(string))
			if errUUID != nil {
				return []model.ClassRoomExamProgressResponse{}, errUUID
			}

			param := api.AllClassRoomExamProgressParam{
				ClassroomID: classRoomID,
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := classRoomExamProgressModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		classroom_exam_progress_detail(
			id: ""
		)
		{
			id,
			classroom_id,
			course_exam_id,
			course_exam_answer_id
		 }
	} */

	classroomExamProgressDetailField = &graphql.Field{
		Type: graphql.NewList(classroomExamProgressType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.ClassRoomExamProgressResponse{}, errUUID
			}

			param := api.OneClassRoomExamProgressParam{
				ID: id,
			}

			one, err := classRoomExamProgressModule.One(ctx, param)
			if err != nil {
				log.Println(err)
				return one, err
			}

			return one, nil
		},
	}

	/* mutation {
		classroom_exam_progress_register(
			classroom_id : "",
			course_exam_id : "",
			course_exam_answer_id : ""
		)
		{
			id,
			classroom_id,
			course_exam_id,
			course_exam_answer_id
		 }
	} */

	classroomExamProgressRegisterField = &graphql.Field{
		Type: classroomExamProgressType,
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_exam_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_exam_answer_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			classRoomID, errUUID := uuid.FromString(p.Args["classroom_id"].(string))
			if errUUID != nil {
				return model.ClassRoomExamProgressResponse{}, errUUID
			}

			courseExamID, errUUID := uuid.FromString(p.Args["course_exam_id"].(string))
			if errUUID != nil {
				return model.ClassRoomExamProgressResponse{}, errUUID
			}

			courseCourseExamAnswerID, errUUID := uuid.FromString(p.Args["course_exam_answer_id"].(string))
			if errUUID != nil {
				return model.ClassRoomExamProgressResponse{}, errUUID
			}

			param := api.AddClassRoomExamParam{
				ClassroomID:        classRoomID,
				CourseExamID:       courseExamID,
				CourseExamAnswerID: courseCourseExamAnswerID,
			}

			data, err := classRoomExamProgressModule.Add(ctx, param)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		classroom_exam_progress_delete(
			classroom_id : "",
		)
		{
			id,
			classroom_id,
			course_exam_id,
			course_exam_answer_id
		 }
	} */

	classroomExamProgressDeleteField = &graphql.Field{
		Type: classroomExamProgressType,
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			classRoomID, errUUID := uuid.FromString(p.Args["classroom_id"].(string))
			if errUUID != nil {
				return model.ClassRoomExamProgressResponse{}, errUUID
			}

			data, err := classRoomExamProgressModule.Delete(ctx, classRoomID)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
