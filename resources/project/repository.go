package project

import (
	"data-pad.app/data-api/repository"
)

type ProjectRepository struct {
	repository.MongoRepository
}

func NewProjectRepository() ProjectRepository {
	r := ProjectRepository{}
	r.Collection = "projects"

	return r
}
