package controller

import (
	"encoding/json"
	"net/http"

	"github.com/msrshahrukh100/go-webservice/entity"
	"github.com/msrshahrukh100/go-webservice/service"
)

var (
	posts []entity.Post
)

type controller struct{}

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

// var repo repository.PostRepository = repository.NewPostRepository()
var postService service.PostService = service.NewPostService()

func init() {
	posts = []entity.Post{entity.Post{Id: 1, Title: "Shahrukh", Text: "Is awesome"}}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling json"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling json"}`))
		return
	}
	// post.Id = len(posts) + 1
	// posts = append(posts, post)
	postService.Validate(&post)
	postService.Create(&post)
	resp.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(post)
	resp.Write(result)

}
