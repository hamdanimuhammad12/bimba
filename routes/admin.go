package routes

import (
	"bimba/models"
	"bimba/sessions"
	"bimba/utils"
	"fmt"
	"net/http"
)

func adminGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	/*fmt.Println(session.Values)*/
	ss := session.Values["USERID"].(int)
	fmt.Println(ss)
	user, err := models.GetUserById(ss)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	url := r.URL.String()
	status, err := models.GetStatusMenu(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)

	utils.ExecuteTemplate(w, "main.html", struct {
		User_v models.User
		Active models.Status_Menu
	}{
		User_v: user,
		Active: status,
	})

}
