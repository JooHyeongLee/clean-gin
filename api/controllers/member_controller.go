package controllers

import (
	"github.com/JooHyeongLee/clean-gin/domains"
	"github.com/JooHyeongLee/clean-gin/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// MemberController data type
type MemberController struct {
	service domains.MemberService
	logger  lib.Logger
}

// NewMemberController creates new user controller
func NewMemberController(memberService domains.MemberService, logger lib.Logger) MemberController {
	return MemberController{
		service: memberService,
		logger:  logger,
	}
}

// GetOneMember gets one user
func (u MemberController) GetOneMember(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	member, err := u.service.GetOneMember(uint(id))

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": member,
	})

}
