package view

import (
	"net/http"
	"net/url"

	"github.com/acoshift/gonews/pkg/model"
)

type IndexData struct {
	List []*model.News
}

// Index renders index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// News renders news view
func News(w http.ResponseWriter, data *model.News) {
	render(tpNews, w, data)
}

type AdminLoginData struct {
	Flash url.Values
}

// AdminLogin renders admin login view
func AdminLogin(w http.ResponseWriter, data *AdminLoginData) {
	render(tpAdminLogin, w, data)
	data.Flash.Del("errors")
}

// AdminRegister renders admin login view
func AdminRegister(w http.ResponseWriter, data interface{}) {
	render(tpAdminRegister, w, data)
}

type AdminListData struct {
	List []*model.News
}

// AdminList renders admin login view
func AdminList(w http.ResponseWriter, data *AdminListData) {
	render(tpAdminList, w, data)
}

// AdminCreate renders admin login view
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit renders admin login view
func AdminEdit(w http.ResponseWriter, data *model.News) {
	render(tpAdminEdit, w, data)
}
