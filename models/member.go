// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
"time"
)

const TableNameMember = "member"

// Member mapped from table <member>
type Member struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Mcode          string    `gorm:"column:mcode;not null" json:"mcode"` // 확인용 Code
	AccessToken    string    `gorm:"column:access_token" json:"access_token"`
	RefreshToken   string    `gorm:"column:refresh_token" json:"refresh_token"`
	Channel        string    `gorm:"column:channel" json:"channel"`                        // 로그인 채널 : G - google plus, F - facebook, N - naver
	NickName       string    `gorm:"column:nick_name" json:"nick_name"`                    // 닉네임
	ProfileImageID int64     `gorm:"column:profile_image_id" json:"profile_image_id"`      // 프로필 이미지 id
	ActiveYn       bool      `gorm:"column:active_yn;not null;default:1" json:"active_yn"` // 활성화 : 1(active), 휴면 : 0 (non-active)
	TUUID          string    `gorm:"column:t_uuid" json:"t_uuid"`
	TGid           string    `gorm:"column:t_gid" json:"t_gid"`
	UseYn          bool      `gorm:"column:use_yn;not null;default:1" json:"use_yn"` // 사용여부
	DelYn          bool      `gorm:"column:del_yn;not null" json:"del_yn"`           // 탈퇴여부
	LastLoginDt    time.Time `gorm:"column:last_login_dt" json:"last_login_dt"`      // 최종 로그인 일시
	LoginToken     string    `gorm:"column:login_token" json:"login_token"`          // 로그인토큰
	CreatedDt      time.Time `gorm:"column:created_dt;not null" json:"created_dt"`   // 등록일
	UpdatedDt      time.Time `gorm:"column:updated_dt;not null" json:"updated_dt"`   // 수정일
	DeletedDt      time.Time `gorm:"column:deleted_dt" json:"deleted_dt"`            // 탈퇴일
}

// TableName Member's table name
func (*Member) TableName() string {
	return TableNameMember
}

