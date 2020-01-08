package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	courseExamSolutionType = graphql.NewObject(graphql.ObjectConfig{
		Name: "course_exam_solution",
		Fields: graphql.Fields{
			"id": &graphql.Field{
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
		course_exam_solution_list(
			course_exam_id : "fb8478fa-b27f-42a9-aebf-acaa09b3f408",
			order_by:"create_at",
			order_dir:"asc",
			offset:0,
			limit:10
		)
			id,
			course_exam_id,
			course_exam_answer_id
		}
	} */

	courseExamSolutionListField = &graphql.Field{
		Type: graphql.NewList(courseExamSolutionType),
		Args: graphql.FieldConfigArgument{
			"course_exam_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
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

			param := api.AllCourseExamSolutionParam{
				CourseExamID: courseExamID,
				OrderBy:      p.Args["order_by"].(string),
				OrderDir:     p.Args["order_dir"].(string),
				Offset:       p.Args["offset"].(int),
				Limit:        p.Args["limit"].(int),
			}

			all, err := courseExamSolutionModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		course_exam_solution_detail(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			course_exam_id,
			course_exam_answer_id
		}
	} */

	courseExamSolutionDetailField = &graphql.Field{
		Type: courseExamSolutionType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseExamSolutionResponse{}, errUUID
			}

			data, err := courseExamSolutionModule.One(ctx, api.OneCourseExamSolutionParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_exam_solution_register(
			course_exam_id:" ",
			course_exam_answer_id:" "
		)
		{
			id,
			course_exam_id,
			course_exam_answer_id
		}
	} */

	courseExamSolutionCreateField = &graphql.Field{
		Type: courseExamSolutionType,
		Args: graphql.FieldConfigArgument{
			"course_exam_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_exam_answer_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseExamID, errUUID := uuid.FromString(p.Args["course_exam_id"].(string))
			if errUUID != nil {
				return model.CourseExamSolutionResponse{}, errUUID
			}

			courseExamAnswerID, errUUID := uuid.FromString(p.Args["course_exam_answer_id"].(string))
			if errUUID != nil {
				return model.CourseExamSolutionResponse{}, errUUID
			}

			courseExamAnswer := api.AddCourseExamSolutionParam{
				CourseExamID:       courseExamID,
				CourseExamAnswerID: courseExamAnswerID,
			}

			data, err := courseExamSolutionModule.Add(ctx, courseExamAnswer)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
