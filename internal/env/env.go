package env

import "os"

type envVars struct {
	PORT string
}

var Env envVars

func LoadEnvVars() {
	Env.PORT = os.Getenv("PORT")
}
