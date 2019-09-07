package model

import (
	uuid "github.com/satori/go.uuid"
)

type menu struct {
	ID       uuid.UUID // 菜单ID
	ParensID uuid.UUID // 菜单父节点ID
	Name     string    // 菜单名称
	Route    string    // 菜单路由
	Code     int64     // 菜单标识
}
