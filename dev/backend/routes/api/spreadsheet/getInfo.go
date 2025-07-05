package api_spreadsheet

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Mateus-MS/stock_checker/dev/backend/models/spreadsheet"
	"github.com/Mateus-MS/stock_checker/dev/features/app"
	"github.com/Mateus-MS/stock_checker/dev/features/utils"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/spreadsheet/info", "GET", SpreadsheetGetInfoRoute)
}

func SpreadsheetGetInfoRoute(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT 
			id,
			name,
			checked_count,
			unchecked_count,
			absent_count
		FROM spreadsheet
		WHERE id = $1;
	`

	var uuid string
	var err error
	var rows *sql.Rows
	var result = spreadsheet.New()

	// Get the uuid from the request
	if uuid, err = utils.GetQueryParam(r, "uuid", true, ""); err != nil {
		http.Error(w, "Error while trying to read the param from URL: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the query
	if rows, err = app.GetInstance().DB.Query(query, uuid); err != nil {
		http.Error(w, "Error while querying from DB: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Scan the rows
	if rows.Next() {
		rows.Scan(&result.ID, &result.Name, &result.Checkeds, &result.Uncheckeds, &result.Absent)
	}

	// Encode the result into a json and returns as response
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error while converting into JSON: "+err.Error(), http.StatusInternalServerError)
	}
}
