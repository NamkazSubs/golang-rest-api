package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NamkazSubs/golang-rest-api/models"
	u "github.com/NamkazSubs/golang-rest-api/utils"
)

var AllPost = func(w http.ResponseWriter, r *http.Request) {

	data := models.AllPost()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
	// fmt.Fprintf(w, "Rest API With GoLang")

}

var CreatePost = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	post := &models.Post{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	post.UserID = user
	resp := post.Create()
	u.Respond(w, resp)
}

var GetPost = func(w http.ResponseWriter, r *http.Request) {

	// params := mux.Vars(r)
	post := &models.Post{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := models.GetPost(post.ID)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdatePost = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	post := &models.Post{}

	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	post.UserID = user
	resp := post.Update()
	u.Respond(w, resp)

}

var DeletePost = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	post := &models.Post{}

	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	post.UserID = user
	resp := post.Delete()
	u.Respond(w, resp)

}
