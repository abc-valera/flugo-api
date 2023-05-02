package api

import (
	"github.com/abc-valera/flugo-api/internal/delivery/handler"
	"github.com/abc-valera/flugo-api/internal/delivery/middleware"
	"github.com/abc-valera/flugo-api/internal/infrastructure/framework"
	"github.com/gofiber/fiber/v2"
)

func routes(app *fiber.App, frameworks *framework.Frameworks, handlers *handler.Handlers) {
	unauth := app.Group("/")
	unauth.Post("sign_up", handlers.UserHandler.SignUp)
	unauth.Get("sign_in", handlers.UserHandler.SignIn)

	// Auth middleware
	app.Use(middleware.NewAuthMiddleware(frameworks.TokenFramework))

	me := app.Group("/me")
	me.Get("", handlers.UserHandler.GetMe)
	me.Put("/password", handlers.UserHandler.UpdateMyPassword)
	me.Put("/fullname", handlers.UserHandler.UpdateMyFullname)
	me.Put("/status", handlers.UserHandler.UpdateMyStatus)
	me.Put("/bio", handlers.UserHandler.UpdateMyBio)
	me.Delete("", handlers.UserHandler.DeleteMe)

	myJokes := me.Group("/jokes")
	myJokes.Post("", handlers.JokeHandler.CreateMyJoke)
	myJokes.Get("", handlers.JokeHandler.GetMyJokes)
	myJokes.Put("/explanation", handlers.JokeHandler.UpdateMyJokeExplanation)
	myJokes.Delete("/:joke_id", handlers.JokeHandler.DeleteMyJoke)

	myLikes := me.Group("/likes")
	myLikes.Post("/:joke_id", handlers.LikeHandler.CreateMyLike)
	myLikes.Get("", handlers.LikeHandler.MyLikedJokes)
	myLikes.Delete("/:joke_id", handlers.LikeHandler.DeleteMyLike)

	myComments := me.Group("/comments")
	myComments.Post("", handlers.CommentHandler.NewMyComment)
	myComments.Delete("/:comment_id", handlers.CommentHandler.DeleteMyComment)

	users := app.Group("/users")
	users.Get("/search", handlers.UserHandler.SearchUsersByUsername) // Note: Not implemented

	jokes := app.Group("/jokes")
	jokes.Get("", handlers.JokeHandler.GetAllJokes)
	jokes.Get("/:joke_id", handlers.JokeHandler.GetJoke)
	jokes.Get("/by/:username", handlers.JokeHandler.GetUserJokes)
	jokes.Get("/search", handlers.JokeHandler.SearchJokesByTitle) // Note: Not implemented

	likes := app.Group("/likes")
	likes.Get("/:joke_id", handlers.LikeHandler.JokeLikes)
	likes.Get("/:joke_id/users", handlers.LikeHandler.UsersWhoLikedJoke)
	likes.Get("/:username/jokes", handlers.LikeHandler.JokesUserLiked)

	comments := app.Group("/comments")
	comments.Get("/:joke_id", handlers.CommentHandler.GetJokeComments)
}
