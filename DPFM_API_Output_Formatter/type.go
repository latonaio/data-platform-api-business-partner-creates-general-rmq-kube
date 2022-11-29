package dpfm_api_output_formatter

type BusinessPartnerCreates struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	BusinessPartner string `json:"business_partner_code"`
	Deleted         bool   `json:"deleted"`
}

type General struct {
	BusinessPartner               int    `json:"BusinessPartner"`
	BusinessPartnerFullName       string `json:"BusinessPartnerFullName"`
	BusinessPartnerName           string `json:"BusinessPartnerName"`
	CreationDate                  string `json:"CreationDate"`
	CreationTime                  string `json:"CreationTime"`
	Industry                      string `json:"Industry"`
	LegalEntityRegistration       string `json:"LegalEntityRegistration"`
	Country                       string `json:"Country"`
	Language                      string `json:"Language"`
	Currency                      string `json:"Currency"`
	LastChangeDate                string `json:"LastChangeDate"`
	LastChangeTime                string `json:"LastChangeTime"`
	OrganizationBPName1           string `json:"OrganizationBPName1"`
	OrganizationBPName2           string `json:"OrganizationBPName2"`
	OrganizationBPName3           string `json:"OrganizationBPName3"`
	OrganizationBPName4           string `json:"OrganizationBPName4"`
	BPGroup1                      string `json:"BPGroup1"`
	BPGroup2                      string `json:"BPGroup2"`
	BPGroup3                      string `json:"BPGroup3"`
	BPGroup4                      string `json:"BPGroup4"`
	BPGroup5                      string `json:"BPGroup5"`
	OrganizationFoundationDate    string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   string `json:"OrganizationLiquidationDate"`
	SearchTerm1                   string `json:"SearchTerm1"`
	SearchTerm2                   string `json:"SearchTerm2"`
	BusinessPartnerBirthplaceName string `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      string `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked      bool   `json:"BusinessPartnerIsBlocked"`
	GroupBusinessPartnerName1     string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     string `json:"GroupBusinessPartnerName2"`
	AddressID                     int    `json:"AddressID"`
	BusinessPartnerIDByExtSystem  string `json:"BusinessPartnerIDByExtSystem"`
	IsMarkedForDeletion           bool   `json:"IsMarkedForDeletion"`
}

type GeneralPDF struct {
	DocType                  string `json:"DocType"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FileName                 string `json:"FileName"`
}

type Role struct {
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	ValidityStartDate   string `json:"ValidityStartDate"`
}

type FinInst struct {
	FinInstIdentification     int    `json:"FinInstIdentification"`
	ValidityEndDate           string `json:"ValidityEndDate"`
	ValidityStartDate         string `json:"ValidityStartDate"`
	FinInstCountry            string `json:"FinInstCountry"`
	FinInstNumber             string `json:"FinInstNumber"`
	FinInstName               string `json:"FinInstName"`
	FinInstBranchName         string `json:"FinInstBranchName"`
	SWIFTCode                 string `json:"SWIFTCode"`
	InternalFinInstCustomerID int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         string `json:"FinInstControlKey"`
	FinInstAccountName        string `json:"FinInstAccountName"`
	FinInstAccount            string `json:"FinInstAccount"`
	IsMarkedForDeletion       bool   `json:"IsMarkedForDeletion"`
}

type Relationship struct {
	RelationshipNumber          int    `json:"RelationshipNumber"`
	ValidityEndDate             string `json:"ValidityEndDate"`
	ValidityStartDate           string `json:"ValidityStartDate"`
	RelationshipCategory        string `json:"RelationshipCategory"`
	RelationshipBusinessPartner int    `json:"RelationshipBusinessPartner"`
	BusinessPartnerPerson       string `json:"BusinessPartnerPerson"`
	IsStandardRelationship      bool   `json:"IsStandardRelationship"`
	IsMarkedForDeletion         bool   `json:"IsMarkedForDeletion"`
}

type Accounting struct {
	ChartOfAccounts     string `json:"ChartOfAccounts"`
	FiscalYearVariant   string `json:"FiscalYearVariant"`
	IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
}
