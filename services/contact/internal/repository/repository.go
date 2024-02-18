package repository

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"database/sql"
)

type ContactRepository struct {
	db *sql.DB
}

func New(db *sql.DB) IContactRepository {
	return &ContactRepository{
		db: db,
	}
}

func (r *ContactRepository) Insert(contact contact.Contact) (int, error) {
	return 0, nil
}

func (r *ContactRepository) GetByID(id int) (contact.Contact, error) {
	return contact.Contact{}, nil
}

func (r *ContactRepository) Update(contact contact.Contact) error {
	return nil
}

func (r *ContactRepository) Delete(id int) error {
	return nil
}

func (r *ContactRepository) InsertGroup(group group.Group) (int, error) {
	return 0, nil
}

func (r *ContactRepository) GetByIDGroup(id int) (group.Group, error) {
	return group.Group{}, nil
}

func (r *ContactRepository) InsertContact(contact contact.Contact, id int) error {
	return nil
}
