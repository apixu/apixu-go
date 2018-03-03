package response

// Condition defines the weather condition item
type Condition struct {
	Code  uint32 `json:"code" xml:"code"`
	Day   string `json:"day" xml:"day"`
	Night string `json:"night" xml:"night"`
	Icon  uint32 `json:"icon" xml:"icon"`
}

// Conditions defines Condition items list
type Conditions []Condition
