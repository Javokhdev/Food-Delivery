package config

import (
  "fmt"
  "os"

  "github.com/joho/godotenv"
  "github.com/spf13/cast"
)

type Config struct {
  HTTPPort string

  PostgresHost     string
  PostgresPort     int
  PostgresUser     string
  PostgresPassword string
  PostgresDatabase string

  DefaultOffset string
  DefaultLimit  string

  TokenKey string
}


func Load() Config {
  if err := godotenv.Load(); err != nil {
    fmt.Println("No .env file found")
  }

  config := Config{}

  config.HTTPPort = cast.ToString(GetOrReturnDefaultValue("HTTP_PORT", ":8070"))

  config.PostgresHost = cast.ToString(GetOrReturnDefaultValue("POSTGRES_HOST", "postgres_dock"))
  config.PostgresPort = cast.ToInt(GetOrReturnDefaultValue("POSTGRES_PORT", 5432))
  config.PostgresUser = cast.ToString(GetOrReturnDefaultValue("POSTGRES_USER", "postgres"))
  config.PostgresPassword = cast.ToString(GetOrReturnDefaultValue("POSTGRES_PASSWORD", "root"))
  config.PostgresDatabase = cast.ToString(GetOrReturnDefaultValue("POSTGRES_DATABASE", "auth_db"))

  config.DefaultOffset = cast.ToString(GetOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
  config.DefaultLimit = cast.ToString(GetOrReturnDefaultValue("DEFAULT_LIMIT", "10"))
  config.TokenKey=cast.ToString(GetOrReturnDefaultValue("TokenKey", "my_secret_key"))
  return config
}

func GetOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
  val, exists := os.LookupEnv(key)

  if exists {
    return val
  }

  return defaultValue
}
