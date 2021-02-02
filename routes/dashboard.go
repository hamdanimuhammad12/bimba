package routes

import (
	"bimba/utils"
	"net/http"
	/*"lakuin/models"*/ /*"fmt"*/)

func dashboardGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "dashboard.html", nil)
}

func kwitanGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "kwitan.html", nil)
}
