package group

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/repository"
	"database/sql"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) repository.IGroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) Insert(group group.Group) (int, error) {
	return 0, nil
}

func (r *GroupRepository) GetByID(id int) (group.Group, error) {
	return group.Group{}, nil
}

func (r *GroupRepository) InsertContact(contact contact.Contact, id int) error {
	return nil
}
