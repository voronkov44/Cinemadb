package film

import "VkProject/internal/actor"

type (
	Film struct {
		ID          string        `json:"id"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Edition     string        `json:"edition"`
		Rating      int8          `json:"rating"`
		Actors      []actor.Actor `json:"actors"`
	}
)
