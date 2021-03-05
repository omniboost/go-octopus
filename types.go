package octopus

type BookYearServiceData struct {
	BookyearDescription string `json:"bookyearDescription"`
	BookyearKey         struct {
		ID int `json:"id"`
	} `json:"bookyearKey"`
	Closed    bool                        `json:"closed"`
	EndDate   Date                        `json:"endDate"`
	Periods   []BookyearPeriodServiceData `json:"periods"`
	StartDate Date                        `json:"startDate"`
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

type VATCodeServiceData struct {
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
	VATRecupPercentage         float64 `json:"vatRecupPercentage"`
	PurchaseVATCode            string  `json:"purchaseVatCode,omitempty"`
	SalesVATCode               string  `json:"salesVatCode,omitempty"`
}

type FinancialDiversBookingAndAttachmentRequest struct {
	FinancialDiversBookingServiceData FinancialDiversBookingServiceData `json:"financialDiversBookingServiceData"`
	Attachments                       Attachments                       `json:"attachments"`
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
		Type               string  `json:"type"`
		AccountKey         int     `json:"accountKey"`
		ExternalRelationID int     `json:"externalRelationId"`
		Reference          string  `json:"reference"`
		Amount             float64 `json:"amount"`
		// CostCentreKey      struct {
		// 	ID int `json:"id,omitempty"`
		// } `json:"costCentreKey"`
		CodaInfo struct {
			CodaVersion struct {
				CodaVersion int `json:"codaVersion"`
			} `json:"codaVersion"`
			CodaMoveLineData []string `json:"codaMoveLineData"`
		} `json:"codaInfo"`
	} `json:"bookingLines"`
}

type Attachments []Attachment

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

type RelationServiceData []RelationIdentificationServiceData

type RelationIdentificationServiceData struct {
	Active                            bool   `json:"active"`
	BankAccountNr                     string `json:"bankAccountNr,omitempty"`
	BicCode                           string `json:"bicCode"`
	City                              string `json:"city,omitempty"`
	Client                            bool   `json:"client"`
	ContactPerson                     string `json:"contactPerson,omitempty"`
	CorporationType                   int    `json:"corporationType,omitempty"`
	Country                           string `json:"country"`
	CurrencyCode                      string `json:"currencyCode,omitempty"`
	DefaultBookingAccountClient       int    `json:"defaultBookingAccountClient,omitempty"`
	DefaultBookingAccountSupplier     int    `json:"defaultBookingAccountSupplier,omitempty"`
	DeliveryCountry                   string `json:"deliveryCountry,omitempty"`
	DeliveryPostalCode                string // `json:"deliveryPostalCode,omitempty"`
	DeliveryStreetAndNr               string // `json:"deliveryStreetAndNr,omitempty"`
	Email                             string `json:"email,omitempty"`
	ExpirationDays                    int    `json:"expirationDays,omitempty"`
	ExpirationType                    int    `json:"expirationType,omitempty"`
	ExternCustomerNr                  string // `json:"externCustomerNr,omitempty"`
	FactLanguage                      int    `json:"factLanguage,omitempty"`
	Fax                               string `json:"fax,omitempty,omitempty"`
	FinancialDiscount                 bool   // `json:"financialDiscount,omitempty"`
	FirstName                         string `json:"firstName"`
	IbanAccountNr                     string `json:"ibanAccountNr,omitempty"`
	Name                              string `json:"name"`
	PostalCode                        string `json:"postalCode"`
	Profession                        string `json:"profession,omitempty"`
	RelationIdentificationServiceData struct {
		ExternalRelationID int `json:"externalRelationId"`
		RelationKey        struct {
			ID int `json:"id,omitempty"`
		} `json:"relationKey"`
	} `json:"relationIdentificationServiceData"`
	Remarks        string `json:"remarks,omitempty"`
	SddActive      bool   `json:"sddActive,omitempty"`
	SddMandateType int    `json:"sddMandateType,omitempty"`
	SddSeqtype     int    `json:"sddSeqtype,omitempty"`
	SearchField1   string `json:"searchField1,omitempty"`
	SearchField2   string `json:"searchField2,omitempty"`
	StreetAndNr    string `json:"streetAndNr,omitempty"`
	Supplier       bool   `json:"supplier,omitempty"`
	Telephone      string `json:"telephone,omitempty"`
	URL            string `json:"url,omitempty"`
	VATNr          string `json:"vatNr,omitempty"`
	// Vat Type (Btw plichtigheid):
	// 0 : Onbekend
	// 1 : Belgische BTW-plichtige
	// 4 : Intracommunautaire BTW-plichtige
	// 6 : BTW-plichtige buiten EU
	// 7 : Particulier Belgie
	// 8 : Particulier EU
	// 9 : Particulier niet-EU
	// 10: Niet btw-plichtig
	VATType int `json:"vatType,omitempty"`
}

type BuySellBookingAndAttachmentRequest struct {
	BuySellBookingServiceData BuySellBookingServiceData `json:"buySellBookingServiceData"`
	// Attachments               Attachments               `json:"attachments"`
}

type BuySellBookingServiceData struct {
	BookyearKey struct {
		ID int `json:"id"`
	} `json:"bookyearKey"`
	JournalKey                        string `json:"journalKey"`
	DocumentSequenceNr                int    `json:"documentSequenceNr"`
	RelationIdentificationServiceData struct {
		RelationKey struct {
			ID int `json:"id"`
		} `json:"relationKey"`
		ExternalRelationID int `json:"externalRelationId"`
	} `json:"relationIdentificationServiceData"`
	BookyearPeriodeNr int     `json:"bookyearPeriodeNr"`
	DocumentDate      Date    `json:"documentDate"`
	ExpiryDate        Date    `json:"expiryDate"`
	Comment           string  `json:"comment"`
	Reference         string  `json:"reference"`
	Amount            float64 `json:"amount"`
	CurrencyCode      string  `json:"currencyCode,omitempty"`
	ExchangeRate      float64 `json:"exchangeRate,omitempty"`
	BookingLines      []struct {
		AccountKey int     `json:"accountKey"`
		BaseAmount float64 `json:"baseAmount"`
		VATCodeKey string  `json:"vatCodeKey"`
		VATAmount  float64 `json:"vatAmount"`
		Comment    string  `json:"comment"`
		// CostCentreKey struct {
		// 	ID int `json:"id"`
		// } `json:"costCentreKey"`
		// VatRecupPercentage   int `json:"vatRecupPercentage"`
		// IntrastatServiceData struct {
		// 	IsoCountrycode  string `json:"isoCountrycode"`
		// 	TransactionCode int    `json:"transactionCode"`
		// 	ProductCode     string `json:"productCode"`
		// 	Region          int    `json:"region"`
		// 	Weight          int    `json:"weight"`
		// 	UnitCount       int    `json:"unitCount"`
		// 	TransportCode   int    `json:"transportCode"`
		// 	IncoTerms       string `json:"incoTerms"`
		// 	OriginCountry   string `json:"originCountry"`
		// } `json:"intrastatServiceData"`
		// OverflowBookingServiceData struct {
		// 	BeginDate                 Date `json:"beginDate"`
		// 	EndDate                   Date `json:"endDate"`
		// 	OverflowBookingPeriodType int  `json:"overflowBookingPeriodType"`
		// 	AccountNr                 int  `json:"accountNr"`
		// } `json:"overflowBookingServiceData"`
	} `json:"bookingLines"`
	PaymentMethod int `json:"paymentMethod"`
}
