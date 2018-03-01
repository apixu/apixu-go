package response

type Condition struct {
	Code  uint32 `json:"code" xml:"code"`
	Day   string `json:"day" xml:"day"`
	Night string `json:"night" xml:"night"`
	Icon  uint32 `json:"icon" xml:"icon"`
}

type Conditions []Condition
