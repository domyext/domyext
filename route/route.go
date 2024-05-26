package route

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/helper"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		// If it's a preflight request, return early.
		return
	}
	config.SetEnv()

	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/":
		controller.GetHome(w, r)
	case method == "GET" && path == "/refresh/token":
		controller.GetNewToken(w, r)
	case method == "GET" && path == "/data/user":
		controller.GetDataUser(w, r)
	case method == "POST" && path == "/data/user":
		controller.PostDataUser(w, r)
	case method == "GET" && path == "/data/proyek":
		controller.GetDataProject(w, r)
	case method == "GET" && path == "/data/proyek/anggota":
		controller.GetDataMemberProject(w, r)
	case method == "POST" && path == "/data/proyek/anggota":
		controller.PostDataMemberProject(w, r)
	case method == "POST" && path == "/data/proyek":
		controller.PostDataProject(w, r)
	case method == "POST" && helper.URLParam(path, "/webhook/github/:proyek"):
		controller.PostWebHookGithub(w, r)
	case method == "POST" && path == "/notif/ux/postlaporan":
		controller.PostLaporan(w, r)
	case method == "POST" && helper.URLParam(path, "/webhook/nomor/:nomorwa"):
		controller.PostInboxNomor(w, r)
	default:
		controller.NotFound(w, r)
	}
}
