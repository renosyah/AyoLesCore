package router

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

var (
	courseQualificationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "course_qualification",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"course_id": &graphql.Field{
				Type: graphql.String,
			},
			"course_level": &graphql.Field{
				Type: graphql.String,
			},
			"min_score": &graphql.Field{
				Type: graphql.Int,
			},
			"course_material_total": &graphql.Field{
				Type: graphql.Int,
			},
			"course_exam_total": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

	/* {
		course_qualification_detail(
			id: "4252869c-ddd2-466f-8528-e1fe8aff4135",
			course_id : "",
		)
		{
			id,
			course_id,
			course_level,
			min_score,
			course_material_total,
			course_exam_total
		}
	} */

	courseQualificationDetailField = &graphql.Field{
		Type: courseQualificationType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, _ := uuid.FromString(p.Args["id"].(string))
			courseID, _ := uuid.FromString(p.Args["course_id"].(string))

			data, err := courseQualificationModule.One(ctx, api.OneCourseQualificationParam{
				ID:       id,
				CourseID: courseID,
			})
			if err != nil {
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_qualification_register(
			course_id:"",
			course_level:"",
			min_score:0,
			course_material_total:0,
			course_exam_total:0
		)
		{
			id,
			course_id,
			course_level,
			min_score,
			course_material_total,
			course_exam_total
		}
	} */

	courseQualificationCreateField = &graphql.Field{
		Type: courseQualificationType,
		Args: graphql.FieldConfigArgument{
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_level": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"min_score": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"course_material_total": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"course_exam_total": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.CourseQualificationResponse{}, errUUID
			}

			courseQualification := api.AddCourseQualificationParam{
				CourseID:            courseID,
				CourseLevel:         p.Args["course_level"].(string),
				MinScore:            int32(p.Args["min_score"].(int)),
				CourseMaterialTotal: int32(p.Args["course_material_total"].(int)),
				CourseExamTotal:     int32(p.Args["course_exam_total"].(int)),
			}

			data, err := courseQualificationModule.Add(ctx, courseQualification)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_qualification_update(
			id:"",
			course_id:"",
			course_level:"",
			min_score:0,
			course_material_total:0,
			course_exam_total:0
		)
		{
			id
		}
	} */

	courseQualificationUpdateField = &graphql.Field{
		Type: courseQualificationType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_level": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"min_score": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"course_material_total": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"course_exam_total": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseQualificationResponse{}, errUUID
			}

			courseID, errUUID := uuid.FromString(p.Args["course_id"].(string))
			if errUUID != nil {
				return model.CourseQualificationResponse{}, errUUID
			}

			courseQualification := api.AddCourseQualificationParam{
				CourseID:            courseID,
				CourseLevel:         p.Args["course_level"].(string),
				MinScore:            int32(p.Args["min_score"].(int)),
				CourseMaterialTotal: int32(p.Args["course_material_total"].(int)),
				CourseExamTotal:     int32(p.Args["course_exam_total"].(int)),
			}

			data, err := courseQualificationModule.Update(ctx, courseQualification, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}

	/* mutation {
		course_qualification_delete(
			id:""
		)
		{
			id
		}
	} */

	courseQualificationDeleteField = &graphql.Field{
		Type: courseQualificationType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			id, errUUID := uuid.FromString(p.Args["id"].(string))
			if errUUID != nil {
				return model.CourseQualificationResponse{}, errUUID
			}

			data, err := courseQualificationModule.Delete(ctx, id)
			if err != nil {
				log.Println(err)
				return data, err
			}

			return data, nil
		},
	}
)
