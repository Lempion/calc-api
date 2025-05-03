package calculate

type Calculations struct {
	ID         string `json:"id" gorm:"primaryKey;type:text"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}
