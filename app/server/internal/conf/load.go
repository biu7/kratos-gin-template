package conf

import (
	"fmt"
	"kratos-gin-template/app/server/internal/constants"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func Load(path string) (*Bootstrap, error) {
	if path == "" {
		path = os.Getenv(constants.ConfEnv)
	}
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)

	defer c.Close()

	if err := c.Load(); err != nil {
		panic(fmt.Errorf("could not load config file on '%s': %w", path, err))
	}

	var bc Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return &bc, nil
}
