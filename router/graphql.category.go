package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	categoryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		category_list(
			search_by:"name",
			search_value:"",
			order_by:"name",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			name,
			image_url
		 }
	} */

	categoryListField = &graphql.Field{
		Type: graphql.NewList(categoryType),
		Args: graphql.FieldConfigArgument{
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

			param := api.AllCategoryParam{
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := categoryModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		category_detail(
				id:"2d9a7cd6-0054-47fa-8ac1-1dbed77a9652"
			) {
				id,
				name,
				image_url
			}
	} */

	categoryDetailField = &graphql.Field{
		Type: categoryType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CategoryResponse{}, errUUID
			}

			data, err := categoryModule.One(ctx, api.OneCategoryParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* {
		category_register(
				name : "sport",
				image_url : "data/category/sport.png"
			) {
				id,
				name,
				image_url
			}
	} */

	categoryCreateField = &graphql.Field{
		Type: categoryType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"image_url": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			category := api.AddCategoryParam{
				Name:     p.Args["name"].(string),
				ImageURL: p.Args["image_url"].(string),
			}

			data, err := categoryModule.Add(ctx, category)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		category_update(
				id : "",
				name : "sport",
				image_url : "data/category/sport.png"
			) {
				id
			}
	} */

	categoryUpdateField = &graphql.Field{
		Type: categoryType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.ArgumentConfig{
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
				return model.CategoryResponse{}, errUUID
			}

			category := api.AddCategoryParam{
				Name:     p.Args["name"].(string),
				ImageURL: p.Args["image_url"].(string),
			}

			data, err := categoryModule.Update(ctx, category, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		category_update(
				id : "",
				name : "sport",
				image_url : "data/category/sport.png"
			) {
				id
			}
	} */

	categoryDeleteField = &graphql.Field{
		Type: categoryType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CategoryResponse{}, errUUID
			}

			data, err := categoryModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
