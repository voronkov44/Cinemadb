package film

type CreateFilmDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Edition     string `json:"edition"`
	Rating      int8   `json:"rating"`
	AuthorID    int    `json:"author_id"`
}
