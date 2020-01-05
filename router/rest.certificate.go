package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renosyah/AyoLesCore/api"
	"github.com/renosyah/AyoLesCore/util"
)

func HandleCertificateQRcode(w http.ResponseWriter, r *http.Request) {

	hash := mux.Vars(r)["hash_id"]

	qr := &util.Qrcode{
		Value: fmt.Sprintf("http://letscourse.com/cert/%s", hash),
		Size:  256,
	}

	qrByte, err := qr.MakeByte()
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	_, err = w.Write(qrByte)
	if err != nil {
		http.Error(w, "500 failed to serve certificate", http.StatusInternalServerError)
	}
}

func HandleCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hash := mux.Vars(r)["hash_id"]

	flag := r.FormValue("print")

	cert, err := classRoomCertificateModule.One(ctx, api.OneClassRoomCertificateParam{
		HashID: hash,
	})
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
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
		"CourseID":   course.ID.String(),
		"Date":       cert.CreateAt.Format("02 January 2006"),
		"HashID":     cert.HashID,
		"Print":      flag,
	}

	errServe := temp.ExecuteTemplate(w, "cert.gohtml", data)
	if errServe != nil {
		http.Error(w, "500 failed to serve certificate", http.StatusInternalServerError)
	}

}
