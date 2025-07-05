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
	app.GetInstance().Router.RegisterRoutes("/api/spreadsheet/rows", "GET", SpreadsheetGetRowsRoute)
}

// TODO: Add a offset parameter for pagination or layload
func SpreadsheetGetRowsRoute(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT 
			c.sku,
			c.piece,
			c.status
		FROM spreadsheet s
		LEFT JOIN spreadsheet_row c ON c.spreadsheet_key = s.id
		WHERE s.id = $1
		LIMIT 20;
	`

	var uuid string
	var err error
	var rows *sql.Rows
	var result = spreadsheet.Rows()

	// Get the uuid from the request
	if uuid, err = utils.GetQueryParam(r, "uuid", true, ""); err != nil {
		http.Error(w, "Error while trying to read the param from URL: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the query search
	if rows, err = app.GetInstance().DB.Query(query, uuid); err != nil {
		http.Error(w, "Error while querying from DB: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		// Write the values queryied from this row, into product
		var product = spreadsheet.Row()
		if err = rows.Scan(&product.SKU, &product.Piece, &product.Status); err != nil {
			http.Error(w, "Error while scanning the rows into product: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Store in list
		result = append(result, product)
	}

	// Encode the []models.product into a json and return as response
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error while converting into JSON: "+err.Error(), http.StatusInternalServerError)
	}
}
