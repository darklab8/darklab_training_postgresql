package settings

import "os"

var DatabaseHost string

func init() {
	var dbhost_exists bool
	DatabaseHost, dbhost_exists = os.LookupEnv("DATABASE_HOST")
	if !dbhost_exists {
		DatabaseHost = "localhost"
	}
}
