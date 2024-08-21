package config

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func LoadEnv(){
	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error Load Env %v", err)
	}
}

func GetEnv(key string) string {
	var val string = os.Getenv(key)
	if (val == "") {
		log.Fatalf("Env Key not found")
	} 
	return val
}

func GothSetup(){
	LoadEnv()
	sessionStore := sessions.NewCookieStore([]byte(GetEnv("SECRET_KEY")))
	sessionStore.Options = &sessions.Options{
		Path:     "/",               // Path untuk cookie
		Domain:   "",                // Domain (kosong untuk default)
		MaxAge:   3600,             // Durasi cookie dalam detik
		HttpOnly: true,             // Hanya bisa diakses melalui HTTP
		Secure:   false,            // Gunakan HTTPS
	}
	goth.UseProviders(
		google.New(GetEnv("GOOGLE_CLIENT"), GetEnv("GOOGLE_SECRET"), GetEnv("WEB_URL")+GetEnv("WEB_PORT")+"/api/login/google/callback", "email", "profile"),
	)
	gothic.Store = sessionStore
	
}