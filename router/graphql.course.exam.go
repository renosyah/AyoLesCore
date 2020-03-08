package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	courseExamType = graphql.NewObject(graphql.ObjectConfig{
		Name: "course_exam",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"course_id": &graphql.Field{
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
		course_exam_list(
			course_id : "2e847a03-5209-4d2b-9e37-b88e461e9c41",
			search_by:"text",
			search_value:"",
			order_by:"exam_index",
			order_dir:"asc",
			offset:0,
			limit:10
			limit_answer:4
		)
		{
			id,
			course_id,
			type_exam,
			exam_index,
			text,
			image_url,
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

	courseExamListField = &graphql.Field{
		Type: graphql.NewList(courseExamType),
		Args: graphql.FieldConfigArgument{
			"course_id": &graphql.ArgumentConfig{
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

			courseID, _ := uuid.FromString(p.Args["course_id"].(string))

			param := api.AllCourseExamParam{
				CourseID:    courseID,
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
				LimitAnswer: p.Args["limit_answer"].(int),
			}

			all, err := courseExamModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		course_exam_detail(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			course_id,
			type_exam,
			exam_index,
			text,
			image_url,
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

	courseExamDetailField = &graphql.Field{
		Type: courseExamType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseExamResponse{}, errUUID
			}

			data, err := courseExamModule.One(ctx, api.OneCourseExamParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_exam_register(
			course_id : "",
			type_exam : "",
			exam_index : "",
			text : "",
			image_url : ""
		)
		{
			id,
			course_id,
			type_exam,
			exam_index,
			text,
			image_url,
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

	courseExamCreateField = &graphql.Field{
		Type: courseType,
		Args: graphql.FieldConfigArgument{
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"type_exam": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"exam_index": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"image_url": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.CourseExamResponse{}, errUUID
			}

			courseExam := api.AddCourseExamParam{
				CourseID:  courseID,
				TypeExam:  int32(p.Args["type_exam"].(int)),
				ExamIndex: int32(p.Args["exam_index"].(int)),
				Text:      p.Args["text"].(string),
				ImageURL:  p.Args["image_url"].(string),
			}

			data, err := courseExamModule.Add(ctx, courseExam)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_exam_update(
			id:"",
			course_id : "",
			type_exam : "",
			exam_index : "",
			text : "",
			image_url : ""
		)
		{
			id
		}
	} */

	courseExamUpdateField = &graphql.Field{
		Type: courseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"type_exam": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"exam_index": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"image_url": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseExamResponse{}, errUUID
			}

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.CourseExamResponse{}, errUUID
			}

			courseExam := api.AddCourseExamParam{
				CourseID:  courseID,
				TypeExam:  int32(p.Args["type_exam"].(int)),
				ExamIndex: int32(p.Args["exam_index"].(int)),
				Text:      p.Args["text"].(string),
				ImageURL:  p.Args["image_url"].(string),
			}

			data, err := courseExamModule.Update(ctx, courseExam, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* {
		course_exam_delete(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id
		}
	} */

	courseExamDeleteField = &graphql.Field{
		Type: courseExamType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseExamResponse{}, errUUID
			}

			data, err := courseExamModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
