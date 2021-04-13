package dtos

type AutocompleteRequest struct {
	Prefix string
}

type AutocompleteResponse struct {
	Result []string `json:"result"`
}
