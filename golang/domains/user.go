package domains

import (
	"clean_architecture/golang/domains/user"
	"time"
)

// TODO: privateに扱うために、小文字に変更する
type User struct {
	ID        user.ID
	Name      user.Name
	Email     user.Email
	Password  user.Password
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewUser(name user.Name, email user.Email, password user.Password) User {
	return User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

// repositoryやfactory経由の生成において使用する関数
// 生成時のバリデーションをしないことに注意
func BuildUser(id user.ID, name user.Name, email user.Email, createdAt time.Time, updatedAt time.Time) User {
	return User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (u *User) UpdateName(name user.Name) *User {
	u.Name = name
	return u
}

func (u *User) UpdateEmail(email user.Email) *User {
	u.Email = email
	return u
}

func (u *User) UpdatePassword(password user.Password) *User {
	u.Password = password
	return u
}
