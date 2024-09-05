package conf

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

func SetupNeo() (neo4j.DriverWithContext, error) {
	var url string = Config.NeoURI
	var username string = Config.NeoUser
	var password string = Config.NeoPassword
	return neo4j.NewDriverWithContext(url, neo4j.BasicAuth(username, password, ""))
}
