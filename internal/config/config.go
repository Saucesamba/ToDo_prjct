// configuration

package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Config Структура для хранения данных о БД и сервере
type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	Server struct {
		Port int
	}
}

//функция для загрузки конфигурации
// Создали флаги для того чтоб можно было прописать и поменять параметры БД и сервака

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.StringVar(&cfg.Database.Host, "dbhost", os.Getenv("DB_HOST"), "Database host")

	flag.IntVar(&cfg.Database.Port, "dbport", func() int {
		portSt := os.Getenv("DB_PORT")
		port, _ := strconv.Atoi(portSt)
		return port
	}(), "Database port")

	flag.StringVar(&cfg.Database.User, "dbuser", os.Getenv("DB_USER"), "Database user")
	flag.StringVar(&cfg.Database.Password, "dbpassword", os.Getenv("DB_PASSWORD"), "Database password")
	flag.StringVar(&cfg.Database.Name, "dbname", os.Getenv("DB_NAME"), "Database name")
	flag.IntVar(&cfg.Server.Port, "serverport", func() int {
		portStr := os.Getenv("SERVER_PORT")
		port, _ := strconv.Atoi(portStr)
		return port
	}(), "Server port")

	flag.Parse() //Распарсили флаги

	//Обрабатываем параметры из среды окружения и если есть - переопределяем их в конфиге

	if envHost := os.Getenv("DB_HOST"); envHost != "" {
		cfg.Database.Host = envHost
	}
	if envPort := os.Getenv("DB_PORT"); envPort != "" {
		port, err := strconv.Atoi(envPort)
		if err != nil {
			cfg.Database.Port = port
		}
	}
	if envUser := os.Getenv("DB_USER"); envUser != "" {
		cfg.Database.User = envUser
	}
	if envPassword := os.Getenv("DB_PASSWORD"); envPassword != "" {
		cfg.Database.Password = envPassword
	}
	if envName := os.Getenv("DB_NAME"); envName != "" {
		cfg.Database.Name = envName
	}

	if envServerPort := os.Getenv("SERVER_PORT"); envServerPort != "" {
		port, err := strconv.Atoi(envServerPort)
		if err == nil {
			cfg.Server.Port = port
		}
	}
	return cfg, nil

}
