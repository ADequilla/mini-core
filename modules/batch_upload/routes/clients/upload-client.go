package clients

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"mini-core/middleware/go-utils/database"
	"mini-core/modules/batch_upload/models/request"
	"mini-core/modules/batch_upload/models/response"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func UploadClient(c *fiber.Ctx) error {
	// Parse the multipart form file
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Save the uploaded file to a temporary location
	tempPath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveFile(file, tempPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Compare headers with the downloaded file
	if err := compareHeaders(tempPath); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Read the content of the uploaded file
	clients, err := readExcelFile(tempPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if no data was found in the uploaded file
	if len(clients) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No data found in the uploaded file"})
	}

	// Insert clients into the database
	var result []response.InsertClientModel
	err = database.DBConn.Transaction(func(tx *gorm.DB) error {
		for _, client := range clients {
			// Format time value as string in a format PostgreSQL can handle
			dobStr := client.DOB.Format("2006-01-02")
			err := tx.Raw("SELECT * FROM ewallet_web.client_batch_upload(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", dobStr, client.CID, client.Mobile, client.FirstName, client.LastName, client.MiddleName, client.MaidenFName, client.MaidenLName, client.MaidenMName, client.BirthPlace, client.Sex, client.CivilStatus, client.MemberMaidenFName, client.MemberMaidenLName, client.MemberMaidenMName, client.Email, client.InstitutionCode, client.UnitCode, client.CenterCode).Scan(&result).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the inserted clients as part of the response
	return c.JSON(result)
}

// Helper function to compare headers with the downloaded file
func compareHeaders(filePath string) error {
	// Open the downloaded file
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	// Get headers from the downloaded file
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil || len(rows) == 0 {
		return fmt.Errorf("failed to get rows from downloaded file")
	}
	downloadedHeaders := rows[0]

	// Define expected headers
	expectedHeaders := []string{"Date of Birth", "CID", "Mobile", "First Name", "Last Name", "Middle Name", "Maiden F Name", "Maiden L Name", "Maiden M Name", "Place of Birth", "Sex", "Civil Status", "Member Maiden F Name", "Member Maiden L Name", "Member Maiden M Name", "Email", "InstitutionCode", "UnitCode", "CenterCode"}

	// Check if the headers match
	if !reflect.DeepEqual(expectedHeaders, downloadedHeaders) {
		return fmt.Errorf("uploaded file headers do not match expected headers")
	}

	return nil
}

// Helper function to read the content of an Excel file
func readExcelFile(filePath string) ([]request.Clients, error) {
	// Open the Excel file
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	// Read data from the Excel file
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	// Extract data into Person struct
	var data []request.Clients
	for rowNum, row := range rows[1:] {
		numericDob, err := strconv.Atoi(row[0])
		if err != nil {
			log.Printf("Error parsing numeric date of birth for row: %v, Error: %v", row, err)
			continue
		}
		dob := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).AddDate(0, 0, numericDob) // Adjust the numeric date value to match Go's epoch

		// Log the received data for debugging
		log.Printf("Received data for row %d: %v", rowNum+1, row)

		person := request.Clients{
			DOB:               dob,
			CID:               row[1],
			Mobile:            row[2],
			FirstName:         row[3],
			LastName:          row[4],
			MiddleName:        row[5],
			MaidenFName:       row[6],
			MaidenLName:       row[7],
			MaidenMName:       row[8],
			BirthPlace:        row[9],
			Sex:               row[10],
			CivilStatus:       row[11],
			MemberMaidenFName: row[12],
			MemberMaidenLName: row[13],
			MemberMaidenMName: row[14],
			Email:             row[15],
			InstitutionCode:   row[16],
			UnitCode:          row[17],
			CenterCode:        row[18],
		}
		data = append(data, person)
	}

	return data, nil
}
