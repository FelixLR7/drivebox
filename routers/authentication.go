package routers

import (
	"drivebox/controllers"
	"net/http"
)

func SetAuthenticationRoutes() {
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/logout", controllers.LogoutHandler)
	http.HandleFunc("/index", controllers.Authentication(controllers.Homepage))
}
