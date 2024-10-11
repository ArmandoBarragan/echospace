package conf

import (
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var neoEngine neo4j.DriverWithContext

func SetupNeo() {
	var url string = Config.NeoURI
	var username string = Config.NeoUser
	var password string = Config.NeoPassword
	engine, err := neo4j.NewDriverWithContext(url, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic("Something went wrong while configuring Neo4j")
	}
	neoEngine = engine
}

func ClosePools() {
	if neoEngine != nil {
		neoEngine.Close(context.Background())
	}
}

func GetNeoSession() (neo4j.SessionWithContext, error) {
	session := neoEngine.NewSession(context.Background(), neo4j.SessionConfig{})
	if session == nil {
		return nil, errors.New("connection to database failed")
	}
	return session, nil
}
