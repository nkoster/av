package main

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func uploadImages(c *fiber.Ctx) error {
	uploadDir := "./ui/images/p"

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot process upload"})
	}

	files := form.File["imagesUpload[]"]
	var uploadedFiles []string

	for _, file := range files {
		// Use original filename
		filePath := filepath.Join(uploadDir, filepath.Base(file.Filename))

		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot save file"})
		}

		uploadedFiles = append(uploadedFiles, file.Filename)

	}

	return c.JSON(fiber.Map{"images": uploadedFiles})
}
