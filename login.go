package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("jouwSecretKeyHier") // Gebruik een veilige, unieke sleutel voor productie!

func generateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kan verzoek niet verwerken"})
	}

	// Query de database voor de gebruiker
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", request.Username).Scan(&hashedPassword)
	if err != nil {
		// Gebruiker niet gevonden of andere fout
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Onjuiste gebruikersnaam of wachtwoord"})
	}

	// Vergelijk het opgegeven wachtwoord met het gehashte wachtwoord in de database
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(request.Password))
	if err != nil {
		// Fout betekent dat het wachtwoord niet overeenkomt
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Onjuiste gebruikersnaam of wachtwoord"})
	}

	// Wachtwoord is correct, genereer een token
	token, err := generateJWT(request.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kan JWT niet genereren"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func authenticate(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("onverwachte ondertekeningsmethode: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Ongeldig token"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("username", claims["name"])
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Ongeldig token"})
	}
}
