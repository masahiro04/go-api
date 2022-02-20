package testData

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/domains/user"
	"strconv"
	"time"
)

func User() domains.User {
	id, _ := user.NewId(1)
	name, _ := user.NewName("大久保")
	email, _ := user.NewEmail("test@example.com")
	return domains.BuildUser(id, name, email, time.Time{}, time.Time{})
}

func UserWithID(newID int) domains.User {
	id, _ := user.NewId(newID)
	name, _ := user.NewName("大久保")
	email, _ := user.NewEmail("test@example.com")
	return domains.BuildUser(id, name, email, time.Time{}, time.Time{})
}

func Users(length int) domains.Users {
	var users []domains.User
	n := 1
	for n < length {
		id, _ := user.NewId(n)
		name, _ := user.NewName("名前" + strconv.Itoa(n))
		email, _ := user.NewEmail("test" + strconv.Itoa(n) + "@example.com")
		users = append(users, domains.BuildUser(id, name, email, time.Time{}, time.Time{}))
		n++
	}
	return domains.NewUsers(users)
}
