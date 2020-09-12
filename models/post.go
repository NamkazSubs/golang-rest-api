package models

import (
	"fmt"

	u "github.com/NamkazSubs/golang-rest-api/utils"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `gorm:"type:text"|json:"content"`
	UserID  uint
	// User    User `gorm:"foreignKey:UserID"`
}

/*
Validasi
*/
func (post *Post) Validate() (map[string]interface{}, bool) {

	if post.Title == "" {
		return u.Message(false, "Title Wajib Diisi"), false
	}

	if post.Slug == "" {
		return u.Message(false, "Slug Wajib Diisi"), false
	}

	if post.Content == "" {
		return u.Message(false, "Content Wajib Diisi"), false
	}

	if post.UserID <= 0 {
		return u.Message(false, "User Invalid"), false
	}

	return u.Message(true, "success"), true
}

//

func (post *Post) ValidateUpdate() (map[string]interface{}, bool) {

	if post.Title == "" {
		return u.Message(false, "Title Wajib Diisi"), false
	}

	if post.Slug == "" {
		return u.Message(false, "Slug Wajib Diisi"), false
	}

	if post.Content == "" {
		return u.Message(false, "Content Wajib Diisi"), false
	}

	if post.ID <= 0 {
		return u.Message(false, "Invalid"), false
	}

	if post.UserID <= 0 {
		return u.Message(false, "User Invalid"), false
	}

	return u.Message(true, "success"), true
}

//

func (post *Post) Validate_Delete() (map[string]interface{}, bool) {

	if post.ID <= 0 {
		return u.Message(false, "User Invalid"), false
	}
	if post.UserID <= 0 {
		return u.Message(false, "User Invalid"), false
	}

	return u.Message(true, "success"), true
}

//
func AllPost() []*Post {

	post := make([]*Post, 0)
	err := GetDB().Table("posts").Find(&post).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return post
}

func (post *Post) Create() map[string]interface{} {

	if resp, ok := post.Validate(); !ok {
		return resp
	}

	GetDB().Create(post)

	resp := u.Message(true, "success")
	resp["post"] = post
	return resp
}

func GetPost(id uint) *Post {
	fmt.Println(id)

	post := &Post{}
	err := GetDB().Table("posts").Where("id = ?", id).First(post).Error
	if err != nil {
		return nil
	}
	return post
}

func GetPostByUser(user uint) []*Post {

	post := make([]*Post, 0)
	err := GetDB().Table("posts").Where("user_id = ?", user).Find(&post).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return post
}

func (post *Post) Update() map[string]interface{} {
	if resp, ok := post.ValidateUpdate(); !ok {
		return resp
	}

	GetDB().Update(post)

	resp := u.Message(true, "success")
	resp["post"] = post
	return resp
}

func (post *Post) Delete() map[string]interface{} {
	if resp, ok := post.Validate_Delete(); !ok {
		return resp
	}

	GetDB().Delete(post)

	resp := u.Message(true, "success")
	return resp
}
