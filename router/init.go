package router

import (
	"context"
	"database/sql"
	"encoding/json"
	template "html/template"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/renosyah/AyoLesCore/api"
)

var (
	studentModule                *api.StudentModule
	categoryModule               *api.CategoryModule
	bannerModule                 *api.BannerModule
	teacherModule                *api.TeacherModule
	courseModule                 *api.CourseModule
	courseDetailModule           *api.CourseDetailModule
	classRoomModule              *api.ClassRoomModule
	courseMaterialModule         *api.CourseMaterialModule
	courseMaterialDetailModule   *api.CourseMaterialDetailModule
	classRoomProgressModule      *api.ClassRoomProgressModule
	courseExamModule             *api.CourseExamModule
	courseExamAnswerModule       *api.CourseExamAnswerModule
	classRoomExamProgressModule  *api.ClassRoomExamProgressModule
	classRoomExamResultModule    *api.ClassRoomExamResultModule
	classRoomCertificateModule   *api.ClassRoomCertificateModule
	courseQualificationModule    *api.CourseQualificationModule
	classRoomQualificationModule *api.ClassRoomQualificationModule
	courseExamSolutionModule     *api.CourseExamSolutionModule
	db                           *sql.DB
	temp                         *template.Template
)

func Init(d *sql.DB) {
	db = d
	temp = template.Must(template.ParseGlob("template/*.gohtml"))
	studentModule = api.NewStudentModule(db)
	categoryModule = api.NewCategoryModule(db)
	bannerModule = api.NewBannerModule(db)
	teacherModule = api.NewTeacherModule(db)
	courseModule = api.NewCourseModule(db)
	courseDetailModule = api.NewCourseDetailModule(db)
	classRoomModule = api.NewClassRoomModule(db)
	courseMaterialModule = api.NewCourseMaterialModule(db)
	courseMaterialDetailModule = api.NewCourseMaterialDetailModule(db)
	classRoomProgressModule = api.NewClassRoomProgressModule(db)
	courseExamModule = api.NewCourseExamModule(db)
	courseExamAnswerModule = api.NewCourseExamAnswerModule(db)
	classRoomExamProgressModule = api.NewClassRoomExamProgressModule(db)
	classRoomExamResultModule = api.NewClassRoomExamResultModule(db)
	classRoomCertificateModule = api.NewClassRoomCertificateModule(db)
	courseQualificationModule = api.NewCourseQualificationModule(db)
	classRoomQualificationModule = api.NewClassRoomQualificationModule(db)
	courseExamSolutionModule = api.NewCourseExamSolutionModule(db)
}

// ParseBodyData parse json-formatted request body into given struct.
func ParseBodyData(ctx context.Context, r *http.Request, data interface{}) error {
	bBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "read")
	}

	err = json.Unmarshal(bBody, data)
	if err != nil {
		return errors.Wrap(err, "json")
	}

	valid, err := govalidator.ValidateStruct(data)
	if err != nil {
		return errors.Wrap(err, "validate")
	}

	if !valid {
		return errors.New("invalid data")
	}

	return nil
}
