package film

import entity "github.com/wonderf00l/filmLib/internal/entity/film"

func filmFromServiceToRepository(film *entity.Film) (*entity.Film, []string) {
	actors := make([]string, len(film.Actors))
	for actor := range film.Actors {
		actors = append(actors, actor)
	}
	return film, actors
}
