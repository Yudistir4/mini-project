package util

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	return viper.GetString(key)
}

func FileHandling(c echo.Context) (string, error) {

	file, err := c.FormFile("file")
	if err != nil {
		return "", nil
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	name := uuid.New().String() + filepath.Ext(file.Filename)

	// Destination
	dst, err := os.Create("images/" + name)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return name, nil
}
