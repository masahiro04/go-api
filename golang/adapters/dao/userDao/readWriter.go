package userDao

import (
	"clean_architecture/golang/domains"
	userModel "clean_architecture/golang/domains/user"
	"time"

	"database/sql"
	"log"
)

type rw struct {
	store *sql.DB
}

func New(db *sql.DB) *rw {
	return &rw{
		store: db,
	}
}

type UserDto struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rw rw) GetAll() (*domains.Users, error) {
	var users []domains.User
	rows, err := rw.store.Query(GetAllSql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userDto UserDto

		if err = rows.Scan(
			&userDto.ID,
			&userDto.Name,
			&userDto.Email,
			&userDto.CreatedAt,
			&userDto.UpdatedAt,
		); err != nil {
			log.Fatal(err)
			return nil, err
		}

		id, _ := userModel.NewId(userDto.ID)
		name, _ := userModel.NewName(userDto.Name)
		email, _ := userModel.NewEmail(userDto.Email)
		newUser := domains.BuildUser(id, name, email, userDto.CreatedAt, userDto.UpdatedAt)

		users = append(users, newUser)
	}
	usersData := domains.NewUsers(users)
	return &usersData, nil
}

func (rw rw) GetById(id int) (*domains.User, error) {
	var userDto UserDto

	result := rw.store.QueryRow(GetByIdSql, id)
	err := result.Scan(
		&userDto.ID,
		&userDto.Name,
		&userDto.Email,
		&userDto.CreatedAt,
		&userDto.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := userModel.NewId(userDto.ID)
	name, _ := userModel.NewName(userDto.Name)
	email, _ := userModel.NewEmail(userDto.Email)
	newUser := domains.BuildUser(_id, name, email, userDto.CreatedAt, userDto.UpdatedAt)
	return &newUser, nil
}

func (rw rw) Create(newUser domains.User) (*domains.User, error) {
	var id int
	err := rw.store.QueryRow(
		CreateSql,
		newUser.Name.Value, newUser.Email.Value, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := userModel.NewId(id)
	name, _ := userModel.NewName(newUser.Name.Value)
	email, _ := userModel.NewEmail(newUser.Email.Value)
	user := domains.BuildUser(_id, name, email, newUser.CreatedAt, newUser.UpdatedAt)
	return &user, nil
}

func (rw rw) CreateTx(newUser domains.User, tx *sql.Tx) (*domains.User, error) {
	var id int
	err := tx.QueryRow(
		CreateSql,
		newUser.Name, newUser.Email, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := userModel.NewId(newUser.ID.Value)
	name, _ := userModel.NewName(newUser.Name.Value)
	email, _ := userModel.NewEmail(newUser.Email.Value)
	user := domains.BuildUser(_id, name, email, newUser.CreatedAt, newUser.UpdatedAt)
	return &user, nil
}

//
func (rw rw) Update(id int, user domains.User) (*domains.User, error) {
	_, err := rw.store.Exec(
		UpdateSql,
		id, user.Name.Value, user.Email.Value, time.Now())

	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := userModel.NewId(id)
	name, _ := userModel.NewName(user.Name.Value)
	email, _ := userModel.NewEmail(user.Email.Value)
	newUser := domains.BuildUser(_id, name, email, user.CreatedAt, user.UpdatedAt)
	return &newUser, nil
}

func (rw rw) Delete(id int) error {
	if _, err := rw.store.Exec(DeleteSql, id, time.Now(), time.Now()); err != nil {
		return err
	}
	return nil
}
