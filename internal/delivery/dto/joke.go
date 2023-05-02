package dto

import "github.com/abc-valera/flugo-api/internal/domain"

type JokeResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Explanation string `json:"explanation"`
}

func NewJokeResponse(joke *domain.Joke) *JokeResponse {
	return &JokeResponse{
		ID:          joke.ID,
		Username:    joke.Username,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
	}
}

type JokesResponse []*JokeResponse

func NewJokesResponse(jokes domain.Jokes) JokesResponse {
	jokesResponse := make(JokesResponse, len(jokes))
	for i, joke := range jokes {
		jokesResponse[i] = NewJokeResponse(joke)
	}
	return jokesResponse
}

type CreateMyJokeRequest struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	Explanation string `json:"explanation"`
}

type UpdateMyJokeExplanationRequest struct {
	JokeID      int    `json:"joke_id"`
	Explanation string `json:"explanation"`
}
