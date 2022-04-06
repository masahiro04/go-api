package factories

import (
	"go-api/domains/models"
	"go-api/domains/models/user"
	"strconv"
	"time"
)

func User() models.User {
	id, _ := user.NewId(1)
	uuid, _ := user.NewUUID("UUID")
	name, _ := user.NewName("大久保")
	email, _ := user.NewEmail("test@example.com")
	return models.BuildUser(id, uuid, name, email, time.Time{}, time.Time{})
}

func UserWithID(newID int) models.User {
	id, _ := user.NewId(newID)
	uuid, _ := user.NewUUID("UUID")
	name, _ := user.NewName("大久保")
	email, _ := user.NewEmail("test@example.com")
	return models.BuildUser(id, uuid, name, email, time.Time{}, time.Time{})
}

func Users(length int) models.Users {
	var users []models.User
	n := 1
	for n < length {
		id, _ := user.NewId(n)
		uuid, _ := user.NewUUID("UUID" + strconv.Itoa(n))
		name, _ := user.NewName("名前" + strconv.Itoa(n))
		email, _ := user.NewEmail("test" + strconv.Itoa(n) + "@example.com")
		users = append(users, models.BuildUser(id, uuid, name, email, time.Time{}, time.Time{}))
		n++
	}
	return models.NewUsers(users)
}
