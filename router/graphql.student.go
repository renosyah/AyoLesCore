package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	studentType = graphql.NewObject(graphql.ObjectConfig{
		Name: "student",
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
		student_list(
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

	studentListField = &graphql.Field{
		Type: graphql.NewList(studentType),
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
			param := api.AllStudentParam{
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := studentModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		student_detail(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			name,
			email
		}
	} */

	studentDetailField = &graphql.Field{
		Type: studentType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.StudentResponse{}, errUUID
			}

			data, err := studentModule.One(ctx, api.OneStudentParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* {
		student_login(
			email:"reno@gmail.com",
			password:"12345"
		)
		{
			id,
			name,
			email
		}
	} */

	studentLoginField = &graphql.Field{
		Type: studentType,
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

			data, err := studentModule.Login(ctx, api.StudentLoginParam{Email: email, Password: password})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		student_register(
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

	studentCreateField = &graphql.Field{
		Type: studentType,
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

			student := api.AddStudentParam{
				Name:     p.Args["name"].(string),
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
			}

			data, err := studentModule.Add(ctx, student)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		student_update(
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

	studentUpdateField = &graphql.Field{
		Type: studentType,
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
				return model.StudentResponse{}, errUUID
			}

			student := api.UpdateStudentParam{
				ID:       id,
				Name:     p.Args["name"].(string),
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
			}

			data, err := studentModule.Update(ctx, student)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		student_delete(
			id : ""
		)
		{
			id
		}
	} */

	studentDeleteField = &graphql.Field{
		Type: studentType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.StudentResponse{}, errUUID
			}

			data, err := studentModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
