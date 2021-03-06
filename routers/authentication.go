package routers

import (
	"drivebox/controllers"
	"net/http"
)

// SetAuthenticationRoutes ...
func SetAuthenticationRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/login", controllers.LoginHandler)
	mux.HandleFunc("/logout", controllers.Authentication(controllers.LogoutHandler))
	mux.HandleFunc("/upload", controllers.Authentication(func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			controllers.UploadPageHandler(response, request)
		} else {
			controllers.UploadHandler(response, request)
		}
	}))
	mux.HandleFunc("/download", controllers.Authentication(controllers.DownloadHandler))
	mux.HandleFunc("/delete", controllers.Authentication(controllers.DeleteHandler))
	mux.HandleFunc("/validation", controllers.ValidationHandler)
	mux.HandleFunc("/index", controllers.Authentication(controllers.Homepage))

	return mux
}
