package contact

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/repository"
	usecase "architecture_go/services/contact/internal/useCase"
)

type ContactUsecase struct {
	contactRepo repository.IContactRepository
}

func NewContactUsecase(repo repository.IContactRepository) usecase.IContactUsecase {
	return &ContactUsecase{
		contactRepo: repo,
	}
}

func (cs *ContactUsecase) CreateContact(contact contact.Contact) (int, error) {
	return 0, nil
}

func (cs *ContactUsecase) GetContact(id int) (contact.Contact, error) {
	return contact.Contact{}, nil
}

func (cs *ContactUsecase) UpdateContact(contact contact.Contact) error {
	return nil
}

func (cs *ContactUsecase) DeleteContact(int) error {
	return nil
}
