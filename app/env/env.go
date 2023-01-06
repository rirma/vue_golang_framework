package env

import (
	"os"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	
	if err != nil {
		panic(".envが読み込めませんでした")
	}
}

func Get(text string) string {
	return os.Getenv(text)
}
