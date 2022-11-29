package requests

type Relationship struct {
	BusinessPartner             *int    `json:"BusinessPartner"`
	RelationshipNumber          *int    `json:"RelationshipNumber"`
	ValidityEndDate             *string `json:"ValidityEndDate"`
	ValidityStartDate           *string `json:"ValidityStartDate"`
	RelationshipCategory        string  `json:"RelationshipCategory"`
	RelationshipBusinessPartner *int    `json:"RelationshipBusinessPartner"`
	BusinessPartnerPerson       string  `json:"BusinessPartnerPerson"`
	IsStandardRelationship      *bool   `json:"IsStandardRelationship"`
	IsMarkedForDeletion         *bool   `json:"IsMarkedForDeletion"`
}
