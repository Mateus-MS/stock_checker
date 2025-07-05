package spreadsheet

type Product struct {
	Piece  string `json:"piece"`
	SKU    string `json:"SKU"`
	Status string `json:"status"`
}

func Row() (p Product) {
	return p
}

func Rows() (p []Product) {
	return p
}
