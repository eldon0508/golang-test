package payment

type Payment struct {
	ID          string  `json:"id"`
	Method      string  `json:"method"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}

var PaymentsData = []Payment{
	{ID: "1", Method: "Card", Amount: 12.30, Description: "Card transaction", Status: "Success"},
	{ID: "2", Method: "Bank", Amount: 34.50, Description: "Bank transaction", Status: "Success"},
	{ID: "3", Method: "Third-Party", Amount: 56.70, Description: "Third-Party transaction", Status: "Success"},
}
