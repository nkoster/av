package main

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func uploadImages(c *fiber.Ctx) error {
	// Map waar de geüploade afbeeldingen worden opgeslagen
	uploadDir := "./ui/images/p"

	// Parse de multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kan upload niet verwerken"})
	}

	files := form.File["imagesUpload[]"]
	var uploadedFiles []string

	for _, file := range files {
		// Gebruik de originele bestandsnaam
		filePath := filepath.Join(uploadDir, filepath.Base(file.Filename))

		// Sla het bestand op in de aangegeven map
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kan bestand niet opslaan"})
		}

		// Voeg de bestandsnaam toe aan de lijst van geüploade bestanden
		uploadedFiles = append(uploadedFiles, file.Filename)

		// Maak een thumbnail naam en voeg die ook toe aan de lijst
		thumbnailName := "thumbnail_" + file.Filename
		// (genereer thumbnails hier)
		uploadedFiles = append(uploadedFiles, thumbnailName)

	}

	// Retourneer de lijst van bestandsnamen gescheiden door komma's
	return c.JSON(fiber.Map{"images": uploadedFiles})
}


