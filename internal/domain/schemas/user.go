package schemas

type RoleUser string

const(
  Admin RoleUser = "ADMIN"
  Coordinator RoleUser = "COORDENADOR"
  Reviewer RoleUser = "REVISOR"
)

type User struct {
  Role RoleUser `gorm:"type:role_user;not null;default:'REVISOR'"`
}
