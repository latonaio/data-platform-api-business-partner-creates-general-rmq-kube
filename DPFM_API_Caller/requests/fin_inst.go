package requests

type FinInst struct {
	BusinessPartner           *int    `json:"BusinessPartner"`
	FinInstIdentification     *int    `json:"FinInstIdentification"`
	ValidityEndDate           *string `json:"ValidityEndDate"`
	ValidityStartDate         *string `json:"ValidityStartDate"`
	FinInstCountry            string  `json:"FinInstCountry"`
	FinInstNumber             string  `json:"FinInstNumber"`
	FinInstName               string  `json:"FinInstName"`
	FinInstBranchName         string  `json:"FinInstBranchName"`
	SWIFTCode                 string  `json:"SWIFTCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         string  `json:"FinInstControlKey"`
	FinInstAccountName        string  `json:"FinInstAccountName"`
	FinInstAccount            string  `json:"FinInstAccount"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}
