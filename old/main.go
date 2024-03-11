package main

import (
	"log"
	"os"
	"time"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/providers"
	"github.com/ESMO-ENTERPRISE/auth-server/routes"
	"github.com/ESMO-ENTERPRISE/auth-server/services"
	"github.com/ESMO-ENTERPRISE/auth-server/token"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
)

var con database.Connector

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load database
	con.InitDatabase()
	// con.MigrateDatabase()

	token.InitTokenService()
	providers.InitGithubFlow()

	// Http server
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowOrigins:     os.Getenv("FRONTEND_URL"),
		AllowCredentials: true,
		AllowMethods:     "GET, POST, DELETE, PATCH",
	}))
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.SlidingWindow{},
		// Storage:                conn.Ratelimter,
	}))

	// Setup routes
	authService := services.AuthService{
		Conn: &con,
	}
	clientService := services.ClientService{
		Conn: &con,
	}

	routes.AuthRoutes(&authService, app)
	routes.GithubRoutes(app)

	// JWT Middleware, from here all routes below must have an authentication token
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    token.PrivateKey.Public(),
		},
	}))

	routes.ClientRoutes(&clientService, app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
