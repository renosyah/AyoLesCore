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
			"nis": &graphql.Field{
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
			nis
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
			nis
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
			nis:"reno@gmail.com",
			password:"12345"
		)
		{
			id,
			name,
			nis
		}
	} */

	studentLoginField = &graphql.Field{
		Type: studentType,
		Args: graphql.FieldConfigArgument{
			"nis": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			nis := p.Args["nis"].(string)
			password := p.Args["password"].(string)

			data, err := studentModule.Login(ctx, api.StudentLoginParam{Nis: nis, Password: password})
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
			nis:"reno@gmail.com",
			password:"12345"
		)
		{
			id,
			name,
			nis
		}
	} */

	studentCreateField = &graphql.Field{
		Type: studentType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"nis": &graphql.ArgumentConfig{
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
				Nis:      p.Args["nis"].(string),
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
			nis:"reno@gmail.com",
			password:"12345"
		)
		{
			id,
			name,
			nis
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
			"nis": &graphql.ArgumentConfig{
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
				Nis:      p.Args["nis"].(string),
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
