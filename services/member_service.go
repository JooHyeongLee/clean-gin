package services

import (
	"github.com/JooHyeongLee/clean-gin/domains"
	"github.com/JooHyeongLee/clean-gin/lib"
	"github.com/JooHyeongLee/clean-gin/models"
	"github.com/JooHyeongLee/clean-gin/repository"
)

// MemberService service layer
type MemberService struct {
	logger     lib.Logger
	repository repository.MemberRepository
}

// NewMemberService creates a new userservice
func NewMemberService(logger lib.Logger, repository repository.MemberRepository) domains.MemberService {
	return MemberService{
		logger:     logger,
		repository: repository,
	}
}

// GetOneMember gets one user
func (s MemberService) GetOneMember(id uint) (member models.Member, err error) {
	return member, s.repository.Find(&member, id).Error
}
