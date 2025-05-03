package calculate

type CalculationType struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequestType struct {
	Expression string `json:"expression"`
}
