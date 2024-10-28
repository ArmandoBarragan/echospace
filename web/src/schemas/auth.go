package schemas

import (
	"context"
	"echospace/conf"
	"echospace/src/utils"
	"errors"
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

func (createAccount CreateAccount) UsernameExists() (bool, error) {
	// Check if the username exists
	session, err := conf.GetNeoSession()
	if err != nil {
		return false, err
	}
	defer session.Close(context.Background())
	var ctx context.Context = context.Background()
	var query string = `
	MATCH (u:User {username: $username}) 
	RETURN CASE WHEN u IS NOT NULL THEN true ELSE false END AS usernameExists
	`
	usernameExists, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(
				ctx,
				query,
				map[string]interface{}{
					"username": createAccount.Username,
				},
			)
			if err != nil {
				return nil, err
			}
			if result.Next(ctx) {
				usernameExistsResult, ok := result.Record().Get("usernameExists")
				if !ok {
					return false, errors.New("unexpected value")
				}
				return usernameExistsResult, nil
			}
			return nil, errors.New("could not execute query")
		},
	)
	if err != nil {
		return false, err
	}
	return usernameExists.(bool), nil
}

func (createAccount CreateAccount) PasswordIsValid() bool {
	var re *regexp.Regexp = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	var reDigit *regexp.Regexp = regexp.MustCompile("[0-9]")
	var passwordContainsSpecialCharacter bool = re.MatchString(createAccount.Password)
	var passwordContainsDigit bool = reDigit.MatchString(createAccount.Password)
	return passwordContainsSpecialCharacter && passwordContainsDigit
}

func (createAccount CreateAccount) Create() (*User, error) {
	// Create the user account
	session, err := conf.GetNeoSession()
	if err != nil {
		return nil, err
	}
	defer session.Close(context.Background())
	var ctx context.Context = context.Background()
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
			record, err := tx.Run(ctx, query, map[string]interface{}{
				"username":   createAccount.Username,
				"first_name": createAccount.FirstName,
				"last_name":  createAccount.LastName,
				"password":   password,
			})
			if err != nil {
				return nil, err
			}
			if record.Next(ctx) {
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
