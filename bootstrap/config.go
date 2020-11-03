package bootstrap

import (
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
func InitConfig() (Config, error) {
	godotenv.Load("../.env")

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

	return c, nil
}
