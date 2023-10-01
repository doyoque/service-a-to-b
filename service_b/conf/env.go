package conf

import (
	"os"
	"time"

	util "github.com/doyoque/service_util"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	EnvLocal       = "local"
	EnvDevelopment = "dev"
	EnvStaging     = "stg"
	EnvProduction  = "prod"
)

var (
	Env     string
	AppName string
	AppPort int
	BatchDB int

	Now = func() time.Time {
		return time.Now().UTC()
	}
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	InitApp()
}

func InitApp() {
	Env = getEnv()

	AppName = util.GetOrDefault("APP_NAME", "service_b")
	AppPort = util.GetOrDefault("APP_PORT", 8200)
	BatchDB = util.GetOrDefault("DB_BATCH_SIZE", 30)
}

func getEnv() string {
	switch os.Getenv("ENVIRONMENT") {
	case EnvDevelopment:
		return EnvDevelopment
	case EnvStaging:
		return EnvStaging
	case EnvProduction:
		return EnvProduction
	default:
		return EnvLocal
	}
}
