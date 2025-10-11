/*
Environment configuration (env config) in Go refers to the practice of loading application settings or secrets (like database credentials, API keys, ports, etc.) from environment variables, instead of hardcoding them in your Go code.
This is a best practice for building secure, portable, and configurable applications — especially when deploying to servers or cloud environments.
⚙️ Why Use Env Config?
✅ Security – you don’t hardcode sensitive data (like passwords).
✅ Flexibility – you can easily change configurations per environment (dev, test, prod).
✅ Portability – makes your Go app easier to deploy on Docker, Kubernetes, etc.
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Set environment variables (for demo purposes)
	os.Setenv("APP_ENV", "development")
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PASS", "secret123")
	os.Setenv("PORT", "8080")

	// Read environment variables
	appEnv := os.Getenv("APP_ENV")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	port := os.Getenv("PORT")

	// Check if any variable is missing
	if appEnv == "" || dbUser == "" || dbPass == "" || port == "" {
		log.Fatal("Missing one or more environment variables")
	}

	fmt.Println("Environment:", appEnv)
	fmt.Println("Database User:", dbUser)
	fmt.Println("Database Password:", dbPass)
	fmt.Println("Server Port:", port)

	// Remove some environment variables
	//sample -> os.Unsetenv("DB_PASS")
}
