package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	courseMaterialType = graphql.NewObject(graphql.ObjectConfig{
		Name: "course_material",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"course_id": &graphql.Field{
				Type: graphql.String,
			},
			"material_index": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		course_material_list(
			course_id : "",
			search_by:"title",
			search_value:"",
			order_by:"material_index",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			course_id,
			material_index,
			title
		}
	} */

	courseMaterialListField = &graphql.Field{
		Type: graphql.NewList(courseMaterialType),
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

			param := api.AllCourseMaterialParam{
				CourseID:    courseID,
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := courseMaterialModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		course_material_detail(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			course_id,
			material_index,
			title
		}
	} */

	courseMaterialDetailField = &graphql.Field{
		Type: courseMaterialType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseMaterialResponse{}, errUUID
			}

			data, err := courseMaterialModule.One(ctx, api.OneCourseMaterialParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_material_register(
			course_id : "4252869c-ddd2-466f-8528-e1fe8aff4135",
			material_index : 1,
			title : "Chapter 1"
		)
		{
			id,
			course_id,
			material_index,
			title
		}
	} */

	courseMaterialCreateField = &graphql.Field{
		Type: courseType,
		Args: graphql.FieldConfigArgument{
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"material_index": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.CourseMaterialResponse{}, errUUID
			}

			courseMaterial := api.AddCourseMaterialParam{
				CourseID:      courseID,
				MaterialIndex: int32(p.Args["material_index"].(int)),
				Title:         p.Args["title"].(string),
			}

			data, err := courseMaterialModule.Add(ctx, courseMaterial)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
