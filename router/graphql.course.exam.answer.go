package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	courseExamAnswerType = graphql.NewObject(graphql.ObjectConfig{
		Name: "course_exam_answer",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"course_exam_id": &graphql.Field{
				Type: graphql.String,
			},
			"type_answer": &graphql.Field{
				Type: graphql.Int,
			},
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		course_exam_answer_list(
			course_exam_id : "",
			search_by:"title",
			search_value:"",
			order_by:"material_index",
			order_dir:"asc",
			offset:0,
			limit:10
		)
			id,
			course_exam_id,
			type_answer,
			label,
			text,
			image_url
		}
	} */

	courseExamAnswerListField = &graphql.Field{
		Type: graphql.NewList(courseExamAnswerType),
		Args: graphql.FieldConfigArgument{
			"course_exam_id": &graphql.ArgumentConfig{
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
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseExamID, _ := uuid.FromString(p.Args["course_exam_id"].(string))

			param := api.AllCourseExamAnswerParam{
				CourseExamID: courseExamID,
				SearchBy:     p.Args["search_by"].(string),
				SearchValue:  p.Args["search_value"].(string),
				OrderBy:      p.Args["order_by"].(string),
				OrderDir:     p.Args["order_dir"].(string),
				Offset:       p.Args["offset"].(int),
				Limit:        p.Args["limit"].(int),
			}

			all, err := courseExamAnswerModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		course_exam_answer_detail(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			course_exam_id,
			type_answer,
			label,
			text,
			image_url
		}
	} */

	courseExamAnswerDetailField = &graphql.Field{
		Type: courseExamAnswerType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseExamAnswerResponse{}, errUUID
			}

			data, err := courseExamAnswerModule.One(ctx, api.OneCourseExamAnswerParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_exam_answer_register(
			course_exam_id : "",
			type_answer : 0,
			label : "",
			text : "",
			image_url : ""
		)
		{
			id,
			course_exam_id,
			type_answer,
			label,
			text,
			image_url
		}
	} */

	courseExamAnswerCreateField = &graphql.Field{
		Type: courseExamAnswerType,
		Args: graphql.FieldConfigArgument{
			"course_exam_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"type_answer": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"label": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"image_url": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseExamID, errUUID := uuid.FromString(p.Args["course_exam_id"].(string))
			if errUUID != nil {
				return model.CourseExamResponse{}, errUUID
			}

			courseExamAnswer := api.AddCourseExamAnswerParam{
				CourseExamID: courseExamID,
				TypeAnswer:   int32(p.Args["type_answer"].(int)),
				Label:        p.Args["label"].(string),
				Text:         p.Args["text"].(string),
				ImageURL:     p.Args["image_url"].(string),
			}

			data, err := courseExamAnswerModule.Add(ctx, courseExamAnswer)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
