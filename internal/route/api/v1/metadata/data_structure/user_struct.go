package datastructurego

type UserMatadata struct {
	UserName   string `json:"gin_account_name"`
	Url        string `json:"url"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	AliasName  string `json:"alias"`
	EMail      string `json:"email"`
	Telephone  string `json:"telephone"`
	ERadNumber string `json:"e_rad_number"`
}
