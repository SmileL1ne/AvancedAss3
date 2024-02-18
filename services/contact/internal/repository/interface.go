package repository

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type IContactRepository interface {
	Insert(contact.Contact) (int, error)  // returns id of created contact
	GetByID(int) (contact.Contact, error) // accepts contact id
	Update(contact.Contact) error
	Delete(int) error // accepts contact id
}

type IGroupRepository interface {
	Insert(group.Group) (int, error)          // returns id of created group
	GetByID(int) (group.Group, error)         // accepts group id
	InsertContact(contact.Contact, int) error // accepts contact and group id
}
