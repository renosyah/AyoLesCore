package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	courseDetailType = graphql.NewObject(graphql.ObjectConfig{
		Name: "course_detail",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"course_id": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"overview_text": &graphql.Field{
				Type: graphql.String,
			},
			"description_text": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		course_detail_list(
			course_id : "",
			search_by:"overview_text",
			search_value:"",
			order_by:"overview_text",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			course_id,
			overview_text,
			description_text,
			image_url
		}
	} */

	courseDetailListField = &graphql.Field{
		Type: graphql.NewList(courseDetailType),
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
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseID, _ := uuid.FromString(p.Args["course_id"].(string))

			param := api.AllCourseDetailParam{
				CourseID:    courseID,
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := courseDetailModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* mutation {
		course_detail_register(
			course_id : "123a1b1e-b822-4035-b9ef-ee133857939f",
			overview_text:"data science",
			description_text : "data science is fun",
			image_url : "path/to/image"
		)
		{
			id,
			course_id,
			overview_text,
			description_text,
			image_url
		}
	} */

	courseDetailCreateField = &graphql.Field{
		Type: courseDetailType,
		Args: graphql.FieldConfigArgument{
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"overview_text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description_text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"image_url": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context

			CourseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.CourseDetailResponse{}, errUUID
			}

			course := api.AddCourseDetailParam{
				CourseID:        CourseID,
				OverviewText:    p.Args["overview_text"].(string),
				DescriptionText: p.Args["description_text"].(string),
				ImageURL:        p.Args["image_url"].(string),
			}

			data, err := courseDetailModule.Add(ctx, course)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_detail_update(
			id:"",
			course_id : "123a1b1e-b822-4035-b9ef-ee133857939f",
			overview_text:"data science",
			description_text : "data science is fun",
			image_url : "path/to/image"
		)
		{
			id
		}
	} */

	courseDetailUpdateField = &graphql.Field{
		Type: courseDetailType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"overview_text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description_text": &graphql.ArgumentConfig{
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
				return model.CourseDetailResponse{}, errUUID
			}

			CourseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.CourseDetailResponse{}, errUUID
			}

			course := api.AddCourseDetailParam{
				CourseID:        CourseID,
				OverviewText:    p.Args["overview_text"].(string),
				DescriptionText: p.Args["description_text"].(string),
				ImageURL:        p.Args["image_url"].(string),
			}

			data, err := courseDetailModule.Update(ctx, course, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_detail_delete(
			id:""
		)
		{
			id
		}
	} */

	courseDetailDeleteField = &graphql.Field{
		Type: courseDetailType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseDetailResponse{}, errUUID
			}

			data, err := courseDetailModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
