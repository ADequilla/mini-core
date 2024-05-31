package accounts

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

func UploadAccount(c *fiber.Ctx) error {
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
	accounts, err := ReadExcelFile(tempPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if no data was found in the uploaded file
	if len(accounts) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No data found in the uploaded file"})
	}

	// Insert clients into the database
	var result []response.InsertAccountModel
	err = database.DBConn.Transaction(func(tx *gorm.DB) error {
		for _, account := range accounts {
			// Format time value as string in a format PostgreSQL can handle
			doStr := account.DateOpen.Format("2006-01-02")
			deStr := account.DateEntry.Format("2006-01-02")
			dRStr := account.DateRecognized.Format("2006-01-02")
			drStr := account.DateResigned.Format("2006-01-02")
			err := tx.Raw("SELECT * FROM ewallet_web.account_batch_upload(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", account.AccountNumber, account.Account, account.AccountType, account.AccountDescription, doStr, account.StatusDescription, account.IiID, account.Status, account.Title, account.Classification, account.SubClassification, deStr, dRStr, drStr, account.InstitutionCode, account.BranchCode, account.UnitCode, account.CenterCode, account.UUID, account.SetCid, account.AreaCode, account.Area, account.Balance, account.Withdrawable, account.LeagerBalance).Scan(&result).Error
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
	expectedHeaders := []string{"Account Number", "Account", "Account Type", "Account Description", "Date Open", "Status Description", "iiID", "Status", "Title", "Classification", "Sub Classification", "Date Entry", "Date Recognized", "Date Resigned", "Institution Code", "Branch Code", "Unit Code", "Center Code", "UUID", "CID", "Area Code", "Area", "Balance", "Withdrawable", "Leager Balance"}

	// Check if the headers match
	if !reflect.DeepEqual(expectedHeaders, downloadedHeaders) {
		return fmt.Errorf("uploaded file headers do not match expected headers")
	}

	return nil
}

// Helper function to read the content of an Excel file
func ReadExcelFile(filePath string) ([]request.Accounts, error) {
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
	var data []request.Accounts
	for rowNum, row := range rows[1:] {
		numericdo, err := strconv.Atoi(row[4])
		if err != nil {
			log.Printf("Error parsing numeric date of birth for row: %v, Error: %v", row, err)
			continue
		}
		do := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).AddDate(0, 0, numericdo) // Adjust the numeric date value to match Go's epoch

		numericde, err := strconv.Atoi(row[11])
		if err != nil {
			log.Printf("Error parsing numeric date of birth for row: %v, Error: %v", row, err)
			continue
		}
		de := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).AddDate(0, 0, numericde) // Adjust the numeric date value to match Go's epoch

		numericdR, err := strconv.Atoi(row[12])
		if err != nil {
			log.Printf("Error parsing numeric date of birth for row: %v, Error: %v", row, err)
			continue
		}
		dR := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).AddDate(0, 0, numericdR) // Adjust the numeric date value to match Go's epoch

		numericdr, err := strconv.Atoi(row[13])
		if err != nil {
			log.Printf("Error parsing numeric date of birth for row: %v, Error: %v", row, err)
			continue
		}
		dr := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).AddDate(0, 0, numericdr) // Adjust the numeric date value to match Go's epoch

		// Log the received data for debugging
		log.Printf("Received data for row %d: %v", rowNum+1, row)

		accountType, err := strconv.Atoi(row[2]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}

		status, err := strconv.Atoi(row[7]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}

		title, err := strconv.Atoi(row[8]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		classification, err := strconv.Atoi(row[9]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		subclassification, err := strconv.Atoi(row[10]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		InstitutionCode, err := strconv.Atoi(row[14]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		BranchCode, err := strconv.Atoi(row[15]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		UnitCode, err := strconv.Atoi(row[16]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		CenterCode, err := strconv.Atoi(row[17]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		UUID, err := strconv.Atoi(row[18]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		SetCid, err := strconv.Atoi(row[19]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}
		AreaCode, err := strconv.Atoi(row[20]) // Convert string to integer
		if err != nil {
			return nil, err // Handle conversion error
		}

		balance, err := strconv.ParseFloat(row[22], 64) // Convert string to float64
		if err != nil {
			return nil, err // Handle conversion error
		}

		withdrawable, err := strconv.ParseFloat(row[23], 64) // Convert string to float64
		if err != nil {
			return nil, err // Handle conversion error
		}

		leagerBalance, err := strconv.ParseFloat(row[24], 64) // Convert string to float64
		if err != nil {
			return nil, err // Handle conversion error
		}

		person := request.Accounts{
			AccountNumber:      row[0],
			Account:            row[1],
			AccountType:        accountType,
			AccountDescription: row[3],
			DateOpen:           do,
			StatusDescription:  row[5],
			IiID:               row[6],
			Status:             status,
			Title:              title,
			Classification:     classification,
			SubClassification:  subclassification,
			DateEntry:          de,
			DateRecognized:     dR,
			DateResigned:       dr,
			InstitutionCode:    InstitutionCode,
			BranchCode:         BranchCode,
			UnitCode:           UnitCode,
			CenterCode:         CenterCode,
			UUID:               UUID,
			SetCid:             SetCid,
			AreaCode:           AreaCode,
			Area:               row[21],
			Balance:            balance,
			Withdrawable:       withdrawable,
			LeagerBalance:      leagerBalance,
		}
		data = append(data, person)
	}

	return data, nil
}
