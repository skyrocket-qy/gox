package configx

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox"
)

func NewConfig[config any](curEnv string, conf *config) (err error) {
	if curEnv == gox.EnvLocal {
		err := godotenv.Load(".env")
		// In a local environment, it's okay for the .env file to be missing.
		// We should only return an error if one occurred that is NOT a "file does not exist" error.
		if err != nil && !os.IsNotExist(err) {
			return erx.W(err)
		}
	}

	if err := env.Parse(conf); err != nil {
		return erx.W(err)
	}

	return nil
}
