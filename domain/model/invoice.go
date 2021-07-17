package domain_model

type Invoice struct {
	Customer    string        `json:"customer"`
	Performance []Performance `json:"performances"`
}
