package repository

import "github.com/JooHyeongLee/clean-gin/lib"

type MemberRepository struct {
	lib.Database
	logger lib.Logger
}
