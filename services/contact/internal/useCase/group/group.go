package group

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/repository"
	usecase "architecture_go/services/contact/internal/useCase"
)

type GroupUsecase struct {
	gropRepo repository.IGroupRepository
}

func NewGroupUseCase(repo repository.IGroupRepository) usecase.IGroupUsecase {
	return &GroupUsecase{
		gropRepo: repo,
	}
}

func (gs *GroupUsecase) CreateGroup(group.Group) (int, error) {
	return 0, nil
}

func (gs *GroupUsecase) GetGroup() (group.Group, error) {
	return group.Group{}, nil
}

func (gs *GroupUsecase) InsertContact(contact contact.Contact, id int) error {
	return nil
}
