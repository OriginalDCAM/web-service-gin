package main

import (
	"os"

	engine "github.com/OrignalDCAM/web-service-gin/router"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)




func main() {
	godotenv.Load()

	engine.Run(os.Getenv("WEB_API_PORT"))
}
