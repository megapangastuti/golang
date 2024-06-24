// ToDo :
// 1. Mendeklarasikan nama package pada file config.go
// 2. Mendeklarasikan struct bernama DBConfig dan APIConfig
// 3. Membuat sebuah fungsi baru yang bernama readConfig
// 4. Membuat sebuah fungsi baru bernama NewConfig

package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type DBConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Driver   string
}

type APIConfig struct {
	ApiPort string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	DBConfig
	APIConfig
	TokenConfig
}

func (c *Config) readConfig() error {
	c.DBConfig = DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Database: "book_db",
		Username: "postgres",
		Password: "123234",
		Driver:   "postgres",
	}

	c.APIConfig = APIConfig{
		ApiPort: "8080",
	}

	accessTokenLifeTime := time.Duration(1) * time.Hour

	c.TokenConfig = TokenConfig{
		ApplicationName:     "Enigma Camp",
		JwtSignatureKey:     []byte("IniSangatRahasia!!!!"),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	if c.Host == "" || c.Port == "" || c.Username == "" || c.Password == "" || c.ApiPort == "" {
		return fmt.Errorf("required config")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cfg.readConfig()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
