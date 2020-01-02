package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	courseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "course",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"course_name": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"teacher": &graphql.Field{
				Type: teacherType,
			},
			"category": &graphql.Field{
				Type: categoryType,
			},
			"course_details": &graphql.Field{
				Type: graphql.NewList(courseDetailType),
			},
		},
	})

	/* {
		course_list(
			category_id : "",
			teacher_id : "",
			search_by:"course_name",
			search_value:"",
			order_by:"course_name",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			course_name,
			image_url,
			teacher {id, name, email } ,
			category {id, name, image_url},
			course_details { id,course_id , overview_text, description_text,image_url }
		}
	} */

	courseListField = &graphql.Field{
		Type: graphql.NewList(courseType),
		Args: graphql.FieldConfigArgument{
			"category_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"teacher_id": &graphql.ArgumentConfig{
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

			categoryID, _ := uuid.FromString(p.Args["category_id"].(string))
			teacherID, _ := uuid.FromString(p.Args["teacher_id"].(string))

			param := api.AllCourseParam{
				CategoryID:  categoryID,
				TeacherID:   teacherID,
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := courseModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		course_detail(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			course_name,
			image_url,
			teacher { id, name, email } ,
			category {id, name, image_url},
			course_details { id,course_id , overview_text, description_text,image_url }
		}
	} */

	courseDetailField = &graphql.Field{
		Type: courseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseResponse{}, errUUID
			}

			data, err := courseModule.One(ctx, api.OneCourseParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_register(
			course_name:"data science",
			teacher_id :"73aa9774-5f93-40b4-b510-4e465f97cfcd",
			category_id:"c6fef7b3-3bc1-4068-b00a-b58d0ffdb699",
			image_url : "path/to/image",
		)
		{
			id,
			course_name,
			image_url,
			teacher { id, name, email } ,
			category {id, name, image_url},
			course_details { id,course_id , overview_text, description_text,image_url }
		}
	} */

	courseCreateField = &graphql.Field{
		Type: courseType,
		Args: graphql.FieldConfigArgument{
			"course_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"teacher_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"category_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"image_url": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			teacherID, errUUID := uuid.FromString(p.Args["teacher_id"].(string))
			if errUUID != nil {
				return model.CourseResponse{}, errUUID
			}

			categoryID, errUUID := uuid.FromString(p.Args["category_id"].(string))
			if errUUID != nil {
				return model.CourseResponse{}, errUUID
			}

			course := api.AddCourseParam{
				CourseName: p.Args["course_name"].(string),
				ImageURL:   p.Args["image_url"].(string),
				Teacher: &model.Teacher{
					ID: teacherID,
				},
				Category: &model.Category{
					ID: categoryID,
				},
			}

			data, err := courseModule.Add(ctx, course)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
