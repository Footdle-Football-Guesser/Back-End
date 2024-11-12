package main

type UserDAO struct {
	users []*User
}

func NewUserDAO() *UserDAO {
	userdao := new(UserDAO)
	userdao.users = make([]*User, 0)
	return userdao
}

func (userdao *UserDAO) GetAll() []*User {
	return userdao.users
}

func (userdao *UserDAO) Insert(user *User) {
	userdao.users = append(userdao.users, user)
}

func (userdao *UserDAO) Get(id int) *User {
	for _, user := range userdao.users {
		if user.Id == id {
			return user
		}
	}
	return nil
}

func (userdao *UserDAO) Update(user *User) string {
	for i, existingUser := range userdao.users {
		if existingUser.Id == user.Id {
			userdao.users[i] = user
			return "ok"
		}
	}
	return "user not found"
}

func (userdao *UserDAO) Delete(id int) string {
	for i, user := range userdao.users {
		if user.Id == id {
			userdao.users = append(userdao.users[:i], userdao.users[i+1:]...)
			return "ok"
		}
	}
	return "user not found"
}
