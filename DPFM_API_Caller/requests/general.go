package requests

type General struct {
	BusinessPartner               *int    `json:"BusinessPartner"`
	BusinessPartnerFullName       string  `json:"BusinessPartnerFullName"`
	BusinessPartnerName           string  `json:"BusinessPartnerName"`
	CreationDate                  *string `json:"CreationDate"`
	CreationTime                  *string `json:"CreationTime"`
	Industry                      string  `json:"Industry"`
	LegalEntityRegistration       string  `json:"LegalEntityRegistration"`
	Country                       string  `json:"Country"`
	Language                      string  `json:"Language"`
	Currency                      string  `json:"Currency"`
	LastChangeDate                *string `json:"LastChangeDate"`
	LastChangeTime                *string `json:"LastChangeTime"`
	OrganizationBPName1           string  `json:"OrganizationBPName1"`
	OrganizationBPName2           string  `json:"OrganizationBPName2"`
	OrganizationBPName3           string  `json:"OrganizationBPName3"`
	OrganizationBPName4           string  `json:"OrganizationBPName4"`
	BPGroup1                      string  `json:"BPGroup1"`
	BPGroup2                      string  `json:"BPGroup2"`
	BPGroup3                      string  `json:"BPGroup3"`
	BPGroup4                      string  `json:"BPGroup4"`
	BPGroup5                      string  `json:"BPGroup5"`
	OrganizationFoundationDate    *string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string `json:"OrganizationLiquidationDate"`
	SearchTerm1                   string  `json:"SearchTerm1"`
	SearchTerm2                   string  `json:"SearchTerm2"`
	BusinessPartnerBirthplaceName string  `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked      *bool   `json:"BusinessPartnerIsBlocked"`
	GroupBusinessPartnerName1     string  `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     string  `json:"GroupBusinessPartnerName2"`
	AddressID                     *int    `json:"AddressID"`
	BusinessPartnerIDByExtSystem  string  `json:"BusinessPartnerIDByExtSystem"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}
