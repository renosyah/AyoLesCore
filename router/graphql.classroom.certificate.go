package router

import (
	"crypto/md5"
	"fmt"
	"log"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	classroomCertificateType = graphql.NewObject(graphql.ObjectConfig{
		Name: "classroom_certificate",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"classroom_id": &graphql.Field{
				Type: graphql.String,
			},
			"hash_id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		classroom_certificate_list(
			student_id:"",
			order_by:"create_at",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			classroom_id,
			hash_id
		 }
	} */

	classroomCertificateListField = &graphql.Field{
		Type: graphql.NewList(classroomCertificateType),
		Args: graphql.FieldConfigArgument{
			"student_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"order_by": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"order_dir": &graphql.ArgumentConfig{
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

			studentID, errUUID := uuid.FromString(p.Args["student_id"].(string))
			if errUUID != nil {
				return []model.ClassRoomCertificateResponse{}, errUUID
			}

			param := api.AllClassRoomCertificateParam{
				StudentID: studentID,
				OrderBy:   p.Args["order_by"].(string),
				OrderDir:  p.Args["order_dir"].(string),
				Offset:    p.Args["offset"].(int),
				Limit:     p.Args["limit"].(int),
			}

			all, err := classRoomCertificateModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		classroom_certificate_detail(
			classroom_id: ""
		)
		{
			id,
			classroom_id,
			hash_id
		 }
	} */

	classroomCertificateDetailField = &graphql.Field{
		Type: classroomCertificateType,
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			classroomID, errUUID := uuid.FromString(p.Args["classroom_id"].(string))
			if errUUID != nil {
				return model.ClassRoomCertificateResponse{}, errUUID
			}

			param := api.OneClassRoomCertificateParam{
				ClassroomID: classroomID,
			}

			one, err := classRoomCertificateModule.One(ctx, param)
			if err != nil {
				return one, err
			}

			return one, nil
		},
	}

	/* mutation {
		classroom_certificate_register(
			classroom_id : ""
		)
		{
			id,
			classroom_id,
			hash_id
		 }
	} */

	classroomCertificateRegisterField = &graphql.Field{
		Type: classroomCertificateType,
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			classRoomID, errUUID := uuid.FromString(p.Args["classroom_id"].(string))
			if errUUID != nil {
				return model.ClassRoomCertificateResponse{}, errUUID
			}

			hash := fmt.Sprint(p.Args["classroom_id"].(string), time.Now())

			param := api.AddClassRoomCertificateParam{
				ClassroomID: classRoomID,
				HashID:      fmt.Sprintf("%x", md5.Sum([]byte(hash))),
			}

			data, err := classRoomCertificateModule.Add(ctx, param)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		classroom_certificate_update(
			id : "",
			classroom_id : "",
			hash_id : ""
		)
		{
			id
		 }
	} */

	classroomCertificateUpdateField = &graphql.Field{
		Type: classroomCertificateType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"hash_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.ClassRoomCertificateResponse{}, errUUID
			}

			classRoomID, errUUID := uuid.FromString(p.Args["classroom_id"].(string))
			if errUUID != nil {
				return model.ClassRoomCertificateResponse{}, errUUID
			}

			param := api.AddClassRoomCertificateParam{
				ClassroomID: classRoomID,
				HashID:      p.Args["classroom_id"].(string),
			}

			data, err := classRoomCertificateModule.Update(ctx, param, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		classroom_certificate_delete(
			id : ""
		)
		{
			id
		 }
	} */

	classroomCertificateDeleteField = &graphql.Field{
		Type: classroomCertificateType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.ClassRoomCertificateResponse{}, errUUID
			}

			data, err := classRoomCertificateModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
