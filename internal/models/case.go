package models

type Case struct {
	Name     string `json:"name"`
	CaseId   uint   `json:"case_id"`
	Duration uint   `json:"duration"`
	Sound    string `json:"sound"`
	Index    uint   `json:"index"`
}
