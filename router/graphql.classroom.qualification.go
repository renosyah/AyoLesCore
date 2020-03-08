package router

import (
	"github.com/graphql-go/graphql"
	"github.com/renosyah/AyoLesCore/api"
	uuid "github.com/satori/go.uuid"
)

var (
	classRoomQualificationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "classroom_qualification",
		Fields: graphql.Fields{
			"classroom_id": &graphql.Field{
				Type: graphql.String,
			},
			"course_qualification": &graphql.Field{
				Type: courseQualificationType,
			},
			"total_score": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

	/* {
		class_qualification_detail(
			classroom_id: "4252869c-ddd2-466f-8528-e1fe8aff4135"
		)
		{
			classroom_id,
			total_score,
			status,
			course_qualification {
				id,
				course_id,
				course_level,
				min_score,
				course_material_total,
				course_exam_total
			}
		}
	} */

	classQualificationDetailField = &graphql.Field{
		Type: classRoomQualificationType,
		Args: graphql.FieldConfigArgument{
			"classroom_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			ctx := p.Context

			classRoomid, _ := uuid.FromString(p.Args["classroom_id"].(string))

			data, err := classRoomQualificationModule.One(ctx, api.OneClassRoomQualificationParam{
				ClassRoomID: classRoomid,
			})
			if err != nil {
				return data, err
			}

			return data, nil
		},
	}
)

// ITS DOESNOT HAVE TABLE
// THIS MODEL VALUE RESULT FROM
// QUERY JOIN
// NO UPDATE
// NO DELETE
