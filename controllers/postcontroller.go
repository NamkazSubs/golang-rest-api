package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetPost(uint(id))
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
