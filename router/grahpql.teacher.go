package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	teacherType = graphql.NewObject(graphql.ObjectConfig{
		Name: "teacher",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		teacher_list(
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
			email
		}
	} */

	teacherListField = &graphql.Field{
		Type: graphql.NewList(teacherType),
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
			param := api.AllTeacherParam{
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := teacherModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		teacher_detail(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			name,
			email
		}
	} */

	teacherDetailField = &graphql.Field{
		Type: teacherType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.TeacherResponse{}, errUUID
			}

			data, err := teacherModule.One(ctx, api.OneTeacherParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* {
		teacher_login(
			email:"reno@gmail.com",
			password:"12345"
		)
		{
			id,
			name,
			email
		}
	} */

	teacherLoginField = &graphql.Field{
		Type: teacherType,
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			email := p.Args["email"].(string)
			password := p.Args["password"].(string)

			data, err := teacherModule.Login(ctx, api.TeacherLoginParam{Email: email, Password: password})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		teacher_register(
			name:"reno",
			email:"reno@gmail.com",
			password:"12345"
		)
		{
			id,
			name,
			email
		}
	} */

	teacherCreateField = &graphql.Field{
		Type: teacherType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			teacher := api.AddTeacherParam{
				Name:     p.Args["name"].(string),
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
			}

			data, err := teacherModule.Add(ctx, teacher)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		teacher_update(
			id : "",
			name:"reno",
			email:"reno@gmail.com",
			password:"12345"
		)
		{
			id,
			name,
			email
		}
	} */

	teacherUpdateField = &graphql.Field{
		Type: teacherType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.TeacherResponse{}, errUUID
			}

			teacher := api.UpdateTeacherParam{
				ID:       id,
				Name:     p.Args["name"].(string),
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
			}

			data, err := teacherModule.Update(ctx, teacher)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		teacher_delete(
			id : ""
		)
		{
			id
		}
	} */

	teacherDeleteField = &graphql.Field{
		Type: teacherType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.TeacherResponse{}, errUUID
			}

			data, err := teacherModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
