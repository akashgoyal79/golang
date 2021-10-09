package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/akashgoyal79/mongo-golang/controllers"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.DELETE("/users/:id", uc.DeleteUser)
	r.POST("/posts", pc.CreatePost)
	r.GET("/posts/:id", pc.GetPostUsingId)
	r.GET("/post/users/:userid", pc.GetAllPostUsingId)

	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
