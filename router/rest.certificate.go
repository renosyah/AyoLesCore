package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renosyah/AyoLesCore/api"
)

func HandleCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hash := mux.Vars(r)["hash_id"]

	flag := r.FormValue("print")

	cert, err := classRoomCertificateModule.One(ctx, api.OneClassRoomCertificateParam{
		HashID: hash,
	})
	if err != nil {
		fmt.Fprintln(w, "Certificate not found")
		return
	}

	classroom, _ := classRoomModule.One(ctx, api.OneClassRoomParam{
		ID: cert.ClassroomID,
	})

	course, _ := courseModule.One(ctx, api.OneCourseParam{
		ID: classroom.Course.ID,
	})

	student, _ := studentModule.One(ctx, api.OneStudentParam{
		ID: classroom.StudentID,
	})

	data := map[string]string{
		"Name":       student.Name,
		"CourseName": course.CourseName,
		"Date":       cert.CreateAt.Format("02 January 2006"),
		"HashID":     cert.HashID,
		"Print" : flag,
	}


	errServe := temp.ExecuteTemplate(w, "cert.gohtml", data)
	if errServe != nil {
		fmt.Fprintln(w, "failed to show Certificate")
	}

}
