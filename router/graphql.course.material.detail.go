package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	courseMaterialDetailType = graphql.NewObject(graphql.ObjectConfig{
		Name: "course_material_detail",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"course_material_id": &graphql.Field{
				Type: graphql.String,
			},
			"position": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"type_material": &graphql.Field{
				Type: graphql.Int,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		course_material_detail_list(
			course_material_id : "",
			search_by:"title",
			search_value:"",
			order_by:"position",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			course_material_id,
			position,
			title,
			type_material,
			content,
			image_url
		}
	} */

	courseMaterialDetailListField = &graphql.Field{
		Type: graphql.NewList(courseMaterialDetailType),
		Args: graphql.FieldConfigArgument{
			"course_material_id": &graphql.ArgumentConfig{
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

			courseMaterialID, _ := uuid.FromString(p.Args["course_material_id"].(string))

			param := api.AllCourseMaterialDetailParam{
				CourseMaterialID: courseMaterialID,
				SearchBy:         p.Args["search_by"].(string),
				SearchValue:      p.Args["search_value"].(string),
				OrderBy:          p.Args["order_by"].(string),
				OrderDir:         p.Args["order_dir"].(string),
				Offset:           p.Args["offset"].(int),
				Limit:            p.Args["limit"].(int),
			}

			all, err := courseMaterialDetailModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* mutation {
		course_material_detail_register(
			course_material_id : "",
			position : 0,
			title  : "",
			type_material : 0,
			content  : "",
			image_url  : "",
		)
		{
			id,
			course_material_id,
			position,
			title,
			type_material,
			content,
			image_url
		}
	} */

	courseMaterialDetailCreateField = &graphql.Field{
		Type: courseMaterialDetailType,
		Args: graphql.FieldConfigArgument{
			"course_material_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"position": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"type_material": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"image_url": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseMaterialID, errUUID := uuid.FromString(p.Args["course_material_id"].(string))
			if errUUID != nil {
				return model.CourseMaterialDetailResponse{}, errUUID
			}

			courseMaterialDetail := api.AddCourseMaterialDetailParam{
				CourseMaterialID: courseMaterialID,
				Position:         int32(p.Args["position"].(int)),
				Title:            p.Args["title"].(string),
				TypeMaterial:     int32(p.Args["type_material"].(int)),
				Content:          p.Args["content"].(string),
				ImageURL:         p.Args["image_url"].(string),
			}

			data, err := courseMaterialDetailModule.Add(ctx, courseMaterialDetail)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_material_detail_update(
			id:"",
			course_material_id : "",
			position : 0,
			title  : "",
			type_material : 0,
			content  : "",
			image_url  : "",
		)
		{
			id
		}
	} */

	courseMaterialDetailUpdateField = &graphql.Field{
		Type: courseMaterialDetailType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_material_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"position": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"type_material": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"content": &graphql.ArgumentConfig{
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
				return model.CourseMaterialDetailResponse{}, errUUID
			}

			courseMaterialID, errUUID := uuid.FromString(p.Args["course_material_id"].(string))
			if errUUID != nil {
				return model.CourseMaterialDetailResponse{}, errUUID
			}

			courseMaterialDetail := api.AddCourseMaterialDetailParam{
				CourseMaterialID: courseMaterialID,
				Position:         int32(p.Args["position"].(int)),
				Title:            p.Args["title"].(string),
				TypeMaterial:     int32(p.Args["type_material"].(int)),
				Content:          p.Args["content"].(string),
				ImageURL:         p.Args["image_url"].(string),
			}

			data, err := courseMaterialDetailModule.Update(ctx, courseMaterialDetail, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_material_detail_delete(
			id:"",
		)
		{
			id
		}
	} */

	courseMaterialDetailDeleteField = &graphql.Field{
		Type: courseMaterialDetailType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseMaterialDetailResponse{}, errUUID
			}

			data, err := courseMaterialDetailModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
