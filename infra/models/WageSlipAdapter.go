package models

type WageSlipAdapter struct {
	ListObjects []struct {
		DocumentaryReference string `json:"referenceDocumentaire"`
	} `json:"listeObjets"`
}
