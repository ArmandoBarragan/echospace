package schemas

import (
	"context"
	"errors"
	"livaf/conf"
	"livaf/src/utils"
	"regexp"

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

func (createAccount CreateAccount) PasswordIsValid() bool {
	var passwordsMatch bool = createAccount.Password == createAccount.PasswordConfirmation
	var re *regexp.Regexp = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]1234567890`)
	var passwordContainsSpecialCharacter bool = re.MatchString(createAccount.Password)
	return passwordsMatch && passwordContainsSpecialCharacter
}

func (createAccount CreateAccount) Create() (*User, error) {
	session, err := conf.GetNeoSession()
	if err != nil {
		return nil, err
	}
	defer session.Close(context.Background())
	var query string = `
		CREATE (u:User {username: $username, first_name: $first_name, last_name: $last_name,
		password: $password})
		RETURN ID(u) AS id, u.username AS username, u.first_name AS first_name, u.last_name AS last_name
	`
	password, err := utils.HashPassword(createAccount.Password)
	if err != nil {
		return nil, err
	}
	result, err := session.ExecuteWrite(
		context.Background(),
		func(tx neo4j.ManagedTransaction) (interface{}, error) {
			record, err := tx.Run(context.Background(), query, map[string]interface{}{
				"username":   createAccount.Username,
				"first_name": createAccount.FirstName,
				"last_name":  createAccount.LastName,
				"password":   password,
			})
			if err != nil {
				return nil, err
			}
			if record.Next(context.Background()) {
				id, _ := record.Record().Get("id")
				idInt64 := id.(int64)
				username, _ := record.Record().Get("username")
				firstName, _ := record.Record().Get("first_name")
				lastName, _ := record.Record().Get("last_name")
				var user *User = &User{
					Id:        int(idInt64),
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
