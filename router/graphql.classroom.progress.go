package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	classroomProgressType = graphql.NewObject(graphql.ObjectConfig{
		Name: "classroom_progress",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"classroom_id": &graphql.Field{
				Type: graphql.String,
			},
			"course_material_id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		classroom_progress_list(
			classroom_id:"",
			offset:0,
			limit:10
		)
		{
			id,
			classroom_id,
			course_material_id
		 }
	} */

	classroomProgressListField = &graphql.Field{
		Type: graphql.NewList(classroomProgressType),
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
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

			classRoomID, errUUID := uuid.FromString(p.Args["classroom_id"].(string))
			if errUUID != nil {
				return []model.ClassRoomProgressResponse{}, errUUID
			}

			param := api.AllClassRoomProgressParam{
				ClassRoomID: classRoomID,
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := classRoomProgressModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		classroom_progress_detail(
			id: ""
		)
		{
			id,
			classroom_id,
			course_material_id
		 }
	} */

	classroomProgressDetailField = &graphql.Field{
		Type: graphql.NewList(classroomProgressType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.ClassRoomProgressResponse{}, errUUID
			}

			param := api.OneClassRoomProgressParam{
				ID: id,
			}

			one, err := classRoomProgressModule.One(ctx, param)
			if err != nil {
				log.Println(err)
				return one, err
			}

			return one, nil
		},
	}

	/* mutation {
		classroom_progress_register(
			classroom_id :"",
			course_material_id :""
		)
		{
			id,
			classroom_id,
			course_material_id
		 }
	} */

	classroomProgressRegisterField = &graphql.Field{
		Type: classroomProgressType,
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_material_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			classRoomID, errUUID := uuid.FromString(p.Args["classroom_id"].(string))
			if errUUID != nil {
				return model.ClassRoomProgressResponse{}, errUUID
			}
			courseMaterialID, errUUID := uuid.FromString(p.Args["course_material_id"].(string))
			if errUUID != nil {
				return model.ClassRoomProgressResponse{}, errUUID
			}

			param := api.AddClassRoomProgressParam{
				ClassRoomID:      classRoomID,
				CourseMaterialID: courseMaterialID,
			}

			data, err := classRoomProgressModule.Add(ctx, param)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
