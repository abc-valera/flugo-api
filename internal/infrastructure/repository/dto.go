package repository

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
)

// dbInsertUser represents user data which should be added into the database
type dbInsertUser struct {
	Username       string `db:"username"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	Fullname       string `db:"fullname"`
	Status         string `db:"status"`
	Bio            string `db:"bio"`
}

func newDBInsertUser(user *domain.User) *dbInsertUser {
	return &dbInsertUser{
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname,
		Status:         user.Status,
		Bio:            user.Bio,
	}
}

// dbReturnUser represents user data which is returned from the database
type dbReturnUser struct {
	Username       string    `db:"username"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
	Fullname       string    `db:"fullname"`
	Status         string    `db:"status"`
	Bio            string    `db:"bio"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func newDomainUser(user *dbReturnUser) *domain.User {
	return &domain.User{
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname,
		Status:         user.Status,
		Bio:            user.Bio,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

// dbReturnUsers represents slice of dbReturnUser type returned from the database
type dbReturnUsers []*dbReturnUser

func newDomainUsers(dbUsers dbReturnUsers) domain.Users {
	users := make(domain.Users, len(dbUsers))
	for i, user := range dbUsers {
		users[i] = newDomainUser(user)
	}
	return users
}

// dbInsertJoke represents joke data which should be added into the database
type dbInsertJoke struct {
	Username    string `db:"username"`
	Title       string `db:"title"`
	Text        string `db:"text"`
	Explanation string `db:"explanation"`
}

func newDBInsertJoke(joke *domain.Joke) *dbInsertJoke {
	return &dbInsertJoke{
		Username:    joke.Username,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
	}
}

// dbReturnJoke represents joke data which is returned from the database
type dbReturnJoke struct {
	ID          int       `db:"id"`
	Username    string    `db:"username"`
	Title       string    `db:"title"`
	Text        string    `db:"text"`
	Explanation string    `db:"explanation"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func newDomainJoke(joke *dbReturnJoke) *domain.Joke {
	return &domain.Joke{
		ID:          joke.ID,
		Username:    joke.Username,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
		CreatedAt:   joke.CreatedAt,
		UpdatedAt:   joke.UpdatedAt,
	}
}

// dbReturnJokes represents slice of dbReturnJoke type returned from the database
type dbReturnJokes []*dbReturnJoke

func newDomainJokes(dbJokes dbReturnJokes) domain.Jokes {
	jokes := make(domain.Jokes, len(dbJokes))
	for i, joke := range dbJokes {
		jokes[i] = newDomainJoke(joke)
	}
	return jokes
}

// dbInsertLike represents like data which should be added into the database
type dbInsertLike struct {
	Username string `db:"username"`
	JokeID   int    `db:"joke_id"`
}

func newDBInsertLike(like *domain.Like) *dbInsertLike {
	return &dbInsertLike{
		Username: like.Username,
		JokeID:   like.JokeID,
	}
}

// dbReturnLike represents like data which is returned from the database
type dbReturnLike struct {
	Username  string    `db:"username"`
	JokeID    int       `db:"joke_id"`
	CreatedAt time.Time `db:"created_at"`
}

func newDomainLike(like *dbReturnLike) *domain.Like {
	return &domain.Like{
		Username:  like.Username,
		JokeID:    like.JokeID,
		CreatedAt: like.CreatedAt,
	}
}

// dbReturnLikes represents slice of dbReturnLike type returned from the database
type dbReturnLikes []*dbReturnLike

func newDomainLikes(dbLikes dbReturnLikes) domain.Likes {
	likes := make(domain.Likes, len(dbLikes))
	for i, like := range dbLikes {
		likes[i] = newDomainLike(like)
	}
	return likes
}

// dbInsertComment represents comment data which should be added into the database
type dbInsertComment struct {
	Username string `db:"username"`
	JokeID   int    `db:"joke_id"`
	Text     string `db:"text"`
}

func newDBInsertComment(comment *domain.Comment) *dbInsertComment {
	return &dbInsertComment{
		Username: comment.Username,
		JokeID:   comment.JokeID,
		Text:     comment.Text,
	}
}

// dbReturnComment represents comment data which is returned from the database
type dbReturnComment struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	JokeID    int       `db:"joke_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func newDomainComment(comment *dbReturnComment) *domain.Comment {
	return &domain.Comment{
		ID:        comment.ID,
		Username:  comment.Username,
		JokeID:    comment.JokeID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}

// dbReturnComments represents slice of dbReturnComment type returned from the database
type dbReturnComments []*dbReturnComment

func newDomainComments(dbComments dbReturnComments) domain.Comments {
	comments := make(domain.Comments, len(dbComments))
	for i, comment := range dbComments {
		comments[i] = newDomainComment(comment)
	}
	return comments
}
