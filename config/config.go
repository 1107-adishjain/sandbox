package config



import(
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	Port string `env:"PORT"`
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	return &Config{
		Port: getEnv("PORT", "8080"),
	}, nil
}