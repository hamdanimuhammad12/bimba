package routes

import (
	"lakuinv5dash/auth"
	"lakuinv5dash/models"
	"lakuinv5dash/sessions"
	"lakuinv5dash/utils"
	"net/http"

	/*"log"*/
	"fmt"
)

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(w, r)
	utils.ExecuteTemplate(w, "login.html", struct {
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	user, err := auth.Signin(email, password)
	checkErrAuthenticate(err, w, r, user)
	fmt.Println(auth.Signin(email, password))
}

func checkErrAuthenticate(err error, w http.ResponseWriter, r *http.Request, user models.User) {
	session, _ := sessions.Store.Get(r, "session")
	message := "Success"
	if err != nil {
		switch err {
		case auth.ErrInvalidEmail,
			auth.ErrInvalidPassword:
			message = fmt.Sprintf("%s", err)
			break
		default:
			utils.InternalServerError(w)
			return
		}
		session.Values["MESSAGE"] = message
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}
	session.Values["MESSAGE"] = message
	session.Values["ALERT"] = "success"
	session.Values["USERID"] = user.Id
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)

}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "USERID")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}
