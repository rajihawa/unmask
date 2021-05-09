package repository

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/rajihawa/unmask/app/domain"
	"github.com/rajihawa/unmask/app/utils"
)

type ProjectMySqlRepo struct {
	db *sql.DB
}

func NewProjectMySqlRepo(db *sql.DB) domain.ProjectRepo {
	return &ProjectMySqlRepo{
		db: db,
	}
}

func (p *ProjectMySqlRepo) GetOne(id string) (*domain.Project, error) {
	projectQuery := squirrel.Select("*").From("projects").Where(squirrel.Eq{"id": id})
	rows, err := projectQuery.RunWith(p.db).Query()
	if err != nil {
		return nil, err
	}
	emptyProject := domain.Project{}
	for rows.Next() {
		// err := rows.Scan(&emptyProject.ID, &emptyProject.Name, &emptyProject.Description, &emptyProject.UserCount, &emptyProject.CreatedAt, &emptyProject.UpdatedAt)
		err := utils.ScanStruct(rows, &emptyProject)
		if err != nil {
			return nil, err
		}
	}
	return &emptyProject, nil
}

func (p *ProjectMySqlRepo) GetAll(limit int, offset int) ([]domain.Project, error) {
	projectsQuery := squirrel.Select("*").From("projects").Limit(uint64(limit)).Offset(uint64(offset))
	rows, err := projectsQuery.RunWith(p.db).Query()
	if err != nil {
		return nil, err
	}
	emptyProjects := []domain.Project{}
	for rows.Next() {
		emptyProject := domain.Project{}
		err := utils.ScanStruct(rows, &emptyProject)
		if err != nil {
			return nil, err
		}
		emptyProjects = append(emptyProjects, emptyProject)
	}
	return emptyProjects, nil
}

func (p *ProjectMySqlRepo) UpdateOne(id string, newProject domain.Project) error {
	updateQuery := squirrel.Update("projects").Set("name", newProject.Name).Set("description", newProject.Description).Set("user_count", newProject.UserCount).Where(squirrel.Eq{"id": id})
	_, err := updateQuery.RunWith(p.db).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectMySqlRepo) CreateOne(newProject domain.Project) error {
	createQuery := squirrel.Insert("projects").Columns("id", "name", "description").Values(newProject.ID, newProject.Name, newProject.Description)
	_, err := createQuery.RunWith(p.db).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectMySqlRepo) DeleteOne(id string) error {
	deleteQuery := squirrel.Delete("projects").Where(squirrel.Eq{"id": id})
	_, err := deleteQuery.RunWith(p.db).Exec()
	if err != nil {
		return err
	}
	return nil
}
