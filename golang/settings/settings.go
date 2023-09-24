package settings

import "os"

var DatabaseHost string
var ENABLED_PERFORMANCE_TESTS bool

func init() {
	var dbhost_exists bool
	DatabaseHost, dbhost_exists = os.LookupEnv("DATABASE_HOST")
	if !dbhost_exists {
		DatabaseHost = "localhost"
	}

	_, ENABLED_PERFORMANCE_TESTS = os.LookupEnv("ENABLE_PERFORMANCE_TESTS")
}
