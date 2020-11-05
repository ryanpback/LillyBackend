package bootstrap

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config holds the structure of the config
type Config struct {
	AppPort           string
	GCSBucket         string
	FormFileName      string
	MultiPartMemLimit int64
	Logger            *logrus.Logger
}

// InitConfig sets the app config
func InitConfig(envFile string) (Config, error) {
	godotenv.Load("../" + envFile)

	multiPartMemLimit, err := strconv.ParseInt(os.Getenv("MULTI_PART_MEM_LIMIT"), 10, 64)
	if err != nil {
		panic(err)
	}

	c := Config{
		AppPort:           os.Getenv("APP_PORT"),
		GCSBucket:         os.Getenv("GCS_BUCKET_NAME"),
		FormFileName:      os.Getenv("FORM_FILE_NAME"),
		MultiPartMemLimit: multiPartMemLimit,
	}

	// initialize logging to stdout
	c.Logger = logrus.New()
	c.Logger.Out = os.Stdout

	if c.AppPort == "" {
		return c, fmt.Errorf("APP_PORT is required")
	}
	if c.GCSBucket == "" {
		return c, fmt.Errorf("GCS_BUCKET_NAME is required")
	}
	if c.FormFileName == "" {
		return c, fmt.Errorf("FORM_FILE_NAME is required")
	}

	return c, nil
}
