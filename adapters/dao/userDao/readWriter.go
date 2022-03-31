package userDao

import (
	"errors"
	"go-api/domains"
	userModel "go-api/domains/user"

	"gorm.io/gorm"
)

type rw struct {
	db *gorm.DB
}

func New(db *gorm.DB) *rw {
	return &rw{
		db: db,
	}
}

type UserDto struct {
	gorm.Model
	ID    int    `json:"id"`
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (UserDto) TableName() string {
	return "users"
}

func (rw rw) GetAll() (*domains.Users, error) {
	var dtos []UserDto
	rw.db.Find(&dtos)
	var users []domains.User

	for _, dto := range dtos {

		id, _ := userModel.NewId(dto.ID)
		uuid, _ := userModel.NewUUID(dto.UUID)
		name, _ := userModel.NewName(dto.Name)
		email, _ := userModel.NewEmail(dto.Email)
		newUser := domains.BuildUser(id, uuid, name, email, dto.CreatedAt, dto.UpdatedAt)

		users = append(users, newUser)
	}

	usersData := domains.NewUsers(users)
	return &usersData, nil
}

func (rw rw) GetById(id int) (*domains.User, error) {
	var dto UserDto

	err := rw.db.First(&dto, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	_id, _ := userModel.NewId(dto.ID)
	uuid, _ := userModel.NewUUID(dto.UUID)
	name, _ := userModel.NewName(dto.Name)
	email, _ := userModel.NewEmail(dto.Email)
	newUser := domains.BuildUser(_id, uuid, name, email, dto.CreatedAt, dto.UpdatedAt)
	return &newUser, nil
}

func (rw rw) Create(newUser domains.User) (*domains.User, error) {
	dto := UserDto{
		UUID:  newUser.UUID.Value,
		Name:  newUser.Name.Value,
		Email: newUser.Email.Value,
	}

	rw.db.Create(&dto)

	_id, _ := userModel.NewId(dto.ID)
	uuid, _ := userModel.NewUUID(dto.UUID)
	name, _ := userModel.NewName(newUser.Name.Value)
	email, _ := userModel.NewEmail(newUser.Email.Value)
	user := domains.BuildUser(_id, uuid, name, email, newUser.CreatedAt, newUser.UpdatedAt)
	return &user, nil
}

// NOTE(okubo): transactioの場合に利用
func (rw rw) CreateTx(newUser domains.User, tx *gorm.DB) (*domains.User, error) {
	dto := UserDto{
		UUID:  newUser.UUID.Value,
		Name:  newUser.Name.Value,
		Email: newUser.Email.Value,
	}

	tx.Create(&dto)

	_id, _ := userModel.NewId(dto.ID)
	uuid, _ := userModel.NewUUID(dto.UUID)
	name, _ := userModel.NewName(newUser.Name.Value)
	email, _ := userModel.NewEmail(newUser.Email.Value)
	user := domains.BuildUser(_id, uuid, name, email, newUser.CreatedAt, newUser.UpdatedAt)

	return &user, nil
}

//
func (rw rw) Update(id int, user domains.User) (*domains.User, error) {
	dto := UserDto{}

	rw.db.First(&dto, id).Updates(UserDto{
		ID:    id,
		Name:  user.Name.Value,
		Email: user.Email.Value,
	})
	_id, _ := userModel.NewId(id)
	uuid, _ := userModel.NewUUID(dto.UUID)
	name, _ := userModel.NewName(user.Name.Value)
	email, _ := userModel.NewEmail(user.Email.Value)
	newUser := domains.BuildUser(_id, uuid, name, email, user.CreatedAt, user.UpdatedAt)
	return &newUser, nil
}

func (rw rw) Delete(id int) error {
	dto := UserDto{}
	rw.db.First(&dto, id).Delete(&dto)
	return nil
}
