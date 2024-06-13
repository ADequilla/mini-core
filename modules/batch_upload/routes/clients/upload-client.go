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
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to get file from request", "details": err.Error()})
	}

	tempPath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveFile(file, tempPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file", "details": err.Error()})
	}

	if err := compareHeaders(tempPath); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Uploaded file headers do not match expected headers", "details": err.Error()})
	}

	clients, err := readExcelFile(tempPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read Excel file", "details": err.Error()})
	}

	if len(clients) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No data found in the uploaded file"})
	}

	var duplicates []response.DuplicateClient // To store details of duplicate clients
	duplicateCount := make(map[string]int)    // To store the count of each duplicate

	// Check for duplicates before starting the transaction
	for _, client := range clients {
		var count int64
		err := database.DBConn.Raw("SELECT COUNT(*) FROM ewallet_web.clients WHERE CID = ? AND Mobile = ? AND First_Name = ? AND Last_Name = ? AND Middle_Name = ? and stats = ?", client.CID, client.Mobile, client.FirstName, client.LastName, client.MiddleName, true).Count(&count).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check for duplicates", "details": err.Error()})
		}
		if count > 0 {
			duplicate := response.DuplicateClient{
				CID:    client.CID,
				Mobile: client.Mobile,
			}
			duplicates = append(duplicates, duplicate)
			duplicateKey := fmt.Sprintf("%s-%s-%s-%s-%s", client.CID, client.Mobile)
			duplicateCount[duplicateKey]++
		}
	}

	if len(duplicates) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":        "Clients uploaded from the file have duplicate data.",
			"data":           duplicates,
			"number of data": len(duplicates),
			"error":          true,
		})
	}

	var result response.InsertClientModel
	var successfulInserts int

	// Start the transaction
	err = database.DBConn.Transaction(func(tx *gorm.DB) error {
		for _, client := range clients {
			dobStr := client.DOB.Format("2006-01-02")
			cid_input := client.CID
			mobile_input := client.Mobile
			fname := client.FirstName
			lname := client.LastName
			mname := client.MiddleName
			mfnam := client.MaidenFName
			mlname := client.MaidenLName
			mmname := client.MaidenMName
			bplace := client.BirthPlace
			sex_input := client.Sex
			cs := client.CivilStatus
			mmfname := client.MemberMaidenFName
			mmlname := client.MemberMaidenLName
			mmmnane := client.MemberMaidenMName
			email_input := client.Email
			i_code := client.InstitutionCode
			u_code := client.UnitCode
			c_code := client.CenterCode

			err := database.DBConn.Raw("SELECT * FROM ewallet_web.client_batch_upload(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", dobStr, cid_input, mobile_input, fname, lname, mname, mfnam, mlname, mmname, bplace, sex_input, cs, mmfname, mmlname, mmmnane, email_input, i_code, u_code, c_code).Scan(&result).Error
			if err != nil {
				return err
			}
			successfulInserts++
		}
		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert clients into the database", "details": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message":        "Clients successfully uploaded",
		"number of data": successfulInserts,
		"error":          false,
	})
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
