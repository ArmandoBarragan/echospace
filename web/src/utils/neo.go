package utils

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

func GetRecordKeys(record neo4j.ResultWithContext) map[string]interface{} {
	return make(map[string]interface{})
}
