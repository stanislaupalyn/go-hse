package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeNameRequest struct {
	Name    string `json:"name"`
	NewName string `json:"new-name"`
}

type GetAccountRequest struct {
	Name string `json:"name"`
}

type ChangeAmountRequest struct {
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	NewAmount int    `json:"new-amount"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
