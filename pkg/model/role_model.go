package model

type RoleID int64
type RoleName string

type Role struct {
	ID        RoleID   `json:"id" db:"id"`
	Name      RoleName `json:"name" db:"name"`
	CreatedAt string   `json:"createdAt" db:"created_at"`
	UpdatedAt string   `json:"updatedAt" db:"updated_at"`
}