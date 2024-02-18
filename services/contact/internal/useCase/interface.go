package usecase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type IContactUsecase interface {
	CreateContact(contact.Contact) (int, error) // returns id of created contact
	GetContact(int) (contact.Contact, error)
	UpdateContact(contact.Contact) error
	DeleteContact(int) error // accepts contact id
}

type IGroupUsecase interface {
	CreateGroup(group.Group) (int, error) // returns id of created group
	GetGroup() (group.Group, error)
	InsertContact(contact.Contact, int) error // accepts contact and group id
}
