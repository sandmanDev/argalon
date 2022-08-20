package requestDM

type DivideRequestDM struct {
	ParamterA float64 `json:"a" binding:"required"`
	ParamterB float64 `json:"b" binding:"required"`
}
