package model

import (
	"time"
)

type Role struct {
	ID        int64     `json:"id" gorm:"primary_key;comment:'角色ID'"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:'创建时间'"`
	RoleName  string    `json:"role_name" gorm:"type:VARCHAR(191);unique;not null;comment:'角色名称'"`
	Polices   []*Policy `json:"polices" gorm:"many2many:role_polices"`
}

type Roles struct {
	List  []*Role `json:"list"`
	Pager *Pager  `json:"pager"`
}

type RoleQueryParam struct {
	Order string
}

type Policy struct {
	ID     int64  `json:"id" gorm:"primary_key;comment:'规则ID'"`
	Router string `json:"router" gorm:"type:VARCHAR(191);not null;comment:'请求路由'"`
	Method string `json:"method" gorm:"type:VARCHAR(30);not null;comment:'请求方式'"`
}

type RolePolices struct {
	RoleID   int64 `json:"role_id" gorm:"comment:'角色ID'"`
	PolicyID int64 `json:"policy_id" gorm:"comment:'规则ID'"`
}
