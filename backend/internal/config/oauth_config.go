package config

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"os"
)

func InitializeOAuth() {
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			os.Getenv("GOOGLE_CALLBACK_URL"),
		),
	)
}
