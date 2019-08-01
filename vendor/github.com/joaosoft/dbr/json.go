package dbr

import "github.com/joaosoft/json"

func Marshal(object interface{}) ([]byte, error) {
	return json.Marshal(object, "db.write", "db")
}

func Unmarshal(bytes []byte, object interface{}) error {
	return json.Unmarshal(bytes, object, "db.read", "db")
}
