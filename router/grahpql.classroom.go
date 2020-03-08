package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	classRoomType = graphql.NewObject(graphql.ObjectConfig{
		Name: "class_room",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"course": &graphql.Field{
				Type: courseType,
			},
			"student_id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/* {
		classroom_list(
			student_id : "4252869c-ddd2-466f-8528-e1fe8aff4135",
			search_by:"date_add",
			search_value:"",
			order_by:"date_add",
			order_dir:"asc",
			offset:0,
			limit:10
		)
		{
			id,
			course {
				id,
				course_name,
				image_url,
				teacher {id, name, email } ,
				category {id, name, image_url},
				course_details { id,course_id , overview_text, description_text,image_url }
			},
			student_id
		}
	} */

	classRoomListField = &graphql.Field{
		Type: graphql.NewList(classRoomType),
		Args: graphql.FieldConfigArgument{
			"student_id": &graphql.ArgumentConfig{
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

			StudentID, errUUID := uuid.FromString(p.Args["student_id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			param := api.AllClassRoomParam{
				StudentID:   StudentID,
				SearchBy:    p.Args["search_by"].(string),
				SearchValue: p.Args["search_value"].(string),
				OrderBy:     p.Args["order_by"].(string),
				OrderDir:    p.Args["order_dir"].(string),
				Offset:      p.Args["offset"].(int),
				Limit:       p.Args["limit"].(int),
			}

			all, err := classRoomModule.All(ctx, param)
			if err != nil {
				log.Println(err)
				return all, err
			}

			return all, nil
		},
	}

	/* {
		classroom_detail(
			id : "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			course {
				id,
				course_name,
				image_url,
				teacher {id, name, email } ,
				category {id, name, image_url},
				course_details { id,course_id , overview_text, description_text,image_url }
			},
			student_id
		}
	} */

	classRoomDetailField = &graphql.Field{
		Type: classRoomType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			data, err := classRoomModule.One(ctx, api.OneClassRoomParam{ID: id})
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/*  {
		classroom_detail_by_id (
			course_id : "4252869c-ddd2-466f-8528-e1fe8aff4135",
			student_id : "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			course {
				id,
				course_name,
				image_url,
				teacher {id, name, email } ,
				category {id, name, image_url},
				course_details { id,course_id , overview_text, description_text,image_url }
			},
			student_id
		}
	} */

	classRoomDetailByStudentAndCourseIdField = &graphql.Field{
		Type: classRoomType,
		Args: graphql.FieldConfigArgument{
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"student_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			studentID, errUUID := uuid.FromString(p.Args["student_id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			classRoom := api.OneClassRoomByIdParam{
				CourseID:  courseID,
				StudentID: studentID,
			}

			data, err := classRoomModule.OneByStudentIdAndCourseId(ctx, classRoom)
			if err != nil {
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		classroom_register(
			course_id : "4252869c-ddd2-466f-8528-e1fe8aff4135",
			student_id : "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id,
			course {
				id,
				course_name,
				image_url,
				teacher {id, name, email } ,
				category {id, name, image_url},
				course_details { id,course_id , overview_text, description_text,image_url }
			},
			student_id
		}
	} */

	classRoomCreateField = &graphql.Field{
		Type: classRoomType,
		Args: graphql.FieldConfigArgument{
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"student_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			studentID, errUUID := uuid.FromString(p.Args["student_id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			classRoom := api.AddClassRoomParam{
				CourseID:  courseID,
				StudentID: studentID,
			}

			data, err := classRoomModule.Add(ctx, classRoom)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		classroom_update(
			id : "36bf5614-bc14-4c98-96f4-4dae483b79e5",
			course_id : "4252869c-ddd2-466f-8528-e1fe8aff4135",
			student_id : "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			id
		}
	} */

	classRoomUpdateField = &graphql.Field{
		Type: classRoomType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"student_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			studentID, errUUID := uuid.FromString(p.Args["student_id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			classRoom := api.AddClassRoomParam{
				CourseID:  courseID,
				StudentID: studentID,
			}

			data, err := classRoomModule.Update(ctx, classRoom, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		classroom_delete(
			id : "36bf5614-bc14-4c98-96f4-4dae483b79e5"
		)
		{
			id
		}
	} */

	classRoomDeleteField = &graphql.Field{
		Type: classRoomType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.ClassRoomResponse{}, errUUID
			}

			data, err := classRoomModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
