package domains

import (
	"github.com/JooHyeongLee/clean-gin/models"
)

type MemberService interface {
	GetOneMember(id uint) (models.Member, error)
}
