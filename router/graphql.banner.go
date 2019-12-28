package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	bannerType = graphql.NewObject(graphql.ObjectConfig{
		Name: "banner",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
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
		banner_list(
			search_by:"title",
			search_value:"",
			order_by:"title",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			title,
			content,
			image_url
		 }
	} */

	bannerListField = &graphql.Field{
		Type: graphql.NewList(bannerType),
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

			param := api.AllBannerParam{
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := bannerModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		banner_detail(
				id:"2d9a7cd6-0054-47fa-8ac1-1dbed77a9652"
			) {
				id,
				title,
				content,
				image_url
			}
	} */

	bannerDetailField = &graphql.Field{
		Type: bannerType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.BannerModelResponse{}, errUUID
			}

			data, err := bannerModule.One(ctx, api.OneBannerParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
	/* {
		banner_register(
				title : ,
				content : ,
				image_url:
			) {
				id,
				title,
				content,
				image_url
			}
	} */

	bannerCreateField = &graphql.Field{
		Type: categoryType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
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

			banner := api.AddBannerParam{
				Title:    p.Args["title"].(string),
				Content:  p.Args["content"].(string),
				ImageURL: p.Args["image_url"].(string),
			}

			data, err := bannerModule.Add(ctx, banner)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
