package api_spreadsheet

import (
	"mime/multipart"
	"net/http"

	"github.com/Mateus-MS/stock_checker/dev/backend/models/spreadsheet"
	"github.com/Mateus-MS/stock_checker/dev/features/app"
	"github.com/Mateus-MS/stock_checker/dev/features/utils"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func init() {
	app.GetInstance().Router.RegisterRoutes(
		"/api/spreadsheet",
		"POST",
		SpreadsheetCreateRoute,
	)
}

func SpreadsheetCreateRoute(w http.ResponseWriter, r *http.Request) {
	query := `
		INSERT INTO spreadsheet(ID, name) values ($1, $2);
	`
	var err error
	// Initiate the spreadsheet
	spreadsheet := spreadsheet.New()
	spreadsheet.ID = uuid.New()

	// Get the name from the request
	if spreadsheet.Name, err = utils.GetQueryParam(r, "name", true, "Planilha sem nome"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the received excel file
	var file multipart.File
	if file, err = parseSpreadsheet(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the excel file content
	var xlFile *excelize.File
	if xlFile, err = excelize.OpenReader(file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer xlFile.Close()

	// Registering the spreadsheet into DB
	_, err = app.GetInstance().DB.Exec(query, spreadsheet.ID, spreadsheet.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Register all rows into DB
	if err = registerRowsInDB(xlFile, spreadsheet.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the uuid back
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(spreadsheet.ID.String()))
}

func parseSpreadsheet(r *http.Request) (file multipart.File, err error) {
	// Parse with 10MB max
	r.ParseMultipartForm(10 << 20)

	// Get the file
	if file, _, err = r.FormFile("file"); err != nil {
		return file, err
	}

	return file, nil
}

func registerRowsInDB(xlFile *excelize.File, _uuid uuid.UUID) (err error) {
	query := `
		INSERT INTO spreadsheet_row(piece, SKU, spreadsheet_key) VALUES ($1, $2, $3);
	`

	// Get rows
	var rows [][]string
	if rows, err = xlFile.GetRows(xlFile.GetSheetName(0)); err != nil {
		return err
	}

	// Iterate over each row
	for rowIndex, row := range rows {
		// TEMP: ignore the first row
		if rowIndex == 0 {
			continue
		}

		// Register into DB
		// TEMP: rowIndex -1, since i want it on DB to star with the 0 also, but if the above TEMP get fixed, attention to this part also.
		if _, err = app.GetInstance().DB.Exec(query, row[0], row[1], _uuid); err != nil {
			return err
		}
	}

	return nil
}
