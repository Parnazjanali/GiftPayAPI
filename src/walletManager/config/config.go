package config

import (
	"log"
	"os"
	"walletManager/internal/models"

	"github.com/joho/godotenv"
)

var AppConfig *models.Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default config")
	}

	AppConfig = &models.Config{
		Port:       GetEnv("PORT", "8081"),
		JwtSeceret: GetEnv("JWT_SECRET", "3c52a999f2f1fa20847f7b3661613f2e9a6c753187a84d057942a140ed493c3a7855c919732029824694661111f0e30fffb21b2661bd3d7cb0c4d9fb3a3060a61a05363388f77f0207bfded5f192e158d05319938ec03764fc61f6c0df2f117d4076a84f3fb7340c742181dff17df91375699eb9a1331783334ad812245a75ff6505a64028f2a35b0d49677a5043bb0d7702a9e7d4b51c69e7d8063bc25017123c563e5f75e88dd1c7fb648bec7ec516ff59142f462858d00f41dcc9009feeb9cdf07b281bffe70d9f243bd2057832d87cf9cf6a7581479b324b5cc4f99ac14a8854b52f30a87fe93bcbbeb0b59a222484de50b3fdde55d137126c48bc49eeda"),
	}

}

func GetEnv(key, Fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return Fallback
}
