package schemas

import (
	"context"
	"errors"
	"livaf/conf"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"id"`
}

type CreateAccount struct {
	Username             string `json:"username"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (createAccount CreateAccount) Create() (*User, error) {
	session, err := conf.GetNeoSession()
	if err != nil {
		return nil, err
	}
	defer session.Close(context.Background())
	var query string = `
		CREATE (u:User {username: $username, first_name: $first_name, last_name: $last_name
		password: $password}
		RETURN u.id as id, u.username AS username, u.first_name AS first_name, u.last_name AS last_name
		)
	`
	result, err := session.ExecuteWrite(
		context.Background(),
		func(tx neo4j.ManagedTransaction) (interface{}, error) {
			record, err := tx.Run(context.Background(), query, map[string]interface{}{
				"username":   createAccount.Username,
				"first_name": createAccount.FirstName,
				"last_name":  createAccount.LastName,
				"password":   createAccount.Password,
			})
			if err != nil {
				return nil, err
			}
			if record.Next(context.Background()) {
				id, _ := record.Record().Get("id")
				username, _ := record.Record().Get("username")
				firstName, _ := record.Record().Get("first_name")
				lastName, _ := record.Record().Get("last_name")
				var user *User = &User{
					Id:        id.(int),
					Username:  username.(string),
					FirstName: firstName.(string),
					LastName:  lastName.(string),
				}
				return user, nil
			}
			return nil, errors.New("could not create user")
		},
	)
	if err != nil {
		return nil, err
	}
	return result.(*User), nil
}
