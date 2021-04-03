package responder

type CommonResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
}

type AdvanceCommonResponse struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}
