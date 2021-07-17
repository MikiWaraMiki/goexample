package domain_model

type Play struct {
	PlayID   string `json:"playID"`
	Name     string `json:"name"`
	TypeName string `json:"typeName"`
}

func (this Play) IsTragedy() bool {
	return this.TypeName == "tragedy"
}

func (this Play) IsComedy() bool {
	return this.TypeName == "comedy"
}
