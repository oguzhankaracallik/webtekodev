package main

import (
	"github.com/gofiber/fiber/v2"
	"oguzodev/controllers"
	"oguzodev/database"
)

func main() {
	database.Connect()
	//database.Conn.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{}, &models.Follower{}, &models.Media{})
	app := fiber.New()

	api := app.Group("/api")

	postR := api.Group("/posts")
	postR.Get("/", controllers.PostC{}.GetPosts)
	postR.Get("/:id", controllers.PostC{}.GetPost)
	postR.Post("/", controllers.PostC{}.CreatePost)
	postR.Put("/:id", controllers.PostC{}.UpdatePost)
	postR.Put("/like/:id", controllers.PostC{}.LikePost)
	postR.Delete("/:id", controllers.PostC{}.DeletePost)

	commentR := api.Group("/comments")
	commentR.Get("/:post-id", controllers.CommentC{}.GetComments)
	commentR.Post("/", controllers.CommentC{}.CreateComment)
	commentR.Delete("/:id", controllers.CommentC{}.DeleteComment)

	userR := api.Group("/users")
	userR.Get("/", controllers.UserC{}.GetUsers)
	userR.Get("/:id", controllers.UserC{}.GetUser)
	userR.Post("/", controllers.UserC{}.CreateUser)
	userR.Post("/follow", controllers.UserC{}.FollowUser)
	userR.Put("/:id", controllers.UserC{}.UpdateUser)
	userR.Delete("/:id", controllers.UserC{}.DeleteUser)

	app.Listen(":3000")
}
