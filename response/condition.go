package response

// Condition defines the weather condition item
type Condition struct {
	Code  uint32 `json:"code" xml:"ROW>code"`
	Day   string `json:"day" xml:"ROW>day"`
	Night string `json:"night" xml:"ROW>night"`
	Icon  uint32 `json:"icon" xml:"ROW>icon"`
}

// Conditions defines Condition items list
type Conditions []Condition
