package userDao

import (
	"errors"
	"go-api/domains/models"
	"go-api/domains/models/user"
	"go-api/usecases"

	"gorm.io/gorm"
)

type rw struct {
	db *gorm.DB
}

// NOTE(okubo): interfaceで抽象実装
func New(db *gorm.DB) usecases.UserRepository {
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

func (rw rw) GetAll() (*models.Users, error) {
	var dtos []UserDto
	rw.db.Find(&dtos)
	var users []models.User

	for _, dto := range dtos {
		id, _ := user.NewId(dto.ID)
		uuid, _ := user.NewUUID(dto.UUID)
		name, _ := user.NewName(dto.Name)
		email, _ := user.NewEmail(dto.Email)
		newUser := models.BuildUser(id, uuid, name, email, dto.CreatedAt, dto.UpdatedAt)

		users = append(users, newUser)
	}

	usersData := models.NewUsers(users)
	return &usersData, nil
}

func (rw rw) GetById(id int) (*models.User, error) {
	var dto UserDto

	err := rw.db.First(&dto, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	_id, _ := user.NewId(dto.ID)
	uuid, _ := user.NewUUID(dto.UUID)
	name, _ := user.NewName(dto.Name)
	email, _ := user.NewEmail(dto.Email)
	newUser := models.BuildUser(_id, uuid, name, email, dto.CreatedAt, dto.UpdatedAt)
	return &newUser, nil
}

func (rw rw) Create(newUser models.User) (*models.User, error) {
	dto := UserDto{
		UUID:  newUser.UUID.Value,
		Name:  newUser.Name.Value,
		Email: newUser.Email.Value,
	}

	rw.db.Create(&dto)

	_id, _ := user.NewId(dto.ID)
	uuid, _ := user.NewUUID(dto.UUID)
	name, _ := user.NewName(newUser.Name.Value)
	email, _ := user.NewEmail(newUser.Email.Value)
	user := models.BuildUser(_id, uuid, name, email, newUser.CreatedAt, newUser.UpdatedAt)
	return &user, nil
}

// NOTE(okubo): transactioの場合に利用
func (rw rw) CreateTx(newUser models.User, tx *gorm.DB) (*models.User, error) {
	dto := UserDto{
		UUID:  newUser.UUID.Value,
		Name:  newUser.Name.Value,
		Email: newUser.Email.Value,
	}

	tx.Create(&dto)

	_id, _ := user.NewId(dto.ID)
	uuid, _ := user.NewUUID(dto.UUID)
	name, _ := user.NewName(newUser.Name.Value)
	email, _ := user.NewEmail(newUser.Email.Value)
	user := models.BuildUser(_id, uuid, name, email, newUser.CreatedAt, newUser.UpdatedAt)

	return &user, nil
}

//
func (rw rw) Update(id int, newUser models.User) (*models.User, error) {
	dto := UserDto{}

	rw.db.First(&dto, id).Updates(UserDto{
		ID:    id,
		Name:  newUser.Name.Value,
		Email: newUser.Email.Value,
	})
	_id, _ := user.NewId(id)
	uuid, _ := user.NewUUID(dto.UUID)
	name, _ := user.NewName(newUser.Name.Value)
	email, _ := user.NewEmail(newUser.Email.Value)
	updatedUser := models.BuildUser(_id, uuid, name, email, newUser.CreatedAt, newUser.UpdatedAt)
	return &updatedUser, nil
}

func (rw rw) Delete(id int) error {
	dto := UserDto{}
	rw.db.First(&dto, id).Delete(&dto)
	return nil
}
