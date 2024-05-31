package clients

import (
	"github.com/gofiber/fiber/v2"
)

// func DownloadClientTemplate(c *fiber.Ctx) error {
// 	// Open the template file from the assets folder
// 	templatePath := "assets/clients.xlsx"
// 	templateFile, err := os.Open(templatePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer templateFile.Close()

// 	// Create the directory for saving files if it doesn't exist
// 	desktopPath, err := os.UserHomeDir()
// 	if err != nil {
// 		return err
// 	}
// 	desktopPath = filepath.Join(desktopPath, "Desktop")
// 	excelFileName := "Clients.xlsx"
// 	excelFilePath := filepath.Join(desktopPath, excelFileName)

// 	// Check if the file already exists, if so, append a number to the filename
// 	number := 1
// 	for fileExists(excelFilePath) {
// 		excelFileName = fmt.Sprintf("Clients-%d.xlsx", number)
// 		excelFilePath = filepath.Join(desktopPath, excelFileName)
// 		number++
// 	}

// 	// Create a new file on the user's desktop
// 	newFile, err := os.Create(excelFilePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer newFile.Close()

// 	// Copy the contents of the template file to the new file
// 	_, err = io.Copy(newFile, templateFile)
// 	if err != nil {
// 		return err
// 	}

// 	// Set response headers for file download
// 	c.Set(fiber.HeaderContentType, "application/octet-stream")
// 	c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", excelFileName))

// 	// Create JSON response
// 	jsonResponse := map[string]interface{}{
// 		"status":     "success",
// 		"message":    "Template successfully downloaded",
// 		"excel_path": excelFilePath,
// 	}

// 	// Return JSON response
// 	return c.JSON(jsonResponse)
// }

// // Helper function to check if a file exists
// func fileExists(filename string) bool {
// 	_, err := os.Stat(filename)
// 	return err == nil
// }

func DownloadClientTemplate(c *fiber.Ctx) error {
	excelFileName := "assets/Clients.xlsx"
	if err := c.Download(excelFileName); err != nil {
		return err
	}
	return c.Status(200).SendString(excelFileName)
}
