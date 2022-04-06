package models

type Users struct {
	// Userの集合を表現
	Value []User
}

func NewUsers(value []User) Users {
	return Users{Value: value}
}

func (users Users) Size() int {
	return len(users.Value)
}

func EmptyUsers() Users {
	return Users{Value: []User{}}
}
func (users *Users) ApplyLimitAndOffset(limit, offset int) []User {
	if limit <= 0 {
		return []User{}
	}

	usersSize := users.Size()
	min := offset
	if min < 0 {
		min = 0
	}

	if min > usersSize {
		return []User{}
	}

	max := min + limit
	if max > usersSize {
		max = usersSize
	}

	return users.Value[min:max]
}
