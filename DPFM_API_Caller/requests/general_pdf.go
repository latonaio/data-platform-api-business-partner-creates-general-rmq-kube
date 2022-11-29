package requests

type GeneralPDF struct {
	BusinessPartner          *int   `json:"BusinessPartner"`
	DocType                  string `json:"DocType"`
	DocVersionID             *int   `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner *int   `json:"DocIssuerBusinessPartner"`
	FileName                 string `json:"FileName"`
}
