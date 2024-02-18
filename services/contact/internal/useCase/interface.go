package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type IContactUsecase interface {
	CreateContact(contact.Contact) (int, error) // returns id of created contact
	GetContact(int) (contact.Contact, error)    // accepts contact id
	UpdateContact(contact.Contact) error
	DeleteContact(int) error // accepts contact id

	CreateGroup(group.Group) (int, error)     // returns id of created group
	GetGroup(int) (group.Group, error)        // accepts group id
	InsertContact(contact.Contact, int) error // accepts contact and group id
}
