package db

// Sale represents a single sale stored to the datastore
type Sale struct {
	CustomerName    string  `json:"customer_name" db:"customer_name"`
	ItemDescription string  `json:"item_description" db:"item_description"`
	ItemPrice       float64 `json:"item_price" db:"item_price"`
	Quantity        int     `json:"quantity" db:"quantity"`
	MerchantName    string  `json:"merchant_name" db:"merchant_name"`
	MerchantAddress string  `json:"merchant_address" db:"merchant_address"`
}
