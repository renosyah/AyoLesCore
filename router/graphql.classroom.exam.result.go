package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	classroomExamResultType = graphql.NewObject(graphql.ObjectConfig{
		Name: "classroom_exam_result",
		Fields: graphql.Fields{
			"course_exam_id": &graphql.Field{
				Type: graphql.String,
			},
			"course_id": &graphql.Field{
				Type: graphql.String,
			},
			"classroom_id": &graphql.Field{
				Type: graphql.String,
			},
			"student_answer_id": &graphql.Field{
				Type: graphql.String,
			},
			"valid_answer_id": &graphql.Field{
				Type: graphql.String,
			},
			"type_exam": &graphql.Field{
				Type: graphql.Int,
			},
			"exam_index": &graphql.Field{
				Type: graphql.Int,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"answers": &graphql.Field{
				Type: graphql.NewList(courseExamAnswerType),
			},
		},
	})

	/* {
		classroom_exam_result_list(
			classroom_id : "2e847a03-5209-4d2b-9e37-b88e461e9c41",
			search_by:"course_exam.text",
			search_value:"",
			order_by:"classroom_exam_progress.create_at",
			order_dir:"asc",
			offset:0,
			limit:10,
			limit_answer:4
		)
		{
			course_exam_id,
			course_id,
			classroom_id,
			student_answer_id,
			right_answer_id,
			type_exam,
			exam_index,
			text,
			image_url
			answers {
				id,
				course_exam_id,
				type_answer,
				label,
				text,
				image_url
			}
		}
	} */

	classroomExamResultListField = &graphql.Field{
		Type: graphql.NewList(classroomExamResultType),
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"search_by": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"search_value": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"order_by": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"order_dir": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"limit_answer": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			classRoomID, _ := uuid.FromString(p.Args["classroom_id"].(string))

			param := api.AllClassRoomExamResultParam{
				ClassRoomID: classRoomID,
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
				LimitAnswer: p.Args["limit_answer"].(int),
			}

			all, err := classRoomExamResultModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		classroom_exam_result_detail(
			course_exam_id: "4252869c-ddd2-466f-8528-e1fe8aff4135",
			course_id: "4252869c-ddd2-466f-8528-e1fe8aff4135",
			limit_answer: 4
		)
		{
			course_exam_id,
			course_id,
			classroom_id,
			student_answer_id,
			right_answer_id,
			type_exam,
			exam_index,
			text,
			image_url
			answers {
				id,
				course_exam_id,
				type_answer,
				label,
				text,
				image_url
			}
		}
	} */

	classroomExamResultDetailField = &graphql.Field{
		Type: classroomExamResultType,
		Args: graphql.FieldConfigArgument{
			"course_exam_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"limit_answer": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseExamID, errUUID := uuid.FromString(p.Args["course_exam_id"].(string))
			if errUUID != nil {
				return model.ClassRoomExamResultResponse{}, errUUID
			}

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.ClassRoomExamResultResponse{}, errUUID
			}

			data, err := classRoomExamResultModule.One(ctx, api.OneClassRoomExamResultParam{
				CourseExamID: courseExamID,
				CourseID:     courseID,
				LimitAnswer:  p.Args["limit_answer"].(int),
			})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)

// ITS DOESNOT HAVE TABLE
// THIS MODEL VALUE RESULT FROM
// QUERY JOIN
// NO UPDATE
// NO DELETE
