package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/repository"
)

type ContactUsecase struct {
	contactRepo repository.IContactRepository
}

func New(repo repository.IContactRepository) IContactUsecase {
	return &ContactUsecase{
		contactRepo: repo,
	}
}

func (cs *ContactUsecase) CreateContact(contact contact.Contact) (int, error) {
	return cs.contactRepo.Insert(contact)
}

func (cs *ContactUsecase) GetContact(id int) (contact.Contact, error) {
	return cs.contactRepo.GetByID(id)
}

func (cs *ContactUsecase) UpdateContact(contact contact.Contact) error {
	return cs.contactRepo.Update(contact)
}

func (cs *ContactUsecase) DeleteContact(id int) error {
	return cs.contactRepo.Delete(id)
}

func (cs *ContactUsecase) CreateGroup(group group.Group) (int, error) {
	return cs.contactRepo.InsertGroup(group)
}

func (cs *ContactUsecase) GetGroup(id int) (group.Group, error) {
	return cs.contactRepo.GetByIDGroup(id)
}

func (cs *ContactUsecase) InsertContact(contact contact.Contact, groupID int) error {
	return cs.contactRepo.InsertContact(contact, groupID)
}
