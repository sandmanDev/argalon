package responseDM

type RequestResponse struct {
	Success bool
	Msg     string                 `json:"msg"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Errors  map[string]interface{} `json:"errors,omitempty"`
}

type RequestErrorResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}
