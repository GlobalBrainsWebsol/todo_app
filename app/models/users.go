package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func UserLists() (users []User, err error) {
	cmd := `select id, uuid, name,email, created_at from users`
	rows, err := Db.Query(cmd)
	users = []User{}
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}
func (u *User) CreateUser() (err error) {
	cmd := `insert into users(
		uuid,
		name,
		email,
		password,
		created_at
	) values(
		:1,:2,:3,:4,:5
	)`
	_, err = Db.Exec(cmd, CreateUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func GetUser(id int) (user User, err error) {
	cmd := `select id, uuid, name, email,created_at from users where id = :1`
	user = User{}
	err = Db.QueryRow(cmd, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	fmt.Println("update")
	cmd := `update users set name = :1, email = :2 where id = :3`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		return err
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = :1`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
