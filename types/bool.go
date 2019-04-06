package types

import (
	"strconv"
)

// Bool is used to convert int 1/0 to bool true/false
type Bool bool

// UnmarshalJSON converts int to bool from JSON
func (b *Bool) UnmarshalJSON(data []byte) (err error) {
	v, err := strconv.ParseBool(string(data))
	*b = Bool(v)
	return
}
