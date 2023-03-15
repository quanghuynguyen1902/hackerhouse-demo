package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string
	Postgres    DBConnection
	Helius      Helius
}

type DBConnection struct {
	Host string
	Port string
	User string
	Name string
	Pass string

	SSLMode string
}

type Helius struct {
	API string
	Key string
}

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

// generateConfigFromViper generate config from viper data
func generateConfigFromViper(v *viper.Viper) *Config {
	tokenTTLInDay := v.GetInt("ACCESS_TOKEN_TTL")
	if tokenTTLInDay == 0 {
		tokenTTLInDay = 7
	}

	return &Config{
		ServiceName: v.GetString("SERVICE_NAME"),
		Postgres: DBConnection{
			Host:    v.GetString("DB_HOST"),
			Port:    v.GetString("DB_PORT"),
			User:    v.GetString("DB_USER"),
			Name:    v.GetString("DB_NAME"),
			Pass:    v.GetString("DB_PASS"),
			SSLMode: v.GetString("DB_SSL_MODE"),
		},
		Helius: Helius{
			API: v.GetString("HELIUS_API"),
			Key: v.GetString("HELIUS_KEY"),
		},
	}
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) *Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("ENV", "local")
	v.SetDefault("HELIUS_KEY", "199bc455-a4c7-441a-b099-45d4335229e8")
	v.SetDefault("HELIUS_API", "https://api.helius.xyz")
	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return generateConfigFromViper(v)
}

// DefaultConfigLoaders is default loader list
func DefaultConfigLoaders() []Loader {
	loaders := []Loader{}
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}
