package api_spreadsheet

import (
	"net/http"

	"github.com/Mateus-MS/stock_checker/dev/features/app"
	"github.com/Mateus-MS/stock_checker/dev/features/utils"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/spreadsheet/rows", "PATCH", SpreadsheetUpdateRowsRoute)
}

func SpreadsheetUpdateRowsRoute(w http.ResponseWriter, r *http.Request) {
	query := `
		UPDATE spreadsheet_row c
		SET status = $3
		FROM spreadsheet s
		WHERE c.spreadsheet_key = s.id
		AND s.id = $1
		AND c.sku = $2;
	`

	var err error
	var uuid string
	var sku string
	var status string

	if uuid, err = utils.GetQueryParam(r, "uuid", true, ""); err != nil {
		http.Error(w, "Error while trying to read the param from URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	if sku, err = utils.GetQueryParam(r, "sku", true, ""); err != nil {
		http.Error(w, "Error while trying to read the param from URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	if status, err = utils.GetQueryParam(r, "status", true, ""); err != nil {
		http.Error(w, "Error while trying to read the param from URL: "+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err = app.GetInstance().DB.Exec(query, uuid, sku, status); err != nil {
		http.Error(w, "Error while updating the row in DB: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
