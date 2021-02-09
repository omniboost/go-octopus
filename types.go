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

type FinancialDiversBookingAndAttachmentRequest struct {
	FinancialDiversBookingServiceData FinancialDiversBookingServiceData `json:"financialDiversBookingServiceData"`
	Attachments                       []Attachment                      `json:"attachments"`
}

type FinancialDiversBookingServiceData struct {
	BookyearKey struct {
		ID int `json:"id"`
	} `json:"bookyearKey"`
	JournalKey         string `json:"journalKey"`
	DocumentSequenceNr int    `json:"documentSequenceNr"`
	BookyearPeriodeNr  int    `json:"bookyearPeriodeNr"`
	DocumentDate       string `json:"documentDate"`
	ExchangeRate       int    `json:"exchangeRate"`
	BookingLines       []struct {
		Type               string `json:"type"`
		AccountKey         int    `json:"accountKey"`
		ExternalRelationID int    `json:"externalRelationId"`
		Reference          string `json:"reference"`
		Amount             int    `json:"amount"`
		CostCentreKey      struct {
			ID int `json:"id"`
		} `json:"costCentreKey"`
		CodaInfo struct {
			CodaVersion struct {
				CodaVersion int `json:"codaVersion"`
			} `json:"codaVersion"`
			CodaMoveLineData []string `json:"codaMoveLineData"`
		} `json:"codaInfo"`
	} `json:"bookingLines"`
}

type Attachment struct {
	FileName string `json:"fileName"`
	FileData string `json:"fileData"`
}

type InvoiceServiceData struct {
	BookyearKey struct {
		ID int `json:"id"`
	} `json:"bookyearKey"`
	JournalKey                        string `json:"journalKey"`
	DocumentSequenceNr                int    `json:"documentSequenceNr"`
	BookyearPeriodeNr                 int    `json:"bookyearPeriodeNr"`
	DocumentDate                      string `json:"documentDate"`
	ExpiryDate                        string `json:"expiryDate"`
	CurrencyCode                      string `json:"currencyCode"`
	ExchangeRate                      int    `json:"exchangeRate"`
	RelationIdentificationServiceData struct {
		RelationKey struct {
			ID int `json:"id"`
		} `json:"relationKey"`
		ExternalRelationID int `json:"externalRelationId"`
	} `json:"relationIdentificationServiceData"`
	Comment              string `json:"comment"`
	Reference            string `json:"reference"`
	FinancialDiscount    int    `json:"financialDiscount"`
	CustomFieldValueList []struct {
		CustomFieldKey struct {
			ID int `json:"id"`
		} `json:"customFieldKey"`
		Value string `json:"value"`
	} `json:"customFieldValueList"`

	InvoiceLines []InvoiceLineServiceData `json:"invoiceLines"`
}

type InvoiceLineServiceData struct {
	ExternProductNr    string `json:"externProductNr"`
	Description        string `json:"description"`
	Count              int    `json:"count"`
	Unit               string `json:"unit"`
	UnitPrice          int    `json:"unitPrice"`
	DiscountPercentage int    `json:"discountPercentage"`
	VATCodeKey         string `json:"vatCodeKey"`
	BookingAccountNr   int    `json:"bookingAccountNr"`
	CostCentreKey      struct {
		ID int `json:"id"`
	} `json:"costCentreKey"`
	CustomFieldValueList []struct {
		CustomFieldKey struct {
			ID int `json:"id"`
		} `json:"customFieldKey"`
		Value string `json:"value"`
	} `json:"customFieldValueList"`
	IntrastatServiceData struct {
		IsoCountrycode  string `json:"isoCountrycode"`
		TransactionCode int    `json:"transactionCode"`
		ProductCode     string `json:"productCode"`
		Region          int    `json:"region"`
		Weight          int    `json:"weight"`
		UnitCount       int    `json:"unitCount"`
		TransportCode   int    `json:"transportCode"`
		IncoTerms       string `json:"incoTerms"`
		OriginCountry   string `json:"originCountry"`
	} `json:"intrastatServiceData"`
}
