package models

type UserRepo struct {
	conn ConnectionInterface
}

func NewUserRepo() UserRepo {
	return UserRepo{}
}

func (repo UserRepo) Find(id int) (User, error) {
	user := &User{}
	err := conn().SelectOne(&user, "select id, name from user where id = ?", id)
	if err != nil {
		return *user, err
	}
	return *user, nil
}

func (repo UserRepo) Insert(user *User) (error) {
	err := conn().Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepo) Update(user *User) (error) {
	_, err := conn().Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepo) Delete(id int) (User, error) {
	user := &User{Id: id}
	_, err := conn().Delete(user)
	if err != nil {
		return *user, err
	}
	return *user, nil
}
