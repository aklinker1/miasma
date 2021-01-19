package env

import "os"

var IS_PROD = os.Getenv("MODE") == "production"
