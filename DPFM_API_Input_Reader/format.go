package dpfm_api_input_reader

import (
	"data-platform-api-business-partner-creates-general-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBpExistenceConf() {

}

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.BusinessPartner
	return &requests.General{
		BusinessPartner:               data.BusinessPartner,
		BusinessPartnerFullName:       data.BusinessPartnerFullName,
		BusinessPartnerName:           data.BusinessPartnerName,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		Industry:                      data.Industry,
		LegalEntityRegistration:       data.LegalEntityRegistration,
		Country:                       data.Country,
		Language:                      data.Language,
		Currency:                      data.Currency,
		LastChangeDate:                data.LastChangeDate,
		LastChangeTime:                data.LastChangeTime,
		OrganizationBPName1:           data.OrganizationBPName1,
		OrganizationBPName2:           data.OrganizationBPName2,
		OrganizationBPName3:           data.OrganizationBPName3,
		OrganizationBPName4:           data.OrganizationBPName4,
		BPGroup1:                      data.BPGroup1,
		BPGroup2:                      data.BPGroup2,
		BPGroup3:                      data.BPGroup3,
		BPGroup4:                      data.BPGroup4,
		BPGroup5:                      data.BPGroup5,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
		SearchTerm1:                   data.SearchTerm1,
		SearchTerm2:                   data.SearchTerm2,
		BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
		BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
		BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
		GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
		AddressID:                     data.AddressID,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToGeneralPDF(num int) *requests.GeneralPDF {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.GeneralPDF[num]
	return &requests.GeneralPDF{
		BusinessPartner:          dataBusinessPartner.BusinessPartner,
		DocType:                  data.DocType,
		DocVersionID:             data.DocVersionID,
		DocID:                    data.DocID,
		DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		FileName:                 data.FileName,
	}
}

func (sdc *SDC) ConvertToRole(num int) *requests.Role {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Role[num]
	return &requests.Role{
		BusinessPartner:     dataBusinessPartner.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidityEndDate:     data.ValidityEndDate,
		ValidityStartDate:   data.ValidityStartDate,
	}
}

func (sdc *SDC) ConvertToFinInst(num int) *requests.FinInst {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.FinInst[num]
	return &requests.FinInst{
		BusinessPartner:           dataBusinessPartner.BusinessPartner,
		FinInstIdentification:     data.FinInstIdentification,
		ValidityEndDate:           data.ValidityEndDate,
		ValidityStartDate:         data.ValidityStartDate,
		FinInstCountry:            data.FinInstCountry,
		FinInstNumber:             data.FinInstNumber,
		FinInstName:               data.FinInstName,
		FinInstBranchName:         data.FinInstBranchName,
		SWIFTCode:                 data.SWIFTCode,
		InternalFinInstCustomerID: data.InternalFinInstCustomerID,
		InternalFinInstAccountID:  data.InternalFinInstAccountID,
		FinInstControlKey:         data.FinInstControlKey,
		FinInstAccountName:        data.FinInstAccountName,
		FinInstAccount:            data.FinInstAccount,
		IsMarkedForDeletion:       data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToRelationship(num int) *requests.Relationship {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Relationship[num]
	return &requests.Relationship{
		BusinessPartner:             dataBusinessPartner.BusinessPartner,
		RelationshipNumber:          data.RelationshipNumber,
		ValidityEndDate:             data.ValidityEndDate,
		ValidityStartDate:           data.ValidityStartDate,
		RelationshipCategory:        data.RelationshipCategory,
		RelationshipBusinessPartner: data.RelationshipBusinessPartner,
		BusinessPartnerPerson:       data.BusinessPartnerPerson,
		IsStandardRelationship:      data.IsStandardRelationship,
		IsMarkedForDeletion:         data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToAccounting(num int) *requests.Accounting {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Accounting[num]
	return &requests.Accounting{
		BusinessPartner:     dataBusinessPartner.BusinessPartner,
		ChartOfAccounts:     data.ChartOfAccounts,
		FiscalYearVariant:   data.FiscalYearVariant,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}
