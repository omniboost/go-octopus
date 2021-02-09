package octopus

type BookYearServiceData struct {
	BookyearDescription string `json:"bookyearDescription"`
	BookyearKey         struct {
		ID int `json:"id"`
	} `json:"bookyearKey"`
	Closed    bool                        `json:"closed"`
	EndDate   string                      `json:"endDate"`
	Periods   []BookyearPeriodServiceData `json:"periods"`
	StartDate string                      `json:"startDate"`
}

type BookyearPeriodServiceData struct {
	BookyearPeriod int    `json:"bookyearPeriod"`
	EndDate        string `json:"endDate"`
	StartDate      string `json:"startDate"`
}

type JournalServiceData struct {
	BookyearKey struct {
		ID int `json:"id"`
	} `json:"bookyearKey"`
	Closed               bool   `json:"closed"`
	ClosedPeriod         int    `json:"closedPeriod"`
	InsertionType        int    `json:"insertionType"`
	JournalKey           string `json:"journalKey"`
	LastBookedDocumentNr int    `json:"lastBookedDocumentNr"`
	Name                 string `json:"name"`
	ProtectedPeriod      int    `json:"protectedPeriod"`
	CurrencyCode         string `json:"currencyCode,omitempty"`
}

type VatCodeServiceData struct {
	BasePercentage              float64 `json:"basePercentage"`
	Code                        string  `json:"code"`
	DefaultSellBookingAccountNr int     `json:"defaultSellBookingAccountNr"`
	Description                 string  `json:"description"`
	Type                        int     `json:"type"`
}

type AccountServiceData struct {
	AccountKey struct {
		BookyearKey struct {
			ID int `json:"id"`
		} `json:"bookyearKey"`
		ID int `json:"id"`
	} `json:"accountKey"`
	CostCentreType int `json:"costCentreType"`
	Description    struct {
		DescriptionEN string `json:"description_EN"`
		DescriptionFR string `json:"description_FR"`
		DescriptionNL string `json:"description_NL"`
	} `json:"description,omitempty"`
	FiscProfessionalPercentage float64 `json:"fiscProfessionalPercentage"`
	FiscRecupPercentage        float64 `json:"fiscRecupPercentage"`
	VatRecupPercentage         float64 `json:"vatRecupPercentage"`
	PurchaseVatCode            string  `json:"purchaseVatCode,omitempty"`
	SalesVatCode               string  `json:"salesVatCode,omitempty"`
}
