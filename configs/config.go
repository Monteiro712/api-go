package configs

import "github.com/spf13/viper"

//ponteiro para uma instância da estrutura config
var cfg *config
//config contém as configurações para a API e o banco de dados.
type config struct {
	API APIConfig
	DB DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Database string
}

//chamada no start da aplicação
func init() {
	viper.SetDefault("api.port", "5432")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}
//responsável por carregar as configurações do arquivo TOML.
func Load() error{
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	//cria uma instância vazia da estrutura config.
	cfg = new(config)
	//atribuir valores lidos do Viper para a estrutura config.
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		User: viper.GetString("database.user"),
		Pass: viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return nil
	
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}