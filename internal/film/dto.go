package film

type CreateFilmDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Edition     string `json:"edition"`
	Rating      int8   `json:"rating"`
	ActorID    int    `json:"actor_id"`
}
